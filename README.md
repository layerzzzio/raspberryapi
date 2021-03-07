# rpi

RaspiBuddy API (free and paid version).
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

  

# how-to

## detect data race

### per file

/usr/local/bin/go test -timeout 30s -race metrics_test.go

/usr/local/bin/go test -timeout 620s -race actions_test.go

  

### per test

/usr/local/bin/go test -timeout 30s -run ^TestCall$ github.com/raspibuddy/rpi/pkg/utl/actions -race

  

/usr/local/bin/go test -timeout 30s -run ^TestExecutePlanWithoutDependency$ github.com/raspibuddy/rpi/pkg/utl/actions -race

  

/usr/local/bin/go test -timeout 30s -run ^TestExecutePlanWithDependency$ github.com/raspibuddy/rpi/pkg/utl/actions -race

  

## detect leak

  

source: https://medium.com/a-journey-with-go/go-goroutine-leak-detector-61a949beb88

# Daemonize Raspibuddy API
Source: https://baykara.medium.com/how-to-daemonize-a-process-or-service-with-systemd-c34501e646c9

### copy or download the API to the raspberry
    TDB

### make it executable
    chmod 755 raspibuddy
  
 ### copy it to the /usr/bin
	 # it seems to be a good directory on the rpi
	 # as other non-os essential binary are located (vnc, vlc, etc.)
	 # when doing so, the owner changes automatically to root ( <- good!!)
     sudo cp raspibuddy /usr/bin

 ### daemonize the API
	 # create a service file
     sudo vi /etc/systemd/system/raspibuddy.service

Here the tricky question is how to populate the service file.

More information on that here: https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files.

According to the article, a service file is generally composed of the following sections:

##### [Unit] Section Directives (at the beg.)
> The first section found in most unit files is the  `[Unit]`  section. This is generally used for defining metadata for the unit and configuring the relationship of the unit to other units. It includes the **`Description=`** sub-section.

##### The [Service] Section: in between [Unit] and [Install]
> The  `[Service]`  section is used to provide configuration that is only applicable for services. One of the basic things that should be specified within the  `[Service]`  section is the  `Type=`  of the service (**simple**, **forking**, **oneshot**, **idle** etc.). This categorizes services by their process and daemonizing behavior. This is important because it tells  `systemd`  how to correctly manage the service and find out its state.

##### The [Socket] Section: CAUTION ---> NOT IMPLEMENTED HERE !!!
> Socket units are very common in  `systemd`  configurations because many services implement socket-based activation to provide better parallelization and flexibility. **WE ARE NOT GOING TO CREATE A SOCKET SERVICE AS WE ARE CREATING IT DIRECTLY IN THE API.** For more information about that here: https://www.linux.com/training-tutorials/end-road-systemds-socket-units/

##### [Install] Section Directives (at the end)
> On the opposite side of unit file, the last section is often the  `[Install]`  section. This section is optional and is used to define the behavior or a unit if it is enabled or disabled. Enabling a unit marks it to be automatically started at boot. It includes the **`WantedBy=`**. The `WantedBy=` directive is the most common way to specify how a unit should be enabled.

Requirements:
- The API should run as root
> Make sure the service file has a root owner!!
- To start the service once network is up (Wants & After)
> **`Wants=network-online.target`**
**`After=network.target network-online.target`**
Source: https://unix.stackexchange.com/questions/126009/cause-a-script-to-execute-after-networking-has-started
- The service should be started by executing the raspibuddy binary
> **`Type=simple`**: The main process of the service is specified in the start line.
**`ExecStart=/usr/bin/raspibuddy`**: This specifies the full path and the arguments of the command to be executed to start the process.
- The API should restart automatically when the OS is down or the API killed on purpose
> **`Restart=always`**
> **`RestartSec=5`**
- Note:
>The line `WantedBy=multi-user.target` in a service is essentially the same as specifying "this service should start in runlevels 3, 4 and 5" in SysVinit systems: it tells `systemd` that this service should be started as part of normal system start-up, whether or not a local GUI is active.
Source: https://unix.stackexchange.com/questions/126009/cause-a-script-to-execute-after-networking-has-started 

Here is a working version:
````
[Unit]  
Description= Hi there, I am the Raspibuddy API. Hope I can help you :)
Wants=network-online.target
After=network.target network-online.target

[Service]  
Type=simple   
ExecStart=/usr/bin/raspibuddy
Restart=always
RestartSec=5

[Install]  
WantedBy=multi-user.target
````

Then let's create a directory for the conf. file:

    sudo mkdir /etc/raspibuddy
	
	# create the conf file:
    vi raspibuddy.yaml
    
    # include this content:
    server:
	  port: :3333
	  debug: true
	  read_timeout_seconds: 30
	  write_timeout_seconds: 15
    
Finally let's reload the daemon and start the service:
````
sudo chmod 644 /etc/systemd/system/raspibuddy.service
sudo systemctl daemon-reload
sudo systemct enable --now raspibuddy.service #enable and start the service
sudo systemctl status raspibuddy.service
````
