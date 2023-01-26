# Raspberry API Contributing Guide

Hi! We're really excited that you are interested in contributing to Raspberry API. Before submitting your contribution, please make sure to take a moment and read through the following guidelines:

- [Code of Conduct](https://github.com/layerzzzio/raspberryapi/blob/main/CODE_OF_CONDUCT.md)

## Pull Request Guidelines

- Checkout a topic branch from the relevant branch, e.g. `main`, and merge back against that branch.

- If adding a new feature:

  - Provide a convincing reason to add this feature. Ideally, you should open a suggestion issue first and have it approved before working on it.

- If fixing bug:

  - Provide a detailed description of the bug in the PR. Live demo preferred.

- It's OK to have multiple small commits as you work on the PR - GitHub can automatically squash them before merging.

## Development Setup

You will need Golang installed on your local machine as well as a Raspberry Pi to test your code on a real device.

## Branching

There are 3 types of branches:

- one main branch (long-lived)
- one develop branch (long-lived)
- multiple feature/fix/etc. branches

### Add a feature
To implement new functionality, a dedicated branch prefixed with "ft/" is established. Development efforts are focused on this branch, upon completion, it is merged with the "develop" branch via a pull request. The "develop" release is then thoroughly tested in a live environment to evaluate code quality and identify any issues. Once deemed stable, the "develop" release is subsequently merged with the "main" branch.

### Fix bug
To resolve bugs, a "fix/" branch is established and the necessary fixes are implemented. The updated code is then pushed to the "develop" branch and tested on a real device. Upon successful testing, the changes are merged into the "main" branch.

### Hot fix
Create a hotfix/* branch, fix the bug and push directly to main.
