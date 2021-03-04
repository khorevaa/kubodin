package main

import (
	"github.com/khorevaa/go-app-template/cmd"
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

func main() {

	app := &cli.App{
		Name:    "go-app-template",
		Version: version,
		Authors: []*cli.Author{
			{
				Name: "Aleksey Khorev",
			},
		},
		Usage:     "Description for go-app-template",
		Copyright: "(c) 2021 Khorevaa",
		//Description: "Command line utilities for server 1S.Enterprise",
	}

	for _, command := range cmd.Commands {
		app.Commands = append(app.Commands, command.Cmd())
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
