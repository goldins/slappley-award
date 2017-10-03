package lib

import (
	"net/http"
	"time"
)

type Config struct {
	searchUrl  string
	captionUrl string
	imageUrl   string
	client     *http.Client
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func NewConfig(searchUrl string, captionUrl string, imageUrl string) *Config {
	return &Config{
		searchUrl:  searchUrl,
		captionUrl: captionUrl,
		imageUrl:   imageUrl,
		client:     myClient,
	}
}
