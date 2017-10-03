package cmd

import (
	"github.com/goldins/slappley-award/lib"
	"net/http"
)

const BASE_URL = "https://masterofallscience.com/"
const SEARCH_URL = BASE_URL + "api/search?q=%s"
const CAPTION_URL = BASE_URL + "api/caption?e=%s&t=%d"
const IMAGE_URL = BASE_URL + "meme/%s/%d.jpg?b64lines=%s" // Episode/Timestamp.jpg?b64lines=base64encoded

func ScienceHandler() http.HandlerFunc {
	config := lib.NewConfig(SEARCH_URL, CAPTION_URL, IMAGE_URL)
	return lib.FetchHandler(config)
}
