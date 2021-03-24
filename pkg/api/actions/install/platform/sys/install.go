package sys

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Install represents an empty Install entity on the current system.
type Install struct{}

// ExecuteAG returns an action response after installing package with apt-get
func (ins Install) ExecuteAG(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.InstallAptGet,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteNV returns an action response after installing nordvpn
func (ins Install) ExecuteNV(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.InstallNordVPN,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}
