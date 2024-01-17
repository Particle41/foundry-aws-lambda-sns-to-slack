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
	// Change this line to use the provider you want (in the future we will use a factory to create the provider)
	slackNotificationsProvider := providers.NewSlackNotificationsWebhookProvider()
	// slackNotificationsProvider := providers.NewSlackNotificationsApiProvider()
	service := services.NewAlertsService(slackNotificationsProvider)

	for _, record := range request.Records {
		println("Message", record.SNS.Message)
		service.SendNotification(record.SNS.Message)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
