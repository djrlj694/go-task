// Copyright 2022 The Task Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Utility functions used across functions/methods.

// task: A package that provides primitives for orchestrating tasks.
package task

import "log"

// Check: A function that logs non-nil errors and exits the progam with an
// error.
func Check(e error) {
	if e != nil {
		log.Fatalf("execution failed: %s", e)
	}
}
