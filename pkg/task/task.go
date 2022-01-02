// Copyright 2022 The Task Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Model a task.

// task: A package task provides primitives for orchestrating tasks.
package task

import (
	"bytes"
	"fmt"
	"os/exec"
	"text/template"
)

// Types

// Task: A struct that models commands to execute.
type Task struct {
	Name        string
	Description string
	Parameters  Parameters
	Commands    []string
}

// Public Methods

// Run: A method that runs a task's commands.
func (t *Task) Run(data Task) {
	for _, rawCmd := range t.Commands {
		renderedCmd := render(rawCmd, data)
		run(renderedCmd)
	}
}

// Private Methods

// render: A method that renders a template to a street.
func render(pattern string, data Task) string {
	var buffer bytes.Buffer

	tmpl, err := template.New("command").Parse(pattern)
	Check(err)

	err = tmpl.Execute(&buffer, data)
	Check(err)

	return buffer.String()
}

// run: A method that runs a command.
func run(cmdString string) {
	var out bytes.Buffer

	cmd := exec.Command("bash", "-c", cmdString)
	cmd.Stdout = &out

	err := cmd.Run()
	Check(err)

	fmt.Print(out.String())
}
