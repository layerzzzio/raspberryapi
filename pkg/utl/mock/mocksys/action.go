package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Action mock
type Action struct {
	ExecuteFn func(map[int]rpi.Exec) (rpi.Action, error)
}

// Execute mock
func (a *Action) Execute(execs map[int]rpi.Exec) (rpi.Action, error) {
	return a.ExecuteFn(execs)
}
