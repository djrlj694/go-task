# Task

A command and Go package for orchestrating configurable, parameterized pipelines
of tasks.

## Introduction

Task is a software project that is premised on the possibility of easier
systems programming. It seeks to solve several persistent problems plaguing
typical system programming projects:

1. Hard-coded settings
2. Insufficient reusability
3. Opaque workflows
4. Overengineering

These anti-patterns are especially apparent in shell scripts. To make this point
clearer, consider the following Bash shell code block from a data pipeline
application that loads data from a flat file to an Oracle database and then
processes that data.

```bash
echo -n "Extracting data..." >> "${HOME}/data_pipeline_01/logs/elt-${YYYYMMDD}.log"
wget -c http://geolite.maxmind.com/download/geoip/database/GeoLite2-Country.tar.gz -O - | ${TAR} -xz -C "${HOME}/data_pipeline_01/data/incoming/GeoLite2-Country-${YYYYMMDD}.csv"
[ $? -eq 0 ] echo "done." || echo "fail."

echo -n "Loading data..." >> "${HOME}/data_pipeline_01/logs/load.log"
sqlldr userid=$(cat "${HOME}/data_pipeline_01/.oracle_db.txt"), parfile="${HOME}/data_pipeline_01/etc/oracle_db.par", data="${HOME}/data_pipeline_01/data/incoming/GeoLite2-Country-${YYYYMMDD}.csv"
[ $? -eq 0 ] echo "done." || echo "fail."

echo -n "Transforming data..." >> "${HOME}/data_pipeline_01/logs/load.log"
sqlplus $(cat "${HOME}/data_pipeline_01/.oracle_db.txt") @"${HOME}/data_pipeline_01/src/oracle_etl.sql"
[ $? -eq 0 ] echo "done." || echo "fail."
```

Let's look at this program from the perspective of our problems list:

1. Hard-coded settings
2. Insufficient reusability
3. Opaque workflows

Task seeks to remedy these problems via [YAML][YAML]-formatted configuration files.
These files specify sets of sequentially run commands inside abstractions called
*tasks*. Task configuration files also define two other abstractions: hosts and environments. These abstractions respectively define host-dependent and environment-dependent parameters for a task's command sequence.

Task configuration files resemble those used by continuous integration/delivery
[(CI/CD)][CICD] services. Two key differences between Task and CI/CD
services, however, are that Task has:

1. A more minimalistic parameterization model;
2. Broader orchestration applicability.

## Getting Started

### System Requirements

Task supports 3 major operating systems:

* [Linux][Linux]
* [macOS][macOS]
* [Windows][Windows]

To use or test Task, Task following software must first be installed on your
system:

* [Go][Go] 1.17.2 or higher

In addition, Task following 3rd-party Go package dependencies must also be
installed:

* [cli][cli] v2.3.0 or higher
* [yaml][yaml] v2.2.3 or higher

### Installation

Installation instructions for this project are pending.

### Usage

Task can be used from a shell via Task command `Task`. Its syntax is as follows:

```sh
$ Task [-h|--help] | TASK [-c|--config YAML_PATH] [-d|--data JSON_STRING] [-|--environment SDLC_ENV]
```

For example:

```sh
$ task build -c cicd.yml
$ task all -c etl.yml -d '{"start_date":"2009-01-01","end_date":"2021-07-31"}'
$ task train -c ml.yml
```

## Files

Files in this project are organized as follows:

```bash
.
├── .editorconfig
├── .gitattributes
├── .gitignore
├── .make/
│   ├── features/
│   │   ├── formatting.mk
│   │   └── helping.mk
│   ├── platforms/
│   │   ├── Go.mk
│   │   └── Python.mk
│   └── test/
│       └── task.mk
├── .pre-commit-config.yaml
├── .vscode/
│   └── settings.json
├── CODE_OF_CONDUCT.md
├── CONTRIBUTING.md
├── README.md
├── REFERENCES.md
├── cmd/
│   └── task/
│       └── main.go
├── examples/
│   ├── sample.etl.yml
│   └── sample.hello.yml
├── go.mod
├── go.sum
└── pkg/
    └── task/
        ├── config.go
        ├── data.go
        ├── file.go
        ├── parameters.go
        ├── task.go
        ├── task_test.go
        └── utils.go
```

## Builds and Testing

Build and testing instructions for this project are pending.

## Documentation

Documentation for Task is this README itself.

## Known Issues

Currently, there are no known issues.

## References

API documentation, tutorials, and other online references and made portions of
Task possible. See [REFERENCES.md](REFERENCES.md) for a list of some.

[CICD]: https://en.wikipedia.org/wiki/CI/CD
[CLI]: https://en.wikipedia.org/wiki/Command-line_interface
[cli]: https://github.com/urfave/cli/
[exit status]: https://en.wikipedia.org/wiki/Exit_status
[Git for Windows]: https://gitforwindows.org
[Go]: https://golang.org/dl/
[Linux]: https://www.linuxfoundation.org
[Make]: https://www.gnu.org/software/make/
[pipeline]: https://en.wikipedia.org/wiki/Task_(Unix)
[software development lifecycle]: https://en.wikipedia.org/wiki/Software_development_process
[Windows]: https://www.microsoft.com/en-us/windows
[WSL]: https://docs.microsoft.com/en-us/windows/wsl/about
[macOS]: https://www.apple.com/macos/
[YAML]: https://yaml.org
[yaml]: https://gopkg.in/yaml.v2
