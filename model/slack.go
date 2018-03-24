package model

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/BurntSushi/toml"
)

// Slack has webhook parameters
type Slack struct {
	Payload SlackPayload
	URL     string
}

// SlackPayload includes `payload` slack parameters
type SlackPayload struct {
	Channel   string `json:"channel"`
	UserName  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	LinkNames int    `json:"link_names"`
}

// NewSlack returns a initialized Slack struct by config.toml
func NewSlack(filePath string) (s Slack, err error) {
	_, err = toml.DecodeFile(filePath, &s)
	return
}

// Post submits post request to slack
func (s *Slack) Post() (err error) {
	b, err := json.Marshal(s.Payload)
	if err != nil {
		return
	}

	var v = url.Values{}
	v.Set("payload", string(b))

	res, err := http.PostForm(s.URL, v)
	defer res.Body.Close()
	return
}
