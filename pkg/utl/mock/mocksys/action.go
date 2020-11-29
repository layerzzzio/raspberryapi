package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Action mock
type Action struct {
	ExecuteDFFn  func(map[int]rpi.Exec) (rpi.Action, error)
	ExecuteSUSFn func(map[int]rpi.Exec) (rpi.Action, error)
	ExecuteKPFn  func(map[int]rpi.Exec) (rpi.Action, error)
}

// ExecuteDF mock
func (a *Action) ExecuteDF(execs map[int]rpi.Exec) (rpi.Action, error) {
	return a.ExecuteDFFn(execs)
}

// ExecuteSUS mock
func (a *Action) ExecuteSUS(execs map[int]rpi.Exec) (rpi.Action, error) {
	return a.ExecuteSUSFn(execs)
}

// ExecuteKP mock
func (a *Action) ExecuteKP(execs map[int]rpi.Exec) (rpi.Action, error) {
	return a.ExecuteKPFn(execs)
}
