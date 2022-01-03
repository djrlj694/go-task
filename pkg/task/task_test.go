// ============================================================================
// Copyright © 2022 | All rights reserved.
// ============================================================================
// PROGRAM: task_test.go
// PROJECT: Task
// COMPANY: djrlj694.dev
// LICENSE: MIT
//
// PURPOSE:
// - To test the struct `Task`.
//
// AUTHORS:
// - Robert (Bob) L. Jones
//
// CREATED: Jan 02, 2022
// REVISED: Jan 02, 2022
// ============================================================================

// task: Provides primitives for orchestrating tasks.
package task

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Constructor Method Tests

// TestTask: Function to test the struct `Task`.
func TestTask(t *testing.T) {

	// An immediately-invoked function expression (IIFE) for testing if 2
	// objects are identical: actual ("got") vs. expected ("want").
	var check = func(got interface{}, want interface{}) {
		if !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	// An immediately-invoked function expression (IIFE) for standardizing a
	// test message.
	var msg = func(descr string) string {
		return fmt.Sprintf("initialize a %s", descr)
	}

	t.Run(msg("multi-command task"), func(t *testing.T) {
		cmds := []string{
			"echo -n $(date +'%F %T') \"[INFO]:\" \"Running task {{.Name}}...\"",
			"echo \"Hello, ${USER}!\"",
		}
		task := Task{
			Name:        "test1",
			Description: "A task for testing purposes",
			Commands:    cmds,
		}

		check(task.Commands[1], "echo \"Hello, ${USER}!\"")
	})

}
