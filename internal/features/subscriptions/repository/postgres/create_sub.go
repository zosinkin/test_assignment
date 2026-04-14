package subscriptions_postgres_repository

import (
	"context"
	"fmt"

	"github.com/zosinkin/test_assignment.git/internal/core/domain"
)


func (r *SubRepository) CreateSub(
	ctx context.Context,
	sub domain.Subscription,
) (domain.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO subscription_service.subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, service_name, price, user_id, start_date, end_date;
	`

	row := r.pool.QueryRow(
		ctx, 
		query, 
		sub.ServiceName, 
		sub.Price, 
		sub.UserID, 
		sub.StartDate, 
		sub.EndDate,
	)

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
		return domain.Subscription{}, fmt.Errorf("scan error: %w", err)	
	}

	subDomain := domain.NewSubscription(
		subModel.ID,
		subModel.ServiceName,
		subModel.Price,
		subModel.UserID,
		subModel.StartDate,
		subModel.EndDate,
	)

	return subDomain, nil
}