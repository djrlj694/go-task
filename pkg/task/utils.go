// ============================================================================
// Copyright © 2022 | All rights reserved.
// ============================================================================
// PROGRAM: task_test.go
// PROJECT: Task
// COMPANY: djrlj694.dev
// LICENSE: MIT
//
// PURPOSE:
// - To define utility functions used across functions/methods.
//
// AUTHORS:
// - Robert (Bob) L. Jones
//
// CREATED: Jan 02, 2022
// REVISED: Jan 02, 2022
// ============================================================================

// task: Package to provide primitives for orchestrating tasks.
package task

import "log"

// Check: Function to log non-nil errors and exits the progam with an error.
func Check(e error) {
	if e != nil {
		log.Fatalf("execution failed: %s", e)
	}
}
