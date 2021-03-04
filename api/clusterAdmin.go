package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/kubodin/service"
)

type clusterAdminApi struct {
	service service.Service
}

func (a *clusterAdminApi) Routes(r fiber.Router) {

	r.Get("/admins", withClient(a.List))
	r.Post("/admins", withClient(a.Create))
	r.Delete("/admins/:admin", withClient(a.Delete))

}

// List получение списка администраторов кластера
//  Swagger-spec:
//		@Summary получение списка администраторов кластера
// 		@Description получение списка администраторов кластера
// 		@Tags admins,clusters
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=serialize.UsersList}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/cluster/{cluster}/admins [get]
func (a *clusterAdminApi) List(client service.ClientContext, ctx *fiber.Ctx) error {
	return NoAllowResponse(ctx)
}

// Create выполняет регистрацию нового администратор на кластере
//  Swagger-spec:
//		@Summary выполняет регистрацию нового администратор на кластере
// 		@Description выполняет регистрацию нового администратор на кластере
// 		@Tags admins,clusters
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param req body serialize.UserInfo true "user info"
// 		@Success 200 {object} Response{data=serialize.UserInfo}
//		@Failure 404 {object} Response
//		@Failure 500 {object} Response
// 		@Router /app/{app}/cluster/{cluster}/admins [post]
func (a *clusterAdminApi) Create(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)

}

// Delete Удаление администратора агента на кластере
//  Swagger-spec:
//		@Summary Удаление администратора агента на кластере
// 		@Description Удаление администратора агента на кластере
// 		@Tags admins,clusters
// 		@Accept  json
// 		@Produce json
// 		@Param app path string true "app name"
// 		@Param cluster path string true "cluster uuid"
// 		@Param admin path string true "admin name"
// 		@Param force query bool false "force update ignore cache"
// 		@Success 200 {object} Response{data=string}
// 		@Failure 500 {object} Response
// 		@Router /app/{app}/cluster/{cluster}/admins/{admin} [delete]
func (a *clusterAdminApi) Delete(client service.ClientContext, ctx *fiber.Ctx) error {
	return NoAllowResponse(ctx)
}
