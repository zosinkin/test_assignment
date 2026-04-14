package subscriptions_postgres_repository

import (
	"time"
	"github.com/google/uuid"
)

type SubModel struct {
	ID          int
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
}

