package subscriptions_postgres_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/zosinkin/test_assignment.git/internal/core/domain"
)

type SubModel struct {
	ID          int
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
}


func subDomainsFromModels(subs []SubModel) []domain.Subscription {
	subDomains := make([]domain.Subscription, len(subs))

	for i, sub := range subs {
		subDomains[i] = domain.NewSubscription(
			sub.ID,
			sub.ServiceName,
			sub.Price,
			sub.UserID,
			sub.StartDate,
			sub.EndDate,
		)
	}
	return subDomains
}