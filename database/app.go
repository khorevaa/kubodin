package database

import (
	"context"
	"github.com/khorevaa/kubodin/errors"

	"github.com/khorevaa/kubodin/models"
	"time"
)

type Repository interface {
	GetAppServers() (apps []*models.AppServer, err error)
	GetAppServer(name string) (*models.AppServer, error)
	AddAppServer(app models.AppServer) error
	DeleteAppServer(appName string) error
	Db() string
	Clear() error
}

var (
	ErrorNotFound = errors.Internal.New("app by id not found")
)

func prepareAppServer(app *models.AppServer) error {

	if len(app.Port) == 0 {
		app.Port = "1545"
	}

	if len(app.Name) == 0 {
		app.Name = app.Addr
	}

	client, err := app.Client()

	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	version, err := client.GetAgentVersion(ctx)
	app.Version = client.Version()

	if err != nil {
		// При первом  получении версии только определяется версия сервиса.
		// Надо ждать изменений в ras-client, получение версии сервиса в отдельном потоке, а не с первой папыткой открытия endpoint
		version, err = client.GetAgentVersion(ctx)

	}

	app.AgentVersion = version

	return err
}
