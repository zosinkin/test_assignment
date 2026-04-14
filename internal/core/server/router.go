package core_http_server

import (
	"fmt"
	"net/http"
)

type APIVersion string

var (
	ApiVersion1 = APIVersion("v1")
	ApiVersion2 = APIVersion("v2")
	ApiVersion3 = APIVersion("v3")
)

type ApiVersionRouter struct {
	*http.ServeMux
	apiVersion  APIVersion
}

func NewAPIVersionRouter(
	apiVersion APIVersion,
) *ApiVersionRouter {
	return &ApiVersionRouter{
		ServeMux: http.NewServeMux(),
		apiVersion: apiVersion,
	}
}


func (r *ApiVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.Handler)
	}
}