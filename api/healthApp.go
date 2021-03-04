package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/kubodin/service"
)

type healthAppApi struct {
	service service.Service
}

func (a *healthAppApi) Routes(r fiber.Router) {

	r.Get("/health", withClient(a.health))
}

// health запрос о состонии сервера приложений 1С Предприятие
//  Swagger-spec:
//		@Summary запрос о состонии сервера приложений 1С Предприятие
// 		@Description запрос о состонии сервера приложений 1С Предприятие
// 		@Tags app
// 		@Accept  json
// 		@Produce json
// 		@Success 200 {object} StatusResponse
//		@Failure 500 {object} StatusResponse
// 		@Router /app/{app}/health [get]
func (a *healthAppApi) health(client service.ClientContext, ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)

}
