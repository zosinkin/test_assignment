package subscription_transport_http

import (
	"net/http"

	core_logger "github.com/zosinkin/test_assignment.git/internal/core/logger"
	core_http_response "github.com/zosinkin/test_assignment.git/internal/core/transport/http/response"
	core_http_utils "github.com/zosinkin/test_assignment.git/internal/core/utils"
)


type GetSubResponse subDTOResponse


func (h *SubHTTPHandler) GetSub(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get Subscription ID path value",
		)

		return
	}

	user, err := h.subService.GetSub(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err, 
			"failed to get subscription",
		)

		return
	}

	response := GetSubResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}