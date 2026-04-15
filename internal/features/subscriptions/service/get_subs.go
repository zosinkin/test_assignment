package subscription_service

import (
	"context"
	"fmt"
	"github.com/zosinkin/test_assignment.git/internal/core/domain"
)

func (s *SubService) GetSubs(
	ctx context.Context,
) ([]domain.Subscription, error) {
	subs, err := s.subRepository.GetSubs(
		ctx,
	)
	if err != nil {
		return nil, fmt.Errorf("get subs from repository: %w", err)
	}

	return subs, nil
}