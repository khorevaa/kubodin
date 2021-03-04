package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/kubodin/service"
)

type connectionsApi struct {
	service service.Service
}

func (a *connectionsApi) Routes(r fiber.Router) {

	r.Get("/connections", withClient(a.List))
	r.Post("/connections/terminate", withClient(a.Terminate))
	r.Delete("/connections/:connection.:process", withClient(a.TerminateOne))

}

// List получение списка подключений на кластере
//  Swagger-spec:
//		@Summary получение списка подключений на кластере
// 		@Description пполучение списка подключений на кластере
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=serialize.ConnectionShortInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/connections [get]
func (a *connectionsApi) List(client service.ClientContext, ctx *fiber.Ctx) error {
	return NoAllowResponse(ctx)
}

// List получение списка подключений для информационной базы на кластере
//  Swagger-spec:
//		@Summary получение списка подключений для информационной базы на кластере
// 		@Description получение списка подключений для информационной базы на кластере
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
//		@Success 200 {object} Response{data=serialize.ConnectionShortInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase}/connections [get]
func (a *connectionsApi) ListClusterInfobase() {}

// List получение списка подключений на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка подключений на сервере 1С Предприятие
// 		@Description получение списка подключений на сервере 1С Предприятие
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Param force query bool false "force update ignore cache"
//		@Success 200 {object} Response{data=serialize.ConnectionShortInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/connections [get]
func (a *connectionsApi) ListApp() {}

// List получение списка подключений для информационной базы на сервер 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка подключений для информационной базы на сервер 1С Предприятие
// 		@Description получение списка подключений для информационной базы на сервер 1С Предприятие
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
// 		@Param force query bool false "force update ignore cache"
//		@Success 200 {object} Response{data=serialize.ConnectionShortInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/infobases/{infobase}/connections [get]
func (a *connectionsApi) ListAppInfobase() {}

// TerminateOne отключение подключения на кластере
//  Swagger-spec:
//		@Summary отключение подключения на кластере
// 		@Description отключение подключения на кластере
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param connection path string true "connection uuid"
// 		@Param process path string true "process uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=models.TerminateConnectionSig}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/connections/{connection}.{process} [delete]
func (a *connectionsApi) TerminateOne(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)
}

// TerminateOne отключение подключения на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary отключение подключения на сервере 1С Предприятие
// 		@Description отключение подключения на сервере 1С Предприятие
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param connection path string true "connection uuid"
// 		@Param process path string true "process uuid"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
// 		@Success 200 {object} Response{data=models.TerminateConnectionSig}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/connections/{connection}.{process} [delete]
func (a *connectionsApi) TerminateOneApp() {}

// Terminate отключение списка подключений на кластере
//  Swagger-spec:
//		@Summary отключение списка подключений на кластере
// 		@Description отключение списка подключений или по информационной базе на кластере
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param req body models.TerminateConnectionsRequest true "request"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
//		@Success 200 {object} Response{data=models.TerminateConnectionsResponse}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/connections/terminate [post]
func (a *connectionsApi) Terminate(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)

}

// Terminate отключение списка подключений на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary отключение списка подключений на сервере 1С Предприятие
// 		@Description отключение списка подключений или по информационной базе на сервере 1С Предприятие
// 		@Tags connections
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param req body models.TerminateConnectionsRequest true "request"
// 		@Param cluster-id query string false "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param infobase-usr query string false "infobase user"
// 		@Param infobase-pwd query string false "infobase password"
//		@Success 200 {object} Response{data=models.TerminateConnectionsResponse}
// 		@Failure 400 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/connections/terminate [post]
func (a *connectionsApi) TerminateApp() {}
