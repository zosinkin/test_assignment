package subscriptions_postgres_repository

import (
	"context"
	"fmt"
	"github.com/zosinkin/test_assignment.git/internal/core/domain"
)


func (r *SubRepository) GetSubs(
	ctx context.Context,
) ([]domain.Subscription, error) {
	ctx, cancel := context.WithTimeout(
		ctx,
		r.pool.OpTimeout(),
	)
	defer cancel()

	query := `
		SELECT id, service_name, price, user_id, start_date, end_date
		FROM subscription_service.subscriptions
	`

	rows, err := r.pool.Query(
		ctx,
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("select subs: %w", err)
	}
	//высвобождаем подключение, которое нам было дано из connection пула
	defer rows.Close()

	var subModels []SubModel
	for rows.Next() {
		var subModel SubModel

			
		err := rows.Scan(
			&subModel.ID,
			&subModel.ServiceName,
			&subModel.Price,
			&subModel.UserID,
			&subModel.StartDate,
			&subModel.EndDate,
		)
		if err != nil {
			return nil, fmt.Errorf("scan subs: %w", err)
		}

		subModels = append(subModels, subModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	subDomains := subDomainsFromModels(subModels)

	return subDomains, nil

}