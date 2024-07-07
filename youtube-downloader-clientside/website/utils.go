package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
)

var (
	ErrUnableToParseUrl        = errors.New("unable_to_parse_provided_url")
	ErrNotYoutubeUrl           = errors.New("provided_url_is_not_youtube")
	ErrNoIdProvidedInsideUrl   = errors.New("provided_url_does_not_include_id")
	ErrYtdlCommandFailed       = errors.New("ytdl_command_failed")
	ErrYtdlUnableToParseStdout = errors.New("ytdl_unable_to_parse_stdout")
)

// Runs the command and exits on an error
func RequireCommandExecution(command string) string {
	stdout, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatalf("Error running the command (%s): %s\n", command, err)
	}
	stdout_string := strings.TrimSpace(string(stdout))
	return stdout_string
}

func LoadEnvRequired(env_name string) string {
	env_value := os.Getenv(env_name)
	if len(env_value) == 0 {
		log.Fatal(env_name + " env variable cannot be empty")
	}
	return env_value
}

func GetIdFromYoutubeUrl(query_url string) (string, error) {
	parsed_url, err := url.Parse(query_url)
	if err != nil {
		return "", ErrUnableToParseUrl
	}

	hostname := parsed_url.Hostname()
	hostname = strings.TrimPrefix(hostname, "www.")
	if hostname != "youtube.com" && hostname != "youtu.be" && hostname != "youtube-nocookie.com" {
		return "", ErrNotYoutubeUrl
	}

	// try to get from https://youtu.be/{ID} and https://www.youtube-nocookie.com/**/{id}
	if hostname == "youtu.be" || hostname == "youtube-nocookie.com" {
		possible_id := path.Base(query_url)

		// handles https://youtu.be/ input <-- (possible_id would be youte.be)
		if possible_id == hostname {
			return "", ErrNoIdProvidedInsideUrl
		}

		return possible_id, nil
	}

	// try to get from https://www.youtube.com/watch?v={ID}
	url_query := parsed_url.Query()
	id_from_v_param := url_query.Get("v")
	if !url_query.Has("v") || id_from_v_param == "" {
		return "", ErrNoIdProvidedInsideUrl
	}

	return id_from_v_param, nil
}

// basically the following regex, translated: ^[A-Za-z0-9_-]{11}$
func IsYoutubeIdValid(id string) bool {
	if len(id) != 11 {
		return false
	}

	for _, r := range id {
		if !((r >= 'a' && r <= 'z') || // must be letter between 'a' and 'z'
			(r >= 'A' && r <= 'Z') || // OR must be letter between 'A' and 'Z'
			(r >= '0' && r <= '9') || // OR must be number between '0' and '9'
			r == '_' || r == '-') { // OR must be '-' or '_'
			return false
		}
	}

	return true
}

// make sure to sanitize EVERYTHING before appending it to a command!
// this actually executes shell commands on the server, if it's not sanitized properly,
// it could lead to a nasty RCE!
func GetYoutubeDownloadUrl(video_id string) (*YoutubeVideoMetadata, error) {
	command := "yt-dlp -i --no-warnings --no-config-locations --no-cache-dir " +
		" --socket-timeout 5 --retries 5 --retry-sleep 5 " +
		" --force-ipv4 --no-playlist " +
		" --proxy socks5://127.0.0.1:9050 -J " +
		" \"https://www.youtube.com/watch?v=" + video_id + "\" "

	// setup timeout for 30s. If I don't get a response by then, it's probably too late anyways
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stdout, err := exec.CommandContext(ctx, "bash", "-c", command).Output()
	if err != nil {
		log.Printf("Error running the command (%s): %s\n", command, err)
		return nil, ErrYtdlCommandFailed
	}

	var video_info YoutubeVideoMetadata
	if err := json.Unmarshal(stdout, &video_info); err != nil {
		log.Printf("Error while parsing stdout: \n ----------------- \n %s \n \n %s \n ----------------- \n", stdout, err)
		return nil, ErrYtdlUnableToParseStdout
	}

	return &video_info, nil
}

func ErrTypeNotHandledResponse(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Error type not handled on the server. Please report: https://github.com/TDiblik/youtube-downloader-clientside/issues/new",
	})
}
