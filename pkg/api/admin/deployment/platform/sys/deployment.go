package sys

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Deployment represents a Deployment entity on the current system.
type Deployment struct{}

// ExecuteDP deploys an API version
func (d Deployment) ExecuteDP(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.DeployVersion,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}
