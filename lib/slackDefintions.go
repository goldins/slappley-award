package lib

type ActionValue struct {
	Url     string `json:"url"`
	Text    string `json:"text"`
	Args    string `json:"args"`
	Command string `json:"command"`
}

type Action struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Style string `json:"style"`
	Value string `json:"value"`
}

type CancelAction struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Style string `json:"style"`
	Value string `json:"value"`
}

type Attachment struct {
	Title      string   `json:"title"`
	ImageUrl   string   `json:"image_url"`
	CallbackId string   `json:"callback_id"`
	Actions    []Action `json:"actions"`
}

type SlackMessage struct {
	ResponseType string       `json:"response_type"`
	Text         string       `json:"text"`
	Username     string       `json:"username"`
	Channel      string       `json:"channel"`
	Icon         string       `json:"icon_emoji"`
	Attachments  []Attachment `json:"attachments"`
}

func GetActionType(actionType string) interface{} {
	switch actionType {
	case "ActionValue":
		var p Action
		return &p
	case "CancelAction":
		var p CancelAction
		return &p
	}
	return nil
}
