package repository

import (
	"database/sql"
	"fmt"
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

type subscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (s *subscriptionRepository) Create(sub *database.Subscription) error {
	err := s.db.QueryRow(
		`INSERT INTO subscriptions (service_name, price, user_name, start_date, end_date) 
			   VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		sub.ServiceName, sub.Price, sub.UserName, sub.StartDate, sub.EndDate).Scan(&sub.ID)
	return err
}

func (s *subscriptionRepository) GetAll() ([]database.Subscription, error) {
	query := `SELECT id, service_name, price, user_name, start_date, end_date FROM subscriptions`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []database.Subscription

	for rows.Next() {
		var sub database.Subscription
		err = rows.Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserName, &sub.StartDate, &sub.EndDate)

		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, sub)
	}
	return subscriptions, nil
}

func (s *subscriptionRepository) GetID(id string) (*database.Subscription, error) {
	query := `SELECT id, service_name, price, user_name, start_date, end_date FROM subscriptions WHERE id = $1`

	row := s.db.QueryRow(query, id)

	var sub database.Subscription

	err := row.Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserName, &sub.StartDate, &sub.EndDate)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("subscription not found", err)
	}

	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (s *subscriptionRepository) Delete(id string) error {
	query := `DELETE FROM subscriptions WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *subscriptionRepository) Summary(start time.Time, end time.Time, serviceName string, userName string) (int64, error) {
	query := `SELECT SUM(price) FROM subscriptions WHERE start_date <= $2 AND (end_date IS NULL OR end_date >= $1)`

	args := []interface{}{start, end}
	argPos := 3

	if serviceName != "" {
		query += fmt.Sprintf(" AND service_name = $%d", argPos)
		args = append(args, serviceName)
		argPos++
	}

	if userName != "" {
		query += fmt.Sprintf(" AND user_name = $%d", argPos)
		args = append(args, userName)
		argPos++
	}

	var total sql.NullInt64
	err := s.db.QueryRow(query, args...).Scan(&total)

	return total.Int64, err
}
