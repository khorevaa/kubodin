package main

import (
	"fmt"
	"github.com/khorevaa/kubodin/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

// main
// @title KUBOdin: Remote Administration for 1S.Enterprise Application Servers
// @version 1.0
// @description KUBOdin Swagger UI
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email khorevaa@yandex.ru
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /api/v1
func main() {

	mainCmd := &cmd.MainCommand{
		Version: version,
		BuildBy: builtBy,
		Date:    date,
	}

	app := &cli.App{
		Name:    "kubodin",
		Version: buildVersion(),
		Authors: []*cli.Author{
			{
				Name: "Aleksey Khorev",
			},
		},
		Usage:       "Start API server for kubernetes & 1C.Enterprise",
		Copyright:   "(c) 2021 Khorevaa",
		Description: "API server for Kubernetes & 1C.Enterprise",
		Flags:       mainCmd.Flags(),
		Action:      mainCmd.Run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func buildVersion() string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}

	return result
}
