// http://guzalexander.com/2017/09/15/cowsay-slack-command.html
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
	"math/big"
)

const URL = "https://masterofallscience.com/api/search?q="

// todo: move to a lib?
type SlackMessage struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	Channel  string `json:"channel"`
	Icon     string `json:"icon_emoji"`
}

type respItem struct {
	Id		  big.Int `json:"number"`
	Episode   string `json:"string"`
	Timestamp big.Int `json:"number"`
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

		url := string(URL + query)
		items := new([]respItem) // or &Foo{}
		getJson(url, items)
		log.Printf("%s", items)

		m := SlackMessage{
			Text:     fmt.Sprintf("%s", items),
			Username: "sciencebot",
			Channel:  channelId,
			Icon:     ":d20:",
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, string(m.Text))
	}
}
