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

//dto
type CreateSubRequest struct {
	ServiceName string `json:"service_name" validate:"required,min=1,max=100"`
	UserID      string `json:"user_id" validate:"required,uuid"`
	Price       int    `json:"price" validate:"required,gte=0"`
	StartDate   string `json:"start_date" validate:"required,monthyear"`
	EndDate     string `json:"end_date" validate:"omitempty,monthyear"`
}


type CreateSubResponse subDTOResponse


func (h *SubHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)


	var request CreateSubRequest

	//декодируем и валидируем запрос
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate CreateUser request")

		return 
	}

	//создаем домен из dto
	userDomain, err := domainFromDTO(request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to map request to domain")
		return
	}

	//обращаемся к сервису
	userDomain, err = h.subService.CreateSub(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return 
	}

	//формирует ответ
	response := CreateSubResponse(userDTOFromDomain(userDomain))
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