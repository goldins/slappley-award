package lib

import (
	"net/http"
	"encoding/json"
	"log"
	"strings"
)

func CallbackHandler(config *Config) http.HandlerFunc {
	_config = config
	return func(w http.ResponseWriter, r *http.Request) {
		payload := r.PostFormValue("payload")
		var response CallbackResponse
		json.Unmarshal([]byte(payload), &response)
		var actionValue ActionValue
		json.Unmarshal([]byte(response.Actions[0].Value), &actionValue)

		switch actionValue.Args {
		case "send":
			sendHandler(w, actionValue, response)
		case "cancel":
			cancelHandler(w, actionValue, response)
		case "shuffle":
			shuffleAction(w, actionValue, response)
		default:
			log.Panic("invalid action " + actionValue.Args)
		}
	}
}

func shuffleAction(w http.ResponseWriter, value ActionValue, payload CallbackResponse) {
	// todo: refactor FetchHandler to reuse most of its logic
}

func sendHandler(w http.ResponseWriter, value ActionValue, payload CallbackResponse) {
	updateMessage(w, "sent")
	newAttachment := Attachment{
		ImageUrl: value.Url,
		Color:    _config.messageColor,
	}

	newM := SlackMessage{
		ResponseType:    "in_channel",
		Text:            value.Text,
		Username:        _config.username,
		Channel:         _channelId,
		Attachments:     []Attachment{newAttachment},
		ReplaceOriginal: false,
	}

	marshalled, err := json.Marshal(newM)
	if err != nil {
		log.Panic(err)
		return
	}

	r, err := myClient.Post(
		payload.ResponseUrl,
		"application/json",
		strings.NewReader(string(marshalled)),
	)

	r.Body.Close()
}

func cancelHandler(w http.ResponseWriter, value ActionValue, payload CallbackResponse) {
	updateMessage(w, "canceled")
}

func updateMessage(w http.ResponseWriter, action string) {
	// Replace original response with "Canceled" text.
	// Temporary until I figure out how to delete messages.
	m := SlackMessage{
		ResponseType:    "ephemeral",
		Text:            action,
		Username:        _config.username,
		Channel:         _channelId,
		ReplaceOriginal: true,
	}
	handleReturn(w, m)
}
