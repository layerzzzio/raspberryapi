package sys

import (
	"github.com/raspibuddy/rpi"
)

// DeleteFile represents an empty DeleteFile entity on the current system.
type DeleteFile struct{}

// Execute returns a DeleteFile execution response
func (d DeleteFile) Execute(name string, steps map[int]string, execs []rpi.Exec, startTime uint64, endTime uint64) (rpi.Action, error) {
	// the exit status is the status of the last exec
	return rpi.Action{
		Name:          name,
		Steps:         steps,
		NumberOfSteps: uint16(len(steps)),
		Executions:    execs,
		ExitStatus:    execs[len(execs)-1].ExitStatus,
		StartTime:     startTime,
		EndTime:       endTime,
	}, nil
}
