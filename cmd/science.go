package cmd

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"math/rand"
	"time"
	"encoding/base64"
	"net/url"
	"os"
)

const BASE_URL = "https://masterofallscience.com/"
const SEARCH_URL = BASE_URL + "api/search?q=%s"
const CAPTION_URL = BASE_URL + "api/caption?e=%s&t=%d"
const IMAGE_URL = BASE_URL + "meme/%s/%d.jpg?b64lines=%s" // Episode/Timestamp.jpg?b64lines=base64encoded

type ActionValue struct {
	Url     string `json:"url"`
	Text    string `json:"text"`
	Args    string `json:"args"`
	Command string `json:"command"`
}

type Action struct {
	Name  string      `json:"name"`
	Text  string      `json:"text"`
	Type  string      `json:"type"`
	Style string      `json:"style"`
	Value ActionValue `json:"value"`
}

type Attachment struct {
	Title    string   `json:"title"`
	ImageUrl string   `json:"image_url"`
	Actions  []Action `json:"actions"`
}

// todo: move to a lib?
type SlackMessage struct {
	ResponseType string       `json:"response_type"`
	Text         string       `json:"text"`
	Username     string       `json:"username"`
	Channel      string       `json:"channel"`
	Icon         string       `json:"icon_emoji"`
	Attachments  []Attachment `json:"attachments"`
}

type Episode struct {
	Id              int
	Key             string
	Season          int
	EpisodeNumber   int
	Title           string
	Director        string
	Writer          string
	OriginalAirDate string
	WikiLink        string
}

type Subtitle struct {
	Id                      int
	RepresentativeTimestamp int
	Episode                 string
	StartTimestamp          int
	EndTimestamp            int
	Content                 string
	Language                string
}

type Frame struct {
	Id        int
	Episode   string
	Timestamp int
}

type CaptionResponse struct {
	Episode   Episode
	Frame     Frame
	Subtitles []Subtitle
}

type SearchResponse struct {
	Id        int
	Episode   string
	Timestamp int
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func ScienceHandler() http.HandlerFunc {
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

		randomItem := getSearchResult(query)
		captionText := getCaptionResult(randomItem)
		imageUrl := getImage(randomItem, captionText)

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

func getSearchResult(q string) SearchResponse {
	searchUrl := fmt.Sprintf(SEARCH_URL, q)

	var searchJson []SearchResponse
	getJson(searchUrl, &searchJson)
	// pick a random item from 0 to length
	return searchJson[rand.Intn(len(searchJson)-1)]
}

func getCaptionResult(item SearchResponse) string {
	captionUrl := fmt.Sprintf(CAPTION_URL, item.Episode, item.Timestamp)
	log.Printf("%s", captionUrl)
	var captionJson CaptionResponse
	getJson(captionUrl, &captionJson)
	return captionJson.Subtitles[0].Content
}

func getImage(item SearchResponse, text string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	return fmt.Sprintf(IMAGE_URL, item.Episode, item.Timestamp, string(encoded))
}
