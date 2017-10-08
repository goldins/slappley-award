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
	Value string `json:"value"` // JSONified representation of ActionValue
}

type ReturnedAction struct {
	Name  string      `json:"name"`
	Text  string      `json:"text"`
	Type  string      `json:"type"`
	Style string      `json:"style"`
	Value ActionValue `json:"value"`
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
	Color      string   `json:"color"`
	Actions    []Action `json:"actions"`
}

type SlackMessage struct {
	ResponseType    string       `json:"response_type"`
	Text            string       `json:"text"`
	Username        string       `json:"username"`
	Channel         string       `json:"channel"`
	Icon            string       `json:"icon_emoji"`
	Attachments     []Attachment `json:"attachments"`
	ReplaceOriginal bool         `json:"replace_original"`
}

type Channel struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type CallbackResponse struct {
	Actions      []Action `json:"actions"`
	CallbackId   string   `json:"callback_id"`
	ActionTs     string   `json:"action_ts"`
	MessageTs    string   `json:"message_ts"`
	AttachmentId string   `json:"attachment_id"`
	Token        string   `json:"token"`
	Type         string   `json:"type"`
	ResponseUrl  string   `json:"response_url"`
	TriggerId    string   `json:"trigger_id"`
	Channel      Channel  `json:"channel"`
}

type ChatUpdateMessage struct {
	Token       string       `json:"token"`
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	MessageTs   string       `json:"message_ts"`
	Attachments []Attachment `json:"attachments"`
}

type ChatDeleteMessage struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
	AsUser  bool   `json:"as_user"`
}

type ChatUpdateResponse struct {
	Ok      bool   `json:"ok"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
	Text    string `json:"text"`
}
