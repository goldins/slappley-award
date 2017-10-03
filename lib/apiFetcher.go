package lib

import (
	"net/http"
	"os"
	"log"
	"net/url"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"math/rand"
)

func FetchHandler(config *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		token := os.Getenv("SCIENCE_TOKEN")
		if token != r.PostFormValue("token") {
			log.Printf("%s", r.PostFormValue("token"))
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		query := url.QueryEscape(r.PostFormValue("text"))
		// user := r.PostFormValue("user_name")
		// channel := r.PostFormValue("channel_name")
		channelId := r.PostFormValue("channel_id")

		randomItem := getSearchResult(config.searchUrl, query)
		captionText := getCaptionResult(config.captionUrl, randomItem)
		imageUrl := getImage(config.imageUrl, randomItem, captionText)

		a := Attachment{
			ImageUrl: imageUrl,
		}

		m := SlackMessage{
			ResponseType: "in_channel",
			Text:         captionText,
			Username:     "sciencebot",
			Channel:      channelId,
			Icon:         ":d20:",
			Attachments:  []Attachment{a},
		}

		w.Header().Add("Content-Type", "application/json")
		b, err := json.Marshal(m)
		if err != nil {
			return
		}
		log.Print(string(b))
		fmt.Fprintf(w, string(b))
	}
}

func getSearchResult(url string, q string) SearchResponse {
	searchUrl := fmt.Sprintf(url, q)

	var searchJson []SearchResponse
	getJson(searchUrl, &searchJson)
	// pick a random item from 0 to length
	return searchJson[rand.Intn(len(searchJson)-1)]
}

func getCaptionResult(url string, item SearchResponse) string {
	captionUrl := fmt.Sprintf(url, item.Episode, item.Timestamp)
	log.Printf("%s", captionUrl)
	var captionJson CaptionResponse
	getJson(captionUrl, &captionJson)
	return captionJson.Subtitles[0].Content
}

func getImage(url string, item SearchResponse, text string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	return fmt.Sprintf(url, item.Episode, item.Timestamp, string(encoded))
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
