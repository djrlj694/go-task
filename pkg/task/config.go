// Copyright 2022 The Task Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Model a configuration file for orchestrating tasks.

// task: A package that provides primitives for orchestrating tasks.
package task

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config: A struct that models the configuration file read by the `task`
// command.
type Config struct {
	Globals      map[string]string            `yaml:"globals"`
	Environments map[string]map[string]string `yaml:"environments"`
	Hosts        map[string]map[string]string `yaml:"hosts"`
	Tasks        map[string]Task              `yaml:"tasks"`
	Pipelines    map[string][]string          `yaml:"pipelines"`
}

// Utility Methods

// Read: A method that reads a YAML file to instantiate a configuration.
func (c *Config) Read(path string) {
	absPath, err := filepath.Abs(os.ExpandEnv(path))
	Check(err)

	cfgContent, err := ioutil.ReadFile(absPath)
	Check(err)

	err = yaml.Unmarshal(cfgContent, &c)
	Check(err)
}
