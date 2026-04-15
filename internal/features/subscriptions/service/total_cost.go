package subscription_service

import (
	"context"
	"fmt"
	"time"
	core_errors "github.com/zosinkin/test_assignment.git/internal/core/errors"
)


func (s *SubService) TotalCost(
	ctx context.Context,
	userID string,
	serviceName string,
	periodStart string,
	periodEnd string,
) (int, error) {
	start, err := time.Parse("01-2006", periodStart)
	fmt.Println(start)
	if err != nil {
		return 0, fmt.Errorf("invalid period_start: %w", core_errors.ErrInvalidArgument)
	}

	end, err := time.Parse("01-2006", periodEnd)
	if err != nil {
		return 0, fmt.Errorf("invalid period_end: %w", core_errors.ErrInvalidArgument)
	}

	if end.Before(start) {
		return 0, fmt.Errorf("period_end before period_start: %w", core_errors.ErrInvalidArgument)
	}

	total, err := s.subRepository.TotalCost(
		ctx,
		userID,
		serviceName,
		start,
		end,
	)
	if err != nil {
		return 0, fmt.Errorf("repository total cost: %w", err)
	}

	return total, nil
}