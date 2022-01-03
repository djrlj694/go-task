// ============================================================================
// Copyright © 2022 | All rights reserved.
// ============================================================================
// PROGRAM: config.go
// PROJECT: Task
// COMPANY: djrlj694.dev
// LICENSE: MIT
//
// PURPOSE:
// - To model a task's parameters.
//
// AUTHORS:
// - Robert (Bob) L. Jones
//
// CREATED: Jan 02, 2022
// REVISED: Jan 02, 2022
// ============================================================================

// task: Package to provide primitives for orchestrating tasks.
package task

// Parameters: Struct to model parameters associated with a task.
type Parameters struct {
	Global      map[string]string
	Environment map[string]string
	Host        map[string]string
}
