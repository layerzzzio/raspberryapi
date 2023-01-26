[![codecov](https://codecov.io/gh/layerzzzio/raspberryapi/branch/main/graph/badge.svg?token=IKY9WGCDIY)](https://codecov.io/gh/layerzzzio/raspberryapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspibuddy/rpi)](https://goreportcard.com/report/github.com/raspibuddy/rpi)

![raspberryapi_github_banner](https://user-images.githubusercontent.com/98493964/214852032-edbbc2b7-2fe7-4a8b-8560-f50c66116402.png)

# Getting started

Raspberry API is an open-source API written in Golang for fetching critical metrics and other information from a Raspberry Pi. It also allows executing actions on it, such as installing packages or configuring the device.

<a href="https://raspberryapi.com/docs/getting-started.html">Getting started 👩‍💻</a>

# Branching: how-to?

There are 3 types of branches:
- one main branch (long-lived)
- one develop branch (long-lived)
- multiple feature/fix/etc. branches

## Regular workflow

If a feature has to be added, a branch starting with ft/[] is created.
The development is done in this branch.
Once done, the feature branch is merged to develop with a PR.
The develop release is then tested in a real environment to test the quality of the code and find bugs.
If after some time, the develop release works well, it is merge with main.

## A minor bug to fix

Create a fix/* branch and fix the bug.
Push to develop and test with a real device.
If it works, merge to main.

## Hot fix

Create a hotfix/* branch, fix the bug and push directly to main.

# Info about the linter

The following linter is being used: https://golangci-lint.run/usage/install#github-actions. It is the same "company" that used to maintain GolangCI.com. They closed GolangCI.com as explained here https://medium.com/golangci/golangci-com-is-closing-d1fc1bd30e0e, but kept maintaining the linter: golangci-lint

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
