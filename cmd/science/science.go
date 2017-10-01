// http://guzalexander.com/2017/09/15/cowsay-slack-command.html

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"bytes"
	sc "github.com/goldins/slappley-award"
	"os"
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

		resp := sc.getScience(desc)

		m := SlackMessage{
			Text:     fmt.Sprintf("Hey, *<@%s>*, echo: %s", user, resp),
			Username: "rollbot",
			Channel:  channelId,
			Icon:     ":d20:",
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		err := enc.Encode(m)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}

		log.Printf("Said Hey to %s in %s(%s)", user, channel, channelId)
	}
}

type Config struct {
	Listen   string
	token string
}

func main() {
	c := Config{Listen: ":8080", token: os.Getenv("SCIENCE_TOKEN")}

	// http.HandleFunc("/roll", rollHandler(c, false))
	http.HandleFunc("/science", scienceHandler(c));
	log.Fatal(http.ListenAndServe(c.Listen, nil))
}
