package subscription_transport_http

import (
	"net/http"
	"regexp"
	"github.com/google/uuid"
	"time"
	"github.com/go-playground/validator/v10"
	"github.com/zosinkin/test_assignment.git/internal/core/domain"
	core_logger "github.com/zosinkin/test_assignment.git/internal/core/logger"
	core_http_request "github.com/zosinkin/test_assignment.git/internal/core/transport/http/request"
	core_http_response "github.com/zosinkin/test_assignment.git/internal/core/transport/http/response"
)


func monthYearValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	re := regexp.MustCompile(`^(0[1-9]|1[0-2])-\d{4}$`)
	return re.MatchString(value)
}

func init() {
	if err := core_http_request.RegisterValidation("monthyear", monthYearValidator); err != nil {
		panic(err)
	}
}

//DTO
type CreateSubRequest struct {
	ServiceName string `json:"service_name" validate:"required,min=1,max=100"   example:"Yandex Plus"`
	UserID      string `json:"user_id" validate:"required,uuid"                 example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	Price       int    `json:"price" validate:"required,gte=0"                  example:"400"`
	StartDate   string `json:"start_date" validate:"required,monthyear"         example:"07-2025"`
	EndDate     string `json:"end_date" validate:"omitempty,monthyear"          example:"08-2025"`
}


type CreateSubResponse subDTOResponse

// CreateSub   	godoc
// @Summary     Добавить подписку
// @Description Создать новую подписку в системе
// @Tags        subscription 
// @Accept      json
// @Produce     json
// @Param       request body     CreateSubRequest  true "CreateSub тело запроса"
// @Success     201     {object} CreateSubResponse "Успешно добавленная подписка"
// @Failure     400     {object} core_http_response.ErrorResponse "Bad request"
// @Failure     500     {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /subscriptions [post]
func (h *SubHTTPHandler) CreateSub(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)


	var request CreateSubRequest

	//декодируем и валидируем запрос
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate CreateSub request")

		return 
	}

	//создаем домен из DTO
	subDomain, err := domainFromDTO(request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to map request to domain")
		return
	}

	//обращаемся к сервису
	subDomain, err = h.subService.CreateSub(ctx, subDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create sub")
		return 
	}

	//формирует ответ
	response := CreateSubResponse(subDTOFromDomain(subDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)
}


func domainFromDTO(dto CreateSubRequest) (domain.Subscription, error) {
    userID, err := uuid.Parse(dto.UserID)
    if err != nil {
        return domain.Subscription{}, err
    }

    startDate, err := time.Parse("01-2006", dto.StartDate)
    if err != nil {
        return domain.Subscription{}, err
    }

    var endDate *time.Time
    if dto.EndDate != "" {
        t, err := time.Parse("01-2006", dto.EndDate)
        if err != nil {
            return domain.Subscription{}, err
        }
        endDate = &t
    } else {
        t := startDate.AddDate(0, 1, 0)
        endDate = &t
    }

    return domain.NewSubUninitialized(
        dto.ServiceName,
        dto.Price,
        userID,
        startDate,
        endDate,
    ), nil
}