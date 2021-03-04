package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/kubodin/models"
	"github.com/khorevaa/kubodin/service"
	"time"
)

type healthApi struct {
	name    string
	version string
	route   string
	service service.Service
}

func (a *healthApi) Routes(r fiber.Router) {

	r.Get("/health", a.health)
	r.Get("/health/readiness", a.readiness)
}

// health запрос о состонии приложения
//  Swagger-spec:
//		@Summary запрос о состонии приложения
// 		@Description запрос о состонии приложения
// 		@Tags health
// 		@Accept  json
// 		@Produce json
// 		@Success 200 {object} StatusResponse
//		@Failure 500 {object} StatusResponse
// 		@Router /health [get]
func (a *healthApi) health(ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)

}

// readiness запрос подробного состония приложения
//  Swagger-spec:
//		@Summary запрос подробного состония приложения
// 		@Description запрос подробного состония приложения
// 		@Tags health
// 		@Accept  json
// 		@Produce json
// 		@Success 200 {object} ReadinessCheckStatus
//		@Failure 500 {object} ReadinessCheckStatus
// 		@Router /health/readiness [get]
func (a *healthApi) readiness(ctx *fiber.Ctx) error {

	return NoAllowResponse(ctx)
}

type readinessCheck struct {
	Name    string
	Version string
	service service.Service
	Apps    []AppServiceCheckConfig
}

type ReadinessCheckStatus struct {
	Name    string             `json:"name"`
	Version string             `json:"version"`
	Status  bool               `json:"status"`
	Apps    []AppServiceStatus `json:"apps"`
}

type serviceStatus struct {
	Status bool   `json:"status"`
	Error  string `json:"errors,omitempty"`
}

type AppServiceCheckConfig struct {
	App     *models.AppServer
	Host    string
	TimeOut time.Duration `json:"timeout,omitempty"` // default value: 10
	Headers []HTTPHeader  `json:"headers,omitempty"`
	Ctx     *fiber.Ctx
}

type AppServiceStatus struct {
	Name         string  `json:"name"`
	Host         string  `json:"host"`
	Status       bool    `json:"status"`
	ResponseTime float64 `json:"response_time"`
	URL          string  `json:"url"`
	Error        string  `json:"errors,omitempty"`
}

// HTTPHeader used to setup webservices integrations
type HTTPHeader struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type StatusResponse struct {
	Status bool   `json:"status"`
	Err    string `json:"errors,omitempty"`
}

func (r StatusResponse) Error() string {
	return r.Err
}
