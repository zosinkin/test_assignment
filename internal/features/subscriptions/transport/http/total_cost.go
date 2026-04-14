package subscription_transport_http

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/zosinkin/test_assignment.git/internal/core/errors"
	core_logger "github.com/zosinkin/test_assignment.git/internal/core/logger"
	core_http_response "github.com/zosinkin/test_assignment.git/internal/core/transport/http/response"
)

var queyValidator = validator.New()

func init() {
	if err := queyValidator.RegisterValidation("monthyear", monthYearValidator); err != nil {
		panic(err)
	}
}

type TotalCostRequest struct {
	UserID      string `json:"user_id" validate:"omitempty,uuid"`
	ServiceName string `json:"service_name" validate:"omitempty,min=1,max=100"`
	PeriodStart string `json:"period_start" validate:"required,monthyear"`
	PeriodEnd   string `json:"period_end" validate:"required,monthyear"`
}


type TotalCostResponse struct {
	TotalCost int `json:"total_cost"`
}

func (h *SubHTTPHandler) TotalCost(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	request := TotalCostRequest{
		UserID:   			r.URL.Query().Get("user_id"),
		ServiceName: 		r.URL.Query().Get("service_name"),
		PeriodStart:        r.URL.Query().Get("period_start"),
		PeriodEnd:        r.URL.Query().Get("period_end"),

	}

	if err := queyValidator.Struct(request); err != nil {
		responseHandler.ErrorResponse(
			fmt.Errorf("validate qury params: %w: %w", err, core_errors.ErrInvalidArgument),
			"failed to validate total cost request",
		)
		return
	}

	total, err := h.subService.TotalCost(
		ctx,
		request.UserID,
		request.ServiceName,
		request.PeriodStart,
		request.PeriodEnd,
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to calculate total cost")
		return 
	}

	responseHandler.JSONResponse(
		TotalCostResponse{TotalCost: total},
		http.StatusOK,
	)


}