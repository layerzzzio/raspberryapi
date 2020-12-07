package sys

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Destroy represents an empty Destroy entity on the current system.
type Destroy struct{}

// ExecuteDF returns an action response after deleting a file
func (des Destroy) ExecuteDF(execPlan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(execPlan)
	progress, exitStatus := actions.ExecutePlan(execPlan, progressInit)

	return rpi.Action{
		Name:          actions.DeleteFile,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteSUS returns an action response after stopping a user session
func (des Destroy) ExecuteSUS(execPlan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(execPlan)
	progress, exitStatus := actions.ExecutePlan(execPlan, progressInit)

	return rpi.Action{
		Name:          actions.StopUserSession,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteKP returns an action response after killing a process
func (des Destroy) ExecuteKP(execPlan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(execPlan)
	progress, exitStatus := actions.ExecutePlan(execPlan, progressInit)

	return rpi.Action{
		Name:          actions.KillProcess,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}
