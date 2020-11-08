package actions

import (
	"fmt"
	"os"
	"time"

	"github.com/raspibuddy/rpi"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

var (
	// DeleteFile is the name of the delete file exec
	DeleteFile = "delete_file"

	// KillProcess is the name of the kill process exec
	KillProcess = "kill_process"

	// DisconnectUser is the name of the disconnect user exec
	DisconnectUser = "disconnect_user"
)

// Service represents several system scripts.
type Service struct {
	a Actions
}

// Actions represents multiple system related action scripts.
type Actions interface{}

// New creates a service instance.
func New() *Service {
	return &Service{}
}

// KillProcess kill a given process
func (s Service) KillProcess(pid int) rpi.Exec {
	startTime := uint64(time.Now().Unix())

	exitStatus := 0
	var stdErr string

	ps, err := os.FindProcess(pid)

	if err != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(err)
	} else {
		e := ps.Kill()
		if e != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(e)
		}
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       KillProcess,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}
}

// DisconnectUser disconnect a user from an active tty from the current host
func (s Service) DisconnectUser(path string, username string) rpi.Exec {
	startTime := uint64(time.Now().Unix())

	exitStatus := 0
	var stdErr string
	e := os.Remove(path)
	if e != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(e)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       DisconnectUser,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}
}

// DeleteFile deletes a file or (empty) directory
func (s Service) DeleteFile(path string) rpi.Exec {
	// execution start time
	startTime := uint64(time.Now().Unix())

	exitStatus := 0
	var stdErr string
	e := os.Remove(path)
	if e != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(e)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       DeleteFile,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}
}
