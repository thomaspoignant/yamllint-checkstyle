# yamllint-checkstyle convert [yamllint](https://github.com/adrienverge/yamllint) to [checkstyle](https://checkstyle.sourceforge.io/releasenotes.html) report
[![Build Status](https://travis-ci.com/thomaspoignant/yamllint-checkstyle.svg?branch=master)](https://travis-ci.com/thomaspoignant/yamllint-checkstyle)

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
run `yamllint-checkstyle` and pass `yamllint` (with -f parsable option) output to it
```shell script
yamllint -f parsable test.yaml | yamllint-checkstyle > yamllint-junit.xml
```

## Output
- if there are any lint errors, full Checkstyle XML will be created
- if there are no errors, empty Checkstyle XML will be created.
