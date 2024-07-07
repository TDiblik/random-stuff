package main

import (
	"github.com/golang-jwt/jwt/v5"
)

type YoutubeVideoMetadata struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Formats []struct {
		FormatID   string  `json:"format_id"`
		FormatNote string  `json:"format_note,omitempty"`
		Ext        string  `json:"ext"`
		Protocol   string  `json:"protocol"`
		Acodec     string  `json:"acodec,omitempty"`
		Vcodec     string  `json:"vcodec"`
		URL        string  `json:"url"`
		Width      int     `json:"width,omitempty"`
		Height     int     `json:"height,omitempty"`
		Fps        float64 `json:"fps,omitempty"`

		Resolution  string  `json:"resolution"`
		AspectRatio float64 `json:"aspect_ratio"`

		AudioExt      string  `json:"audio_ext"`
		VideoExt      string  `json:"video_ext"`
		Vbr           float64 `json:"vbr"`
		Abr           float64 `json:"abr"`
		Format        string  `json:"format"`
		Quality       float64 `json:"quality,omitempty"`
		Filesize      int     `json:"filesize,omitempty"`
		AudioChannels int     `json:"audio_channels,omitempty"`
	} `json:"formats"`
	Thumbnails []struct {
		URL        string `json:"url"`
		Preference int    `json:"preference"`
		Height     int    `json:"height,omitempty"`
		Width      int    `json:"width,omitempty"`
		Resolution string `json:"resolution,omitempty"`
	} `json:"thumbnails"`

	Thumbnail string `json:"thumbnail"`

	Description string   `json:"description"`
	ChannelURL  string   `json:"channel_url"`
	Duration    int      `json:"duration"`
	ViewCount   int      `json:"view_count"`
	AgeLimit    int      `json:"age_limit"`
	Categories  []string `json:"categories"`
	Tags        []string `json:"tags"`

	AutomaticCaptions map[string][]struct {
		URL      string `json:"url"`
		Ext      string `json:"ext"`
		Protocol string `json:"protocol"`
	} `json:"automatic_captions"`

	Subtitles map[string][]struct {
		URL      string `json:"url"`
		Ext      string `json:"ext"`
		Protocol string `json:"protocol"`
	} `json:"subtitles"`

	Location  string `json:"location,omitempty"`
	Fulltitle string `json:"fulltitle"`
}

type ProxyJwtClaims struct {
	VideoMetadata YoutubeVideoMetadata `json:"video_metadata"`
	jwt.RegisteredClaims
}
