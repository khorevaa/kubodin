package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/kubodin/service"
)

type sessionsApi struct {
	service service.Service
}

func (a *sessionsApi) Routes(r fiber.Router) {

	r.Get("/sessions", withClient(a.List))
	//r.Get("/sessions/:session", withClient(a.Info))
	r.Delete("/sessions/:session", withClient(a.TerminateOne))

	r.Post("/sessions/terminate", withClient(a.Terminate))

}

// List получение списка сессий на кластере
//  Swagger-spec:
//		@Summary получение списка сессий на кластере
// 		@Description получение списка сессий на кластере
// 		@Tags sessions
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.SessionInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/sessions [get]
func (a *sessionsApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)

}

// List получение списка сессий информационной базы на кластере
//  Swagger-spec:
//		@Summary получение списка сессий информационной базы на кластере
// 		@Description получение списка сессий информационной базы на кластере
// 		@Tags sessions
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.SessionInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/infobases/{infobase}/sessions [get]
func (a *sessionsApi) ListInfobase() {}

// List получение списка сессий информационной базы на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка сессий информационной базы на сервере 1С Предприятие
// 		@Description получение списка сессий информационной базы на сервере 1С Предприятие
// 		@Tags sessions
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param infobase path string true "infobase uuid or name"
// 		@Param cluster-id query string false "cluster uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.SessionInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/infobases/{infobase}/sessions [get]
func (a *sessionsApi) ListAppInfobase() {}

// List получение списка сессий на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка сессий на сервере 1С Предприятие
// 		@Description получение списка сессий  на сервере 1С Предприятие
// 		@Tags sessions
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster-id query string false "cluster uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.SessionInfoList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/sessions [get]
func (a *sessionsApi) ListApp() {}

func (a *sessionsApi) Info(client service.ClientContext, ctx *fiber.Ctx) error {
	return NoAllowResponse(ctx)

}

func (a *sessionsApi) Terminate(client service.ClientContext, ctx *fiber.Ctx) error {
	return NoAllowResponse(ctx)
}

// TerminateOne отключение сессии на кластере
//  Swagger-spec:
//		@Summary отключение сессии на кластере
// 		@Description отключение сессии на кластере
// 		@Tags sessions
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param session path string true "session uuid"
// 		@Param msg query string false "message to user"
// 		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/clusters/{cluster}/sessions/{session} [delete]
func (a *sessionsApi) TerminateOne(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)
}

// TerminateOne отключение сессии на сервер 1С Предприятие
//  Swagger-spec:
//		@Summary отключение сессии на сервер 1С Предприятие
// 		@Description отключение сессии на сервер 1С Предприятие
// 		@Tags sessions
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param session path string true "session uuid"
// 		@Param msg query string false "message to user"
// 		@Param cluster-id query string false "cluster uuid"
//		@Param cluster-usr query string false "cluster user"
// 		@Param cluster-pwd query string false "cluster password"
//		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/sessions/{session} [delete]
func (a *sessionsApi) TerminateOneApp() {}
