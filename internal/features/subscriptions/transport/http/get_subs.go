package subscription_transport_http

import (
	"net/http"

	core_logger "github.com/zosinkin/test_assignment.git/internal/core/logger"
	core_http_response "github.com/zosinkin/test_assignment.git/internal/core/transport/http/response"
)

type GetSubsResponse []subDTOResponse




// GetSubs      godoc
// @Summary      Список всех подписок
// @Description  Просмотр списка всех подписок в базе
// @Tags         subscriptions
// @Produce      json
// @Success      200 {object} GetSubsResponse                 "Успешное получение списка подписок"
// @Failure      400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure      500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router       /subscriptions [get]
func (h *SubHTTPHandler) GetSubs(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	subDomains, err := h.subService.GetSubs(ctx)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get subs",
		)

		return
	}

	response := GetSubsResponse(subsDTOFromDomain(subDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}
