# yamllint-checkstyle convert [yamllint](https://github.com/adrienverge/yamllint) to [checkstyle](https://checkstyle.sourceforge.io/releasenotes.html) report
[![Release version](https://img.shields.io/github/v/release/thomaspoignant/yamllint-checkstyle "Release version")](https://github.com/thomaspoignant/yamllint-checkstyle/releases)
[![Build Status](https://travis-ci.com/thomaspoignant/yamllint-checkstyle.svg?branch=master)](https://travis-ci.com/thomaspoignant/yamllint-checkstyle)
[![Coverage Status](https://coveralls.io/repos/github/thomaspoignant/yamllint-checkstyle/badge.svg?branch=master)](https://coveralls.io/github/thomaspoignant/yamllint-checkstyle?branch=master)
[![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=thomaspoignant_yamllint-checkstyle&metric=alert_status)](https://sonarcloud.io/dashboard?id=thomaspoignant_yamllint-checkstyle)
![Go version](https://img.shields.io/github/go-mod/go-version/thomaspoignant/yamllint-checkstyle?logo=go%20version "Go version")
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/thomaspoignant/yamllint-checkstyle?color=blue&logo=docker&sort=semver)](https://hub.docker.com/r/thomaspoignant/yamllint-checkstyle)

## Installation

### via brew  _(mac & linux)_.
```shell script
brew tap thomaspoignant/homebrew-tap
brew install yamllint-checkstyle
```

### via scoop _(windows)_.
```shell script
scoop bucket add org https://github.com/thomaspoignant/scoop.git
scoop install yamllint-checkstyle
```


## Usage
run `yamllint-checkstyle` and pass `yamllint` (with `-f` parsable option) output to it
```shell script
yamllint -f parsable test.yaml | yamllint-checkstyle > yamllint-junit.xml
```

## Output
- if there are any lint errors, full Checkstyle XML will be created
- if there are no errors, empty Checkstyle XML will be created.

# How can I contribute?
See the [contributor's guide](CONTRIBUTING.md) for some helpful tips.