package providers

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
)

type SlackNotificationsProvider struct {
}

type Slack struct {
	SlackClient *slack.Client
	ChannelID   string
}

func (s *Slack) sendMessage(message string) error {
	ChannelID, timestamp, err := s.SlackClient.PostMessage(s.ChannelID, slack.MsgOptionText(message, false))
	if err != nil {
		return err
	}
	log.Info().Str("message", message).Msgf("Message sent successfully to %s channel at %s", ChannelID, timestamp)
	return nil
}

func initSlack() *Slack {
	return &Slack{
		SlackClient: slack.New(os.Getenv("SLACK_TOKEN")),
		ChannelID:   os.Getenv("SLACK_CHANNEL_ID"),
	}
}

func (s *SlackNotificationsProvider) SendNotification(message string) error {
	slack := initSlack()
	return slack.sendMessage(message)
}

func NewSlackNotificationsProvider() *SlackNotificationsProvider {
	return &SlackNotificationsProvider{}
}
