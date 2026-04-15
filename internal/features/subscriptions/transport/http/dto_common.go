package subscription_transport_http

import (
	"github.com/zosinkin/test_assignment.git/internal/core/domain"
	"time"
)



type subDTOResponse struct {
	ID          int    `json:"id"               example:"10"`
	ServiceName string `json:"service_name"     example:"Yandex Plus"`
	UserID      string `json:"user_id"          example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	Price       int    `json:"price"            example:"400"`
	StartDate   string `json:"start_date"       example:"07-2025"`   
	EndDate     string `json:"end_date"         example:"08-2025"`
}


func subDTOFromDomain(sub domain.Subscription) subDTOResponse {
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


func subsDTOFromDomain(subs []domain.Subscription) []subDTOResponse {
	subsDTO := make([]subDTOResponse, len(subs))

	for i, sub := range subs{
		subsDTO[i] = subDTOFromDomain(sub)
	}

	return subsDTO
}

