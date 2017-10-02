// http://guzalexander.com/2017/09/15/cowsay-slack-command.html
package cmd

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"math/rand"
	"time"
    "encoding/base64"
)

const BASE_URL = "https://masterofallscience.com/"
const SEARCH_URL = BASE_URL + "api/search?q=%s"
const CAPTION_URL = BASE_URL + "api/caption?e=%s&t=%d"
const IMAGE_URL = BASE_URL + "meme/%s/%d.jpg?b64lines=%s" // Episode/Timestamp.jpg?b64lines=base64encoded

// todo: move to a lib?
type SlackMessage struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	Channel  string `json:"channel"`
	Icon     string `json:"icon_emoji"`
}

// https://masterofallscience.com/api/caption?e=S01E09&t=40165
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
	/**
	{"Id":10,"Key":"S01E09","Season":1,"EpisodeNumber":9,"Title":"Something Ricked This Way Comes","Director":"John Rice","Writer":"Mike McMahan","OriginalAirDate":"24-Mar-14","WikiLink":""}
	 */
}

type Subtitle struct {
	Id                      int
	RepresentativeTimestamp int
	Episode                 string
	StartTimestamp          int
	EndTimestamp            int
	Content                 string // <-- need this
	Language                string
}

type Frame struct {
	Id        int
	Episode   string
	Timestamp int
}

type CaptionResponse struct {
	Episode  Episode
	Frame    Frame
	Subtitles []Subtitle // <-- this here
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

		query := r.PostFormValue("text")
		// user := r.PostFormValue("user_name")
		// channel := r.PostFormValue("channel_name")
		channelId := r.PostFormValue("channel_id")

		randomItem := getSearchResult(query)
		captionText := getCaptionResult(randomItem)
		imageUrl := getImage(randomItem, captionText)
		log.Printf("%s", imageUrl)

		m := SlackMessage{
			Text:     fmt.Sprintf("%s", imageUrl),
			Username: "sciencebot",
			Channel:  channelId,
			Icon:     ":d20:",
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, string(m.Text))
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
	encoded, err := base64.StdEncoding.DecodeString(text)
	log.Printf("%s", err)
	return fmt.Sprintf(IMAGE_URL, item.Episode, item.Timestamp, string(encoded))
}