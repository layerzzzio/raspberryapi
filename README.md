![raspberryapi_github_banner_readme](https://user-images.githubusercontent.com/98493964/214853808-9ef7599f-4097-4df8-b5b1-8fcfa40e5ebf.png)

[![codecov](https://codecov.io/gh/layerzzzio/raspberryapi/branch/main/graph/badge.svg?token=IKY9WGCDIY)](https://codecov.io/gh/layerzzzio/raspberryapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspibuddy/rpi)](https://goreportcard.com/report/github.com/raspibuddy/rpi)

# Getting started

Raspberry API is an open-source application programming interface, developed using the Go programming language, which enables retrieval of vital metrics and other data from a Raspberry Pi device. Additionally, it allows for the execution of various actions such as package installation and device configuration.

<a href="https://raspberryapi.com/docs/getting-started.html">Getting started 👩‍💻</a>

# Contributing

There are 3 types of branches:
- one main branch (long-lived)
- one develop branch (long-lived)
- multiple feature/fix/etc. branches

## Prerequisites 

To have Go installed on your machine.

## Regular workflow

To implement new functionality, a dedicated branch prefixed with "ft/" is established. Development efforts are focused on this branch, upon completion, it is merged with the "develop" branch via a pull request. The "develop" release is then thoroughly tested in a live environment to evaluate code quality and identify any issues. Once deemed stable, the "develop" release is subsequently merged with the "main" branch.

## A minor bug to fix

To resolve bugs, a "fix/*" branch is established and the necessary fixes are implemented. The updated code is then pushed to the "develop" branch and tested on a real device. Upon successful testing, the changes are merged into the "main" branch.

## Hot fix

Create a hotfix/* branch, fix the bug and push directly to main.

# Info about the linter

The following linter is used in the project: https://golangci-lint.run/usage/install#github-actions. It is made by the same organization that used to maintain GolangCI.com, as outlined in this Medium article https://medium.com/golangci/golangci-com-is-closing-d1fc1bd30e0e. Despite the discontinuation of the website, the linter golangci-lint is still actively maintained by the same team.

It can be used in VS Code that way:
Code > Preferences > Settings > Extensions > Go > Linter Tool

# Info about CI/CD

The CI/CD pipelines are inspired by Bruno Paz:
https://brunopaz.dev/blog/building-a-basic-ci-cd-pipeline-for-a-golang-application-using-github-actions

There 2 of them located in .github/workflows: 
- build (lint the code, run the tests and compile the code)
- release (compile the code and release a version - trigger when a tag is pushed to the remote repository)

FYI, some of the GitHub Actions used:
- codecov https://github.com/marketplace/actions/codecov
- golang-ci lint https://github.com/golangci/golangci-lint-action

# Useful how-tos

## Compiling the program locally

Compiling for Raspberry Pi OS:
```
# from the root of the project
cd cmd/api
env GOOS=linux GOARCH=arm GOARM=5 go build -o raspberryapi
```

Note: parts of the API work on Mac. You can compile it that way:
```
# from the root of the project
cd cmd/api
go build -o raspberryapi
```

## Detect data race

Per file:
```
/usr/local/bin/go test -timeout 30s -race metrics_test.go
/usr/local/bin/go test -timeout 620s -race actions_test.go
```

Per test:
```
/usr/local/bin/go test -timeout 30s -run ^TestCall$ github.com/raspibuddy/rpi/pkg/utl/actions -race
/usr/local/bin/go test -timeout 30s -run ^TestExecutePlanWithoutDependency$ github.com/raspibuddy/rpi/pkg/utl/actions -race
/usr/local/bin/go test -timeout 30s -run ^TestExecutePlanWithDependency$ github.com/raspibuddy/rpi/pkg/utl/actions -race
```

## Detect leak

Use the following ressource: https://medium.com/a-journey-with-go/go-goroutine-leak-detector-61a949beb88
