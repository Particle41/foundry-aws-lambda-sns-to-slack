package services

type IAlertsService interface {
	SendNotification(message string) error
}

type IAlertsProvider interface {
	SendNotification(message string) error
}

type AlertsService struct {
	Provider IAlertsProvider
}

func (a *AlertsService) SendNotification(message string) error {
	return a.Provider.SendNotification(message)
}

func NewAlertsService(provider IAlertsProvider) *AlertsService {
	return &AlertsService{
		Provider: provider,
	}
}
