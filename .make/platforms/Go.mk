#!/usr/bin/make -f
# =========================================================================== #
# Copyright © 2022 | All rights reserved.
# =========================================================================== #
# PROGRAM: Go.mk
# PROJECT: Task
# COMPANY: djrlj694.dev
# LICENSE: MIT
#
# PURPOSE:
# - To facilitate Go software project activities.
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


# -- Prerequisites for "build" Target -- #

.PHONY: build-go go-format go-install go-tidy go-tidy-1-16 go-mod-tidy-17

# Completes all Go build activities.
build-go: go-tidy go-format go-install

# Completes all Go format activities.
go-format:
	go fmt ./...

# Completes all Go installation activities.
go-install:
#	govvv install $(V) -version $(VERSION) ./...
#	go install $(V) -ldflags "-w -s -X main.Version=$(VERSION)" ./...
	go install $(V) $(LDFLAGS) ./...

# Adds all Go tidy activities.
go-tidy: go.mod go-tidy-1-16 go-tidy-1-17
#	go mod tidy
#	go mod tidy -compat=`go version | { read _ _ v _; echo ${v#go}; }`

# Adds Go 1.16.x tidy activities.
go-tidy-1-16:
	go mod tidy -compat=1.16

# Adds Go 1.17.x tidy activities.
go-tidy-1-17: go.mod
	go mod tidy -compat=1.17

# -- Prerequisites for "clean" Target -- #

.PHONY: clean-go

# Completes all Go cleanup activities.
clean-go:
	go clean $(V)

# -- Prerequisites for "test" Target -- #

.PHONY: test-go go-test

# Completes all Go test activities.
test-go: go-test

# Completes a specific set of Go test activities.
go-test:
	go test $(V) ./...


# ============================================================================ #
# FILE TARGETS
# ============================================================================ #


go.mod:
	go mod init task
