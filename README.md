![raspberryapi_github_banner_final_small](https://user-images.githubusercontent.com/98493964/215260649-002553d8-8c77-40ef-a31a-c3c3a00392d7.jpg)

[![codecov](https://codecov.io/gh/layerzzzio/raspberryapi/branch/main/graph/badge.svg?token=IKY9WGCDIY)](https://codecov.io/gh/layerzzzio/raspberryapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspibuddy/rpi)](https://goreportcard.com/report/github.com/raspibuddy/rpi)

# Raspberry API

Raspberry API is an open-source application programming interface, developed using the Go programming language, which enables retrieval of vital metrics and other data from a Raspberry Pi device. Additionally, it allows for the execution of various actions such as package installation and device configuration.

<a style="font-weight: bold; " href="https://raspberryapi.com/docs/getting-started.html">Getting started + API documentation 👩‍💻</a>

# Capabilities

## 🩺 Fetch vital metrics

It provides access to various metrics of your Raspberry Pi including RAM, CPU usage, disk usage, temperature, active processes, network activity, and user sessions among others.

## 🚑 Fix issues

It enables you to troubleshoot and resolve issues on your Raspberry Pi, such as terminating resource-intensive processes or deleting large files.

## ⚙️ Configure your device

It allows for the execution of common administrative tasks, such as adding or deleting users, changing passwords, and updating the package repository. It also enables or disables common Raspberry Pi interfaces, such as I2C and 1Wire.

## 📱 Manage packages

It enables you to install packages on your Raspberry Pi device and also provides the ability to start or stop them.

# Helpful tutorials for developers

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
