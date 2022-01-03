// ============================================================================
// Copyright © 2022 | All rights reserved.
// ============================================================================
// PROGRAM: task_test.go
// PROJECT: Task
// COMPANY: djrlj694.dev
// LICENSE: MIT
//
// PURPOSE:
// - To model a task.
//
// AUTHORS:
// - Robert (Bob) L. Jones
//
// CREATED: Jan 02, 2022
// REVISED: Jan 02, 2022
// ============================================================================

// task: Package to provide primitives for orchestrating tasks.
package task

import (
	"bytes"
	"fmt"
	"os/exec"
	"text/template"
)

// Types

// Task: Struct to model commands to execute.
type Task struct {
	Name        string
	Description string
	Parameters  Parameters
	Commands    []string
}

// Public Methods

// Run: Method to run a task's commands.
func (t *Task) Run(data Task) {
	for _, rawCmd := range t.Commands {
		renderedCmd := render(rawCmd, data)
		run(renderedCmd)
	}
}

// Private Utility Functions

// render: Function to render a template to a string.
func render(pattern string, data Task) string {
	var buffer bytes.Buffer

	tmpl, err := template.New("command").Parse(pattern)
	Check(err)

	err = tmpl.Execute(&buffer, data)
	Check(err)

	return buffer.String()
}

// run: Function to run a command.
func run(cmdString string) {
	var out bytes.Buffer

	cmd := exec.Command("bash", "-c", cmdString)
	cmd.Stdout = &out

	err := cmd.Run()
	Check(err)

	fmt.Print(out.String())
}
