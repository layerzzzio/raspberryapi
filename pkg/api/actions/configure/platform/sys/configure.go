package sys

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Configure represents an empty Configure entity on the current system.
type Configure struct{}

// ExecuteCH returns an action response after changing hostname
func (con Configure) ExecuteCH(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.ChangeHostname,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteCP returns an action response after changing password
func (con Configure) ExecuteCP(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.ChangePassword,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteWNB returns an action response after enabling or disable wait for network at boot
func (con Configure) ExecuteWNB(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.WaitForNetworkAtBoot,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteOV returns an action response after enabling or disable overscan
func (con Configure) ExecuteOV(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.Overscan,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteBL returns an action response after enabling or disable blanking
func (con Configure) ExecuteBL(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.Blanking,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteAUS returns an action response after adding a user
func (con Configure) ExecuteAUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.AddUser,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteDUS returns an action response after deleting a user
func (con Configure) ExecuteDUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.DeleteUser,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteCA returns an action response after enabling or disable camera
func (con Configure) ExecuteCA(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.CameraInterface,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}
