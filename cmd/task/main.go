// ============================================================================
// Copyright © 2022 | All rights reserved.
// ============================================================================
// PROGRAM: main.go
// PROJECT: Task
// COMPANY: djrlj694.dev
// LICENSE: MIT
//
// PURPOSE:
// - To define the command `task`.
//
// AUTHORS:
// - Robert (Bob) L. Jones
//
// CREATED: Jan 02, 2022
// REVISED: Jan 02, 2022
// ============================================================================

// main: Package to define the command `task`.
package main

import (
	"os"
	"task/pkg/task"

	"github.com/urfave/cli/v2"
)

// Version: Variable modeling the version from the Go linker flags option
// `-ldflags``, set in the main makefile ("Makefile") by the command:
//
// `git describe --tags`
//
// For more info, see:
// 1. https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
// 2. https://github.com/ahmetb/govvv
// 3. https://stackoverflow.com/questions/11354518/application-auto-build-versioning
// 4. https://stackoverflow.com/questions/11354518/application-auto-build-versioning
// 5. https://steemit.com/golang/@anarcher/releasing-with-bumpversion-govvv-and-drone
var Version string

var app = cli.NewApp()

// info: Function to set application metadata used by the `cli` module.
// REFERENCE: https://itnext.io/how-to-create-your-own-cli-with-golang-3c50727ac608
func info() {
	app.Name = "task"
	app.Usage = "orchestrate commands"
	app.Authors = []*cli.Author{{
		Name:  "Robert (Bob) L. Jones",
		Email: "djrlj694@gmail.com",
	}}
	app.Version = Version
	app.Action = action
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			// Value:   "CONFIG.yaml",
			Usage: "read a YAML file",
		},
		&cli.StringFlag{
			Name:    "data",
			Aliases: []string{"d"},
			Usage:   "read a JSON string as data",
		},
		&cli.StringFlag{
			Name:    "environment",
			Aliases: []string{"e"},
			Usage:   "set the environment",
		},
		&cli.StringFlag{
			Name:    "pipeline",
			Aliases: []string{"p"},
			Usage:   "set the task pipeline",
		},
	}
}

// action: Function to define the main action used by the `cli` module.
//
// REFERENCE: https://stackoverflow.com/questions/58085039/testing-urfave-cli-based-applications-with-go
func action(context *cli.Context) error {
	var config task.Config
	config.Read(context.String("config"))

	// var data map[string]interface{}
	// err := json.Unmarshal([]byte(context.String("data")), &data)
	// task.Check(err)

	global := config.Globals
	environment := config.Environments[context.String("environment")]
	host := config.Hosts["default"]
	parameters := task.Parameters{
		Global:      global,
		Environment: environment,
		Host:        host,
	}
	pipeline := config.Pipelines[context.String("pipeline")]

	preTask := config.Tasks["pre"]
	postTask := config.Tasks["post"]

	for _, t := range pipeline {
		mainTask := config.Tasks[t]
		mainTask.Name = t
		mainTask.Parameters = parameters

		preTask.Run(mainTask)
		mainTask.Run(mainTask)
		postTask.Run(mainTask)
	}

	return nil
}

// main: Function to define the command's main function.
func main() {
	// Set application metadata for its online help.
	info()

	// Start the application.
	err := app.Run(os.Args)
	task.Check(err)

	// Completes normally (i.e., with a POSIX exit code of 0).
	os.Exit(0)
}
