package cmd

import (
	"github.com/goldins/slappley-award/lib"
	"net/http"
)

const COMMAND = "science"
const USERNAME = "sciencebot"
const BASE_URL = "https://masterofallscience.com/"
const SEARCH_URL = BASE_URL + "api/search?q=%s"
const CAPTION_URL = BASE_URL + "api/caption?e=%s&t=%d"

// e.g. meme/episode/timestamp.jpg?b64lines=base64encodedCaptionText
// note: no `api` prefix.
const CAPTIONED_IMAGE_URL = BASE_URL + "meme/%s/%d.jpg?b64lines=%s"
const MESSAGE_COLOR = "#A1D6F0"

func ScienceHandler() http.HandlerFunc {
	config := lib.NewConfig(COMMAND, USERNAME, SEARCH_URL, CAPTION_URL, CAPTIONED_IMAGE_URL, MESSAGE_COLOR)
	return lib.FetchHandler(config)
}

func ActionHandler() http.HandlerFunc {
	return lib.CallbackHandler()
}
