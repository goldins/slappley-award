package lib

import (
	"net/http"
	"time"
)

type Config struct {
	command           string
	username          string
	searchUrl         string
	captionUrl        string
	captionedImageUrl string
	messageColor      string
	client            *http.Client
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func NewConfig(command string, username string, searchUrl string, captionUrl string, captionedImageUrl string,
	messageColor string) *Config {
	return &Config{
		command:           command,
		username:          username,
		searchUrl:         searchUrl,
		captionUrl:        captionUrl,
		captionedImageUrl: captionedImageUrl,
		messageColor:      messageColor,
		client:            myClient,
	}
}
