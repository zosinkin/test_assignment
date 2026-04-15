package subscription_transport_http

import (
	"net/http"

	core_logger "github.com/zosinkin/test_assignment.git/internal/core/logger"
	core_http_response "github.com/zosinkin/test_assignment.git/internal/core/transport/http/response"
	core_http_utils "github.com/zosinkin/test_assignment.git/internal/core/utils"
)


type GetSubResponse subDTOResponse


// GetSub godoc
// @Summary      Получить подписку
// @Description  Получить подписку по ID
// @Tags         subscription
// @Produce      json
// @Param        id   path      int   true  "ID подписки"
// @Success      200  {object}  GetSubResponse
// @Failure      400  {object}  core_http_response.ErrorResponse
// @Failure      404  {object}  core_http_response.ErrorResponse
// @Failure      500  {object}  core_http_response.ErrorResponse
// @Router       /subscriptions/{id} [get]
func (h *SubHTTPHandler) GetSub(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	subID, err := core_http_utils.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get Subscription ID path value",
		)

		return
	}

	sub, err := h.subService.GetSub(ctx, subID)
	if err != nil {
		responseHandler.ErrorResponse(
			err, 
			"failed to get subscription",
		)

		return
	}

	response := GetSubResponse(subDTOFromDomain(sub))

	responseHandler.JSONResponse(response, http.StatusOK)
}