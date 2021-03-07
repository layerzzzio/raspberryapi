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
	ExecuteCHFn  func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteCPFn  func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteWNBFn func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteOVFn  func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteBLFn  func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteAUSFn func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteDUSFn func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteCAFn  func(map[int](map[int]actions.Func)) (rpi.Action, error)
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

// ExecuteCH mock
func (a *Action) ExecuteCH(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteCHFn(plan)
}

// ExecuteCP mock
func (a *Action) ExecuteCP(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteCPFn(plan)
}

// ExecuteWNB mock
func (a *Action) ExecuteWNB(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteWNBFn(plan)
}

// ExecuteOV mock
func (a *Action) ExecuteOV(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteOVFn(plan)
}

// ExecuteBL mock
func (a *Action) ExecuteBL(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteBLFn(plan)
}

// ExecuteAUS mock
func (a *Action) ExecuteAUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteAUSFn(plan)
}

// ExecuteDUS mock
func (a *Action) ExecuteDUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteDUSFn(plan)
}

// ExecuteCA mock
func (a *Action) ExecuteCA(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteCAFn(plan)
}
