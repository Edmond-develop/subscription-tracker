package service

import (
	"fmt"
	"github.com/Edmond-develop/subscription-tracker/internal/database"
	"github.com/Edmond-develop/subscription-tracker/internal/repository"
	"github.com/Edmond-develop/subscription-tracker/internal/utils"
)

type SubscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(sub *database.Subscription) error {
	if sub.ServiceName == "" {
		return fmt.Errorf("service_name is required")
	}

	if sub.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	if sub.UserName == "" {
		return fmt.Errorf("user_name is required")
	}

	err := utils.ParseDates(sub.StartDate, sub.EndDate)

	if err != nil {
		return fmt.Errorf("date parse failed: %w", err)
	}
	return s.repo.Create(sub)
}

func (s *SubscriptionService) GetAll() ([]database.Subscription, error) {
	return s.repo.GetAll()
}

func (s *SubscriptionService) GetID(id string) (*database.Subscription, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	return s.repo.GetID(id)
}

func (s *SubscriptionService) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}

	return s.repo.Delete(id)
}

func (s *SubscriptionService) Summary(start, end, serviceName, userName string) (int64, error) {
	if start == "" || end == "" {
		return 0, fmt.Errorf("no start or end date provided")
	}

	startDate, err1 := utils.ParseDate(start)
	endDate, err2 := utils.ParseDate(end)

	if err1 != nil || err2 != nil {
		return 0, fmt.Errorf("date parse failed")
	}

	return s.repo.Summary(startDate, endDate, serviceName, userName)
}
