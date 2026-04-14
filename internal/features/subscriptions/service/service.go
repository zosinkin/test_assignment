package subscription_service

import (
	"context"

	"github.com/zosinkin/test_assignment.git/internal/core/domain"
)


type SubService struct {
	subRepository SubRepository
}


type SubRepository interface {
	CreateSub(
		ctx context.Context,
		subscription domain.Subscription,
	) (domain.Subscription, error)

	TotalCost(
		ctx context.Context,
		userID string,
		serviceName string,
		periodStart string,
		periodEnd string,
	) (int, error)

	GetSub(
		ctx context.Context,
		ID int,
	) (domain.Subscription, error)
}


func NewSubService(
	subRepository SubRepository, 
) *SubService {
	return &SubService{
		subRepository: subRepository,
	}
}