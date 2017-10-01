package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/kelseyhightower/envconfig"
)

type SlackMessage struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	Channel  string `json:"channel"`
	Icon     string `json:"icon_emoji"`
}

func scienceHandler(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		desc := r.PostFormValue("text")
		user := r.PostFormValue("user_name")
		channel := r.PostFormValue("channel_name")
		channelId := r.PostFormValue("channel_id")

		m := SlackMessage{
			Text:     fmt.Sprintf("Hi *<@%s>*. echoing %s", user, desc),
			Username: "science",
			Channel:  channelId,
			Icon:     ":rick:",
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		err := enc.Encode(m)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}

		_, err = http.Post(c.SlackUrl, "text/json", &buf)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}

		log.Printf("said hi to %s in %s (%s)", user, channel, channelId)
	}
}

type Config struct {
	Listen   string
	SlackUrl string `envconfig:"slack_url"`
}

func main() {
	c := Config{Listen: ":8000"}

	err := envconfig.Process("slappley", &c)
	if err != nil {
		log.Fatal("Getting config: " + err.Error())
	}

	http.HandleFunc("/slappley", scienceHandler(c))

	log.Fatal(http.ListenAndServe(c.Listen, nil))
}

