package lib

import (
	"net/http"
	"time"
)

type Config struct {
	username          string
	searchUrl         string
	captionUrl        string
	captionedImageUrl string
	client            *http.Client
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func NewConfig(username string, searchUrl string, captionUrl string, captionedImageUrl string) *Config {
	return &Config{
		username:          username,
		searchUrl:         searchUrl,
		captionUrl:        captionUrl,
		captionedImageUrl: captionedImageUrl,
		client:            myClient,
	}
}
