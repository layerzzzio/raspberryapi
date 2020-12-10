# rpi

Light version of the RaspiBuddy API.

## Branching
- one master branch
- one develop branch
- multiple feature/fix/etc. branches

## Regular workflow
If a feature has to be added, a branch starting with ft/[] is created.
The development is done in this branch.

Once done, the feature branch is merged to develop with a PR.
The develop release is then tested in a real environment to test the quality of the code and find bugs.
If after some time, the develop release works well, it is merge with master.

## A minor bug to fix
Create a fix/* branch, fix the bug.
Then push to develop.
Intensify the tests and push directly to master.

## Hot fix
Create a hotfix/* branch, fix the bug and push directly to master.

## Linter
The following linter is being used: https://golangci-lint.run/usage/install#github-actions
It is the same "company" that used to maintain GolangCI.com.

They closed GolangCI.com as explained below:
https://medium.com/golangci/golangci-com-is-closing-d1fc1bd30e0e

But kept maintaining the linter: golangci-lint

It is used also in VS Code: 
Code > Preferences > Settings > Extensions > Go > Linter Tool

# ci/cd
The CI/CD pipelines are inspired by Bruno Paz: 
https://brunopaz.dev/blog/building-a-basic-ci-cd-pipeline-for-a-golang-application-using-github-actions

Some GitHub Actions used:
- codecov https://github.com/marketplace/actions/codecov
- golang-ci lint https://github.com/golangci/golangci-lint-action

# how to detect data race
## per file
/usr/local/bin/go test -timeout 30s -race metrics_test.go
/usr/local/bin/go test -timeout 620s -race actions_test.go

## per test
/usr/local/bin/go test -timeout 30s -run ^TestCall$ github.com/raspibuddy/rpi/pkg/utl/actions -race
/usr/local/bin/go test -timeout 30s -run ^TestExecutePlanWithoutDependency$ github.com/raspibuddy/rpi/pkg/utl/actions -race
/usr/local/bin/go test -timeout 30s -run ^TestExecutePlanWithDependency$ github.com/raspibuddy/rpi/pkg/utl/actions -race