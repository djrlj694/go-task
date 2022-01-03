#!/usr/bin/make -f
# =========================================================================== #
# Copyright © 2022 | All rights reserved.
# =========================================================================== #
# PROGRAM: task.mk
# PROJECT: Task
# COMPANY: djrlj694.dev
# LICENSE: MIT
#
# PURPOSE:
# - To test the `task` command using a pre-determined set of CLI options.
#
# AUTHORS:
# - Robert (Bob) L. Jones
#
# CREATED: Jan 02, 2022
# REVISED: Jan 02, 2022
# =========================================================================== #


# =========================================================================== #
# PHONY TARGETS
# =========================================================================== #


# -- Prerequisites for "test" Target -- #

.PHONY: test-task test-task-help test-task-1 test-task-2

test-task: test-task-help test-task-1 test-task-2

test-task-help:
	~/go/bin/task -h

test-task-1:
	~/go/bin/task -c examples/sample.hello.yml -e test -p etl1

test-task-2:
	~/go/bin/task -c examples/sample.hello.yml -e test -p etl2
