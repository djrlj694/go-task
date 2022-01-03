// ============================================================================
// Copyright © 2022 | All rights reserved.
// ============================================================================
// PROGRAM: config.go
// PROJECT: Task
// COMPANY: djrlj694.dev
// LICENSE: MIT
//
// PURPOSE:
// - To model a configuration file for orchestrating tasks.
//
// AUTHORS:
// - Robert (Bob) L. Jones
//
// CREATED: Jan 02, 2022
// REVISED: Jan 02, 2022
// ============================================================================

// Model a configuration file for orchestrating tasks.

// task: Package to provide primitives for orchestrating tasks.
package task

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config: Struct to model the configuration file read by the `task` command.
type Config struct {
	Globals      map[string]string            `yaml:"globals"`
	Environments map[string]map[string]string `yaml:"environments"`
	Hosts        map[string]map[string]string `yaml:"hosts"`
	Tasks        map[string]Task              `yaml:"tasks"`
	Pipelines    map[string][]string          `yaml:"pipelines"`
}

// Utility Methods

// Read: Method to read a YAML file to instantiate a configuration.
func (c *Config) Read(path string) {
	absPath, err := filepath.Abs(os.ExpandEnv(path))
	Check(err)

	cfgContent, err := ioutil.ReadFile(absPath)
	Check(err)

	err = yaml.Unmarshal(cfgContent, &c)
	Check(err)
}
