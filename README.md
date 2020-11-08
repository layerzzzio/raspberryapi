# rpi

Light version of the RaspiBuddy API.

## Branching
- one master branch
- one develop branch
- multiple feature/fix/etc. branches

## Regular workflow
If a feature has to be added, a branch starting with ft/[] is created.
The development is done in this branch.

Once done, the feature branch is merge to develop with a PR.

The develop release is then tested in a real environment to test the quality of the code and find bugs.

If after some time, the develop release works well, it is merge with master.

## A minor bug to fix
Create a fix/[] branch, fix the bug.
Then push to develop.
Intensify the tests and push directly to master.

## Hot fix
Create a hotfix/[] branch, fix the bug and push directly to master.
