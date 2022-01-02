#!/usr/bin/make -f
# =========================================================================== #
# Copyright © 2022 | All rights reserved.
# =========================================================================== #
# PROGRAM: helping.mk
# PROJECT: Task
# COMPANY: djrlj694.dev
# LICENSE: MIT
#
# PURPOSE:
# - To enable the generation & display of online help.
#
# AUTHORS:
# - Robert (Bob) L. Jones
#
# CREATED: Jan 02, 2022
# REVISED: Jan 02, 2022
# =========================================================================== #


# =========================================================================== #
# INTERNAL CONSTANTS
# =========================================================================== #


# -- Commands -- #

# A variable for sharing the online help for both:
# 1. The "make" command;
# 2. Any wrapper script (without its parent directory or file extension) around
#    the "make" command.
MAKE_BASE := $(basename $(notdir $(MAKE)))

# -- Help Strings -- #

# Argument syntax for the "make" command.
MAKE_ARGS := [$(FG_CYAN)<target>$(RESET)]


# =========================================================================== #
# INTERNAL VARIABLES
# =========================================================================== #


# -- Help Strings -- #

# "Targets" section line item of the "make" command's online help.
target_help = $(FG_CYAN)%-8s$(RESET) %s


# =========================================================================== #
# MACROS
# =========================================================================== #


# -- Help Strings -- #

# "Targets" section header of the "make" command's online help.
define targets_help

Targets:

endef
export targets_help

# "Usage" section of the "make" command's online help.
define usage_help

Usage:
  make $(MAKE_ARGS)

endef
export usage_help


# =========================================================================== #
# PHONY TARGETS
# =========================================================================== #


# -- Main Targets -- #

.PHONY: help

## help: Shows the "make" command's online help.
help:
	@printf "$$usage_help"
	@printf "$$targets_help"
# Use the makefile set as a data source for displaying a lexicographically
# sorted, color-formatted list of targets.
#
# Note:
# 1. "cat" handles $(MAKEFILE_LIST), a space-delimited list of makefiles that
#    includes Makefiles and makefiles with the extension ".mk".
# 2. "egrep" filters for makefile lines with:
#    a. "## " starting at column 1;
#    b. The name of of the target preceding a colon, followed by a space.
	@cat $(MAKEFILE_LIST) | \
	egrep '^## [a-zA-Z_-]+: ' | \
	sed -e 's/## //' | sort -t ':' -k1,1 -d | \
	awk -F: '{printf "  $(target_help)\n", $$1, $$2}'
	@echo ""
