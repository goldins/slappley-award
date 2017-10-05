package lib

import (
	"net/http"
	"encoding/json"
	"log"
)

func CallbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := r.PostFormValue("payload")
		var origAttachment Attachment
		json.Unmarshal([]byte(payload), &origAttachment)
		var actionValue ActionValue
		json.Unmarshal([]byte(origAttachment.Actions[0].Value), &actionValue)

		switch actionValue.Args {
		case "send":
			sendHandler(w, actionValue)
		case "cancel":
			cancelHandler(w, actionValue)
		case "shuffle":
			shuffleAction(w, actionValue)
		default:
			log.Panic("invalid action " + actionValue.Args)
		}
	}
}

func shuffleAction(w http.ResponseWriter, value ActionValue) {
	// todo: refactor FetchHandler to reuse most of its logic
}

func sendHandler(w http.ResponseWriter, value ActionValue) {
	newAttachment := Attachment{
		ImageUrl: value.Url,
		Color:    _config.messageColor,
	}

	m := SlackMessage{
		ResponseType: "in_channel",
		Text:         value.Text,
		Username:     _config.username,
		Channel:      _channelId,
		Icon:         ":d20:",
		Attachments:  []Attachment{newAttachment},
	}
	handleReturn(w, m)
}

func cancelHandler(w http.ResponseWriter, value ActionValue) {
	m := SlackMessage{
		ResponseType: "ephemeral",
		Text:         "Canceled",
		Username:     _config.username,
		Channel:      _channelId,
		Icon:         ":d20:",
	}
	handleReturn(w, m)
}
