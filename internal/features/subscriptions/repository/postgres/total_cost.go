package subscriptions_postgres_repository

import (
	"context"
	"fmt"
)


func (r *SubRepository) TotalCost(
	ctx 		context.Context,
	userID 		string,
	serviceName string,
	periodStart string,
	periodEnd 	string,
) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT COALESCE(SUM(price), 0)
		FROM subscription_service.subscriptions
		WHERE start_date <= TO_DATE($1, 'MM-YYYY')
		AND end_date >= TO_DATE($2, 'MM-YYYY')
	`

	

	var total int
	if err := r.pool.QueryRow(
		ctx, 
		query, 
		periodStart, 
		periodEnd,
		).Scan(&total); err != nil {
			return 0, fmt.Errorf("scan total cost: %w", err)
	}

	return total, nil
}