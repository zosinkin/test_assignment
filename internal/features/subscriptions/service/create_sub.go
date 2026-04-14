package subscription_service

import (
	"context"
	"fmt"

	"github.com/zosinkin/test_assignment.git/internal/core/domain"
)


func (s *SubService) CreateSub(
	ctx context.Context,
	sub domain.Subscription,
) (domain.Subscription, error) {
	if err := sub.Validate(); err != nil {
		return domain.Subscription{}, fmt.Errorf("validate subscription domain: %w", err)
	}

	sub, err := s.subRepository.CreateSub(ctx, sub)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("create subscription: %w", err)
	}

	return sub, nil
}