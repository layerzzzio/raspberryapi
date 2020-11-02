package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Action mock
type Action struct {
	ExecuteDFFn func(map[int]rpi.Exec) (rpi.Action, error)
}

// ExecuteDF mock
func (a *Action) ExecuteDF(execs map[int]rpi.Exec) (rpi.Action, error) {
	return a.ExecuteDFFn(execs)
}
