// Copyright 2022 The Task Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Model a task's. parameters

// task: A package that provides primitives for orchestrating tasks.
package task

// Parameters: A struct that models parameters associated with a task.
type Parameters struct {
	Global      map[string]string
	Environment map[string]string
	Host        map[string]string
}
