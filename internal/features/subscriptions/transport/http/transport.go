package subscription_transport_http

import (
	"context"
	"net/http"

	"github.com/zosinkin/test_assignment.git/internal/core/domain"
	core_http_server "github.com/zosinkin/test_assignment.git/internal/core/server"
)



type SubHTTPHandler struct{
	subService SubService 
}

type SubService interface {
	CreateSub(
		ctx context.Context,
		subscription domain.Subscription,
	) (domain.Subscription, error)

	TotalCost(
		ctx context.Context,
		userID string,
		serviceName string,
		periodStart string,
		periodEnd string,
	) (int, error)

	GetSub(
		ctx context.Context, 
		ID int,
	) (domain.Subscription, error)
}


func NewSubHTTPHandler(
	subService SubService,
) *SubHTTPHandler {
	return &SubHTTPHandler{
		subService: subService,
	}
}


func (h *SubHTTPHandler) Routes() []core_http_server.Route  {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path: "/subscriptions",
			Handler: h.CreateUser,
		
		},
		{
			Method: http.MethodGet,
			Path: "/subscriptions/total",
			Handler: h.TotalCost,
		},
		{
			Method: http.MethodGet,
			Path: "/subscriptions/{id}",
			Handler: h.GetSub,
		},
	}
}