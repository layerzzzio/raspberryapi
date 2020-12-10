package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Action mock
type Action struct {
	ExecuteDFFn  func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteSUSFn func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteKPFn  func(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// ExecuteDF mock
func (a *Action) ExecuteDF(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteDFFn(plan)
}

// ExecuteSUS mock
func (a *Action) ExecuteSUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteSUSFn(plan)
}

// ExecuteKP mock
func (a *Action) ExecuteKP(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteKPFn(plan)
}
