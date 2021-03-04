package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/kubodin/service"
)

type servicesApi struct {
	service service.Service
}

func (a *servicesApi) Routes(r fiber.Router) {

	r.Get("/services", withClient(a.List))
	r.Get("/services/:service", withClient(a.Info))

}

// List получение списка сервисов на кластере
//  Swagger-spec:
//		@Summary получение списка сервисов на кластере
// 		@Description получение списка сервисов на кластере
// 		@Tags services
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=[]serialize.ServiceInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/services [get]
func (a *servicesApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)
}

// List получение списка сервисов на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка сервисов на сервере 1С Предприятие
// 		@Description получение списка сервисов на сервере 1С Предприятие
// 		@Tags services
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=[]serialize.ServiceInfo}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/services [get]
func (a *servicesApi) ListApp() {}

func (a *servicesApi) Info(client service.ClientContext, ctx *fiber.Ctx) error {
	return NoAllowResponse(ctx)
}
