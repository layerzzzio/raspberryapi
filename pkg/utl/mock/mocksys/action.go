package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Action mock
type Action struct {
	ExecuteDFFn func(map[int]rpi.Exec) (rpi.Action, error)
	ExecuteDUFn func(map[int]rpi.Exec) (rpi.Action, error)
	ExecuteKPFn func(map[int]rpi.Exec) (rpi.Action, error)
}

// ExecuteDF mock
func (a *Action) ExecuteDF(execs map[int]rpi.Exec) (rpi.Action, error) {
	return a.ExecuteDFFn(execs)
}

// ExecuteDU mock
func (a *Action) ExecuteDU(execs map[int]rpi.Exec) (rpi.Action, error) {
	return a.ExecuteDUFn(execs)
}

// ExecuteKP mock
func (a *Action) ExecuteKP(execs map[int]rpi.Exec) (rpi.Action, error) {
	return a.ExecuteKPFn(execs)
}
