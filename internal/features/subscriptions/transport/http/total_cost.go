package subscription_transport_http

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/zosinkin/test_assignment.git/internal/core/errors"
	core_logger "github.com/zosinkin/test_assignment.git/internal/core/logger"
	core_http_response "github.com/zosinkin/test_assignment.git/internal/core/transport/http/response"
)

var queryValidator = validator.New()

func init() {
	if err := queryValidator.RegisterValidation("monthyear", monthYearValidator); err != nil {
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



// TotalCost godoc
// @Summary      Calculate total subscription cost
// @Description  Calculate total cost of subscriptions for a given period with optional filters by user ID and service name
// @Tags         subscription
// @Produce      json
// @Param        user_id       query     string  false  "User ID"
// @Param        service_name  query     string  false  "Service name"
// @Param        period_start  query     string  true   "Start period (MM-YYYY)"
// @Param        period_end    query     string  true   "End period (MM-YYYY)"
// @Success      200  {object}  TotalCostResponse
// @Failure      400  {object}  core_http_response.ErrorResponse
// @Failure      404  {object}  core_http_response.ErrorResponse
// @Failure      500  {object}  core_http_response.ErrorResponse
// @Router       /subscriptions/total [get]
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

	if err := queryValidator.Struct(request); err != nil {
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