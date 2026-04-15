package subscriptions_postgres_repository

import (
	"context"
	"fmt"
	"time"
)


func (r *SubRepository) TotalCost(
	ctx 		context.Context,
	userID 		string,
	serviceName string,
	periodStart time.Time,
	periodEnd 	time.Time,
) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT COALESCE(SUM(price), 0) AS total_price
		FROM subscription_service.subscriptions
		WHERE user_id = $1
  		AND service_name = $2
  		AND start_date <= $3
  		AND end_date >= $4;
	`

	
	var total int
	if err := r.pool.QueryRow(
		ctx,
		query, 
		userID,
		serviceName, 
		periodStart, 
		periodEnd,
		
		).Scan(&total); err != nil {
			return 0, fmt.Errorf("scan total cost: %w", err)
	}

	return total, nil
}