package main

import (
	"errors"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	"github.com/gofiber/template/html/v2"
	"github.com/patrickmn/go-cache"
)

func main() {
	// Setup .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	JWT_SECRET := []byte(LoadEnvRequired("JWT_SECRET"))

	// Check that the environment has yt-dlp. If it does, try to update to the latest version
	{
		ytdlp_version := RequireCommandExecution("yt-dlp --version")
		if ytdlp_version == "" {
			log.Fatalf("Unable to find yt-dlp... Install it.")
		}
		log.Println("yt-dlp version: " + ytdlp_version)

		ytdlp_rm_cache_dir_output := RequireCommandExecution("yt-dlp --rm-cache-dir")
		log.Println("yt-dlp remove cache attempt: " + ytdlp_rm_cache_dir_output)

		ytdlp_update_output := RequireCommandExecution("yt-dlp -U")
		log.Println("yt-dlp update attempt output: \n" + ytdlp_update_output) // tbh, I don't really care if the update succeeds
	}

	// setup services
	cache_ram := cache.New(180*time.Minute, 180*time.Minute)

	engine := html.New("./views", ".html")
	engine.AddFunc(
		"htmlsafe", func(s string) template.HTML {
			return template.HTML(s)
		},
	)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "./assets")

	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"PageType": "index",
		})
	})

	// takes "id" / "url" query parameters
	app.Get("/api/v1/youtube-video-metadata", func(ctx fiber.Ctx) error {
		params := ctx.Queries()
		query_id := params["id"]
		query_url := params["url"]

		if query_id == "" && query_url == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Bad Request",
				"message": "You have to supply either a video id or an url to a video",
			})
		}

		// Get id (which cannot be trusted yet) either from query or url
		var possibly_invalid_id string
		if query_id != "" {
			possibly_invalid_id = query_id
		} else {
			var err error
			possibly_invalid_id, err = GetIdFromYoutubeUrl(query_url)
			if err != nil {
				if errors.Is(err, ErrUnableToParseUrl) {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"message": "Unable to parse the url supplied",
					})
				} else if errors.Is(err, ErrNotYoutubeUrl) {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"message": "Provided url is not officialy from youtube",
					})
				} else if errors.Is(err, ErrNoIdProvidedInsideUrl) {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"message": "The url you provided does not include an id",
					})
				} else {
					return ErrTypeNotHandledResponse(ctx)
				}
			}
		}

		// additional cleanup, probably unnecessary
		possibly_invalid_id = strings.TrimSpace(possibly_invalid_id)
		possibly_invalid_id = strings.TrimRight(possibly_invalid_id, "/")

		if !IsYoutubeIdValid(possibly_invalid_id) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "The youtube video ID provided is NOT valid.",
			})
		}
		valid_youtube_id := possibly_invalid_id // reassign for clarity

		// check cache, looking up info is long and expensive
		if cached_info, found := cache_ram.Get(valid_youtube_id); found {
			return ctx.JSON(cached_info)
		}

		video_info, err := GetYoutubeDownloadUrl(valid_youtube_id)
		if err != nil {
			if errors.Is(err, ErrYtdlCommandFailed) {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Some error with yt-dlp occured. Please report: https://github.com/TDiblik/youtube-downloader-clientside/issues/new.",
				})
			} else if errors.Is(err, ErrYtdlUnableToParseStdout) {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Unable to parse yt-dlp output. Please report: https://github.com/TDiblik/youtube-downloader-clientside/issues/new.",
				})
			} else {
				return ErrTypeNotHandledResponse(ctx)
			}
		}
		cache_ram.Set(valid_youtube_id, video_info, cache.DefaultExpiration)

		// generate jwt for the cloudflare workers proxy
		proxy_token := jwt.NewWithClaims(jwt.SigningMethodHS256, ProxyJwtClaims{
			*video_info,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		})
		proxy_token_signed, err := proxy_token.SignedString(JWT_SECRET)
		if err != nil {
			// todo
			return ErrTypeNotHandledResponse(ctx)
		}

		return ctx.JSON(fiber.Map{"proxy_token": proxy_token_signed})
	})

	app.Listen(":3000")
}
