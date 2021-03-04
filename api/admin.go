package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/kubodin/service"
)

type agentApi struct {
	service service.Service
}

func (a *agentApi) Routes(r fiber.Router) {

	r.Get("/agent/version", withClient(a.Version))
	r.Get("/agent/admins", withClient(a.List))
	r.Post("/agent/admins", withClient(a.RegAgentAdmin))
	r.Delete("/agent/admins/:admin", withClient(a.UnregAgentAdmin))

}

// List получение списка администраторов агента на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение списка администраторов агента на сервере 1С Предприятие
// 		@Description получение списка администраторов агента на сервере 1С Предприятие
// 		@Tags admins,agent
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.UsersList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/agent/admins [get]
func (a *agentApi) List(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)
}

// Version получение версии агента на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary получение версии агента на сервере 1С Предприятие
// 		@Description получение версии агента на сервере 1С Предприятие
// 		@Tags agent
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/agent/version [get]
func (a *agentApi) Version(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)
}

// RegAgentAdmin выполняет регистрацию нового адмнимистратор на агенте сервера 1С Предприятие
//  Swagger-spec:
//		@Summary выполняет регистрацию нового адмнимистратор на агенте сервера 1С Предприятиеи
// 		@Description выполняет регистрацию нового адмнимистратор на агенте сервера 1С Предприятие
// 		@Tags admins,agent
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param req body serialize.UserInfo true "user info"
// 		@Success 200 {object} Response{data=serialize.UserInfo}
//		@Failure 404 {object} Response
//		@Failure 500 {object} Response
// 		@Router /app/{app}/agent/admins [post]
func (a *agentApi) RegAgentAdmin(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)

}

// UnregAgentAdmin Удаление администратора агента на сервере 1С Предприятие
//  Swagger-spec:
//		@Summary Удаление администратора агента на сервере 1С Предприятие
// 		@Description Удаление администратора агента на сервере 1С Предприятие
// 		@Tags admins,agent
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param admin path string true "admin name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/agent/admins/{admin} [delete]
func (a *agentApi) UnregAgentAdmin(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)

}
