package subscriptions_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/zosinkin/test_assignment.git/internal/core/domain"
	core_errors "github.com/zosinkin/test_assignment.git/internal/core/errors"
	core_postgres_pool "github.com/zosinkin/test_assignment.git/internal/core/repository/postgres/pool"
)


func (r *SubRepository) GetSub(
	ctx context.Context,
	id int,
) (domain.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, service_name, price, user_id, start_date, end_date
		FROM subscription_service.subscriptions
		WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var subModel SubModel

	err := row.Scan(
		&subModel.ID,
		&subModel.ServiceName,
		&subModel.Price,
		&subModel.UserID,
		&subModel.StartDate,
		&subModel.EndDate,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Subscription{}, fmt.Errorf(
				"user with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.Subscription{}, fmt.Errorf("scan error: %w", err)
	}

	subDomain := domain.Subscription{
    	ID: subModel.ID,
    	ServiceName: subModel.ServiceName,
    	Price: subModel.Price,
    	UserID: subModel.UserID,
    	StartDate: subModel.StartDate,
    	EndDate: subModel.EndDate,
	}

	return subDomain, nil
}