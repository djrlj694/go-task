#!/usr/bin/make -f
# =========================================================================== #
# Copyright © 2022 | All rights reserved.
# =========================================================================== #
# PROGRAM: Makefile
# PROJECT: Task
# COMPANY: djrlj694.dev
# LICENSE: MIT
#
# PURPOSE:
# - To support the phases of software project development leading to production-
# quality releases.
#
# AUTHORS:
# - Robert (Bob) L. Jones
#
# CREATED: Jan 02, 2022
# REVISED: Jan 02, 2022
# =========================================================================== #


# =========================================================================== #
# ENVIRONMENT
# =========================================================================== #


-include .env


# =========================================================================== #
# DEFAULT CONSTANTS
# =========================================================================== #


# -- Debugging & Error Capture -- #

# 0 = false, 1 = true
# VERBOSE ?= 0
VERBOSE ?= false

# -- Make -- #

# Name of the main makefile.
MAKEFILE ?= $(firstword $(MAKEFILE_LIST))

# Path of the directory containing the main makefile.
MAKEFILE_DIR ?= $(dir $(realpath $(MAKEFILE)))

# Path of the directory containing secondary makefiles.
MAKE_DIR ?= $(MAKEFILE_DIR).make/


# =========================================================================== #
# INTERNAL CONSTANTS
# =========================================================================== #


# -- Commands -- #

# Command options for no verbosity.
ifeq ($(VERBOSE),0)
Q := --quiet
S := --silent
endif

# Command options for verbosity.
ifneq ($(VERBOSE),0)
V := -v
endif

# -- Continuous Integration/Delivery (CI/CD) -- #

BUILD=`date +%FT%T%z`
VERSION := `git tag | tail -1`

# -- Go -- "

LDFLAGS := -ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"


# ============================================================================ #
# PHONY TARGETS
# ============================================================================ #


# -- Main Targets -- #

.PHONY: all build clean help install run run0 run1 run2 test

# Force the default target execution sequence to display the online help if no
# target is specified in the command line following "make".
all: help

## build: Completes all build activities.
build: build-go

## clean: Completes all cleanup activities.
clean: clean-go

## test: Completes all test activities.
test: test-go test-task


# =========================================================================== #
# SECONDARY MAKEFILES
# =========================================================================== #


# -- Feature Dependencies -- #

-include $(MAKE_DIR)features/formatting.mk
-include $(MAKE_DIR)features/helping.mk

# -- Platform Dependencies -- #

-include $(MAKE_DIR)platforms/Go.mk

# -- Test Dependencies -- #

-include $(MAKE_DIR)tests/task.mk
