package providers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type SlackNotificationsWebhookProvider struct {
}

type Notification struct {
	Text string `json:"text"`
}

func (s *SlackNotificationsWebhookProvider) SendNotification(message string) error {
	url := os.Getenv("SLACK_WEBHOOK_URL")
	notification := Notification{Text: message}

	jsonStr, _ := json.Marshal(notification)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return nil
}

func NewSlackNotificationsWebhookProvider() *SlackNotificationsWebhookProvider {
	return &SlackNotificationsWebhookProvider{}
}
