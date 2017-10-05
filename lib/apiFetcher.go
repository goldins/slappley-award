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
	"errors"
	"strings"
)

var _config *Config
var _channelId string

func SendHandler(config *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := r.PostFormValue("payload")
		var a Attachment
		json.Unmarshal([]byte(payload), &a)
		var actionValue ActionValue
		json.Unmarshal([]byte(a.Actions[0].Value), &actionValue)

		newAttachment := Attachment{
			Title:    actionValue.Text,
			ImageUrl: actionValue.Url,
		}

		m := SlackMessage{
			ResponseType: "in_channel",
			Text:         a.Actions[0].Text,
			Username:     _config.username,
			Channel:      _channelId,
			Icon:         ":d20:",
			Attachments:  []Attachment{newAttachment},
		}
		handleReturn(w, m)
	}
}

func FetchHandler(config *Config) http.HandlerFunc {
	_config = config
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		token := os.Getenv("SCIENCE_TOKEN")
		if token != r.PostFormValue("token") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		query := url.QueryEscape(r.PostFormValue("text"))
		log.Printf("Got request for %s from %s", query, r.PostFormValue("user_name"))
		// user := r.PostFormValue("user_name")
		// channel := r.PostFormValue("channel_name")
		_channelId := r.PostFormValue("channel_id")

		searchItem, err := getSearchResult(query)
		if err != nil {
			handleReturn(w, makeErrorMessage(err))
			return
		}

		captionText, err := getCaptionResult(searchItem)
		if err != nil {
			handleReturn(w, makeErrorMessage(err))
			return
		}

		imageUrl, err := getCaptionedImage(searchItem, captionText)
		if err != nil {
			handleReturn(w, makeErrorMessage(err))
			return
		}

		av := ActionValue{
			Url:     imageUrl,
			Text:    captionText,
			Args:    captionText,
			Command: _config.command,
		}

		avm, err := json.Marshal(av)
		if err != nil {
			handleReturn(w, makeErrorMessage(err))
			return
		}

		message := SlackMessage{
			ResponseType: "ephemeral",
			Username:     _config.username,
			Channel:      _channelId,
			Icon:         ":d20:",
			Attachments: []Attachment{{
				ImageUrl:   imageUrl,
				CallbackId: _config.command + "_callback",
				Actions: []Action{{
					Name:  "send",
					Text:  "Send",
					Type:  "button",
					Style: "good",
					Value: string(avm),
				}},
			}},
		}

		handleReturn(w, message)
		return
	}
}

/*
 * Returns a set of SearchResponses from the search endpoint. Current API returns a max of 36 items.
 */
func getSearchResult(q string) (SearchResponse, error) {
	searchUrl := fmt.Sprintf(_config.searchUrl, q)

	var searchJson []SearchResponse
	getJson(searchUrl, &searchJson)
	// pick a random item from 0 to length
	// todo: start with the first and add shuffle
	numResults := len(searchJson)
	if numResults > 0 {
		return searchJson[rand.Intn(numResults-1)], nil
	}
	return SearchResponse{}, errors.New(fmt.Sprintf("no results for %s", q))
}

/*
 * Returns the first caption given the SearchResponse (based on Episode and Timestamp)
 */
func getCaptionResult(item SearchResponse) (string, error) {
	captionUrl := fmt.Sprintf(_config.captionUrl, item.Episode, item.Timestamp)
	var captionJson CaptionResponse
	err := getJson(captionUrl, &captionJson)
	if err != nil {
		return "", err
	}
	subtitles := captionJson.Subtitles
	// todo: return collection of Content strings
	return subtitles[0].Content, nil
}

/*
 * Returns an image URL of the frame with a text overlay.
 */
func getCaptionedImage(item SearchResponse, text string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	encoded = strings.Replace(encoded, "/", "_", -1)
	return fmt.Sprintf(_config.captionedImageUrl, item.Episode, item.Timestamp, string(encoded)), nil
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func handleReturn(w http.ResponseWriter, message SlackMessage) error {
	w.Header().Add("Content-Type", "application/json")
	marshalled, err := json.Marshal(message)
	if err != nil {
		fmt.Fprintf(w, string(marshalled))
		return err
	}
	log.Print(string(marshalled))
	fmt.Fprintf(w, string(marshalled))
	return nil
}

func makeErrorMessage(err error) SlackMessage {
	return SlackMessage{
		Text:     err.Error(),
		Username: _config.username,
		Channel:  _channelId,
		Icon:     ":d20:",
	}
}
