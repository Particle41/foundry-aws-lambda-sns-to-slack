package main

import (
	"hello-world/providers"
	"hello-world/services"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.SNSEvent) error {

	if request.Records == nil {
		return nil
	}

	slackNotificationsProvider := providers.NewSlackNotificationsProvider()
	service := services.NewAlertsService(slackNotificationsProvider)

	for _, record := range request.Records {
		println(record.SNS.Message)
		service.SendNotification(record.SNS.Message)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
