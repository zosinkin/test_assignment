package subscription_transport_http

import (
	"github.com/zosinkin/test_assignment.git/internal/core/domain"
	"time"
)



type subDTOResponse struct {
	ID          int    `json:"id"`
	ServiceName string `json:"service_name"`
	UserID      string `json:"user_id"`
	Price       int    `json:"price"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}


func userDTOFromDomain(sub domain.Subscription) subDTOResponse {
    var endDate string
    if sub.EndDate != nil {
        endDate = sub.EndDate.Format(time.RFC3339)
    }

    return subDTOResponse{
        ID:          int(sub.ID),
        ServiceName: sub.ServiceName,
        UserID:      sub.UserID.String(),
        Price:       sub.Price,
        StartDate:   sub.StartDate.Format(time.RFC3339),
        EndDate:     endDate,
    }
}