package service

import (
	"github.com/Edmond-develop/subscription-tracker/internal/database"
	"time"
)

type SubscriptionRepository interface {
	Create(sub *database.Subscription) error
	GetAll() ([]database.Subscription, error)
	GetID(id string) (*database.Subscription, error)
	Delete(id string) error
	Summary(
		start time.Time,
		end time.Time,
		serviceName string,
		userName string,
	) (int64, error)
}
