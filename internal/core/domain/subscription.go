package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/zosinkin/test_assignment.git/internal/core/errors"
)

const (
	UninitializedID      int64 = 0
	UninitializedVersion int   = 0
)


type Subscription struct {
	ID          int64
	Version     int
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
}


func NewSubscription(
	id int64,
	version int,
	serviceName string,
	price int,
	userID uuid.UUID,
	startDate time.Time,
	endDate *time.Time,
) Subscription {
	return Subscription{
		ID:          id,
		Version:     version,
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}


func NewSubscriptionUninitialized(
	serviceName string,
	price int,
	userID uuid.UUID,
	startDate time.Time,
	endDate *time.Time,
) Subscription {
	return NewSubscription(
		UninitializedID,
		UninitializedVersion,
		serviceName,
		price,
		userID,
		startDate,
		endDate,
	)
}


func (s *Subscription) Validate() error {
	serviceNameLen := len([]rune(s.ServiceName))
	if serviceNameLen < 1 || serviceNameLen > 100 {
		return fmt.Errorf(
			"Invalid `service_name` len: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if s.Price < 0{
		return fmt.Errorf("invalid price: must be >= 0")
	}

	if s.UserID == uuid.Nil {
		return fmt.Errorf("invalid `user_id`: empty uuid")
	}

	if s.StartDate.IsZero() {
		return fmt.Errorf("invalid `start_date`")
	}

	if s.EndDate != nil && s.EndDate.Before(s.StartDate) {
		return fmt.Errorf("invalid `end_date`: before `start_date`")
	}
	return nil
}
