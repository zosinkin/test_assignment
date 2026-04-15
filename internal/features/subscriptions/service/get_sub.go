package subscription_service

import (
	"context"
	"fmt"

	"github.com/zosinkin/test_assignment.git/internal/core/domain"
)


func (s *SubService) GetSub(
	ctx context.Context,
	id int,
) (domain.Subscription, error) {
	sub, err := s.subRepository.GetSub(ctx, id)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("get sub from repository: %w", err)
	}
	return sub, nil
}