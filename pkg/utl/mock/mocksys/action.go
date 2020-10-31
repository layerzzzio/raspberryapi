package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Action mock
type Action struct {
	ExecuteFn func(string, map[int]string, []rpi.Exec, uint64, uint64) (rpi.Action, error)
}

// Execute mock
func (a *Action) Execute(name string, steps map[int]string, execs []rpi.Exec, startTime uint64, endTime uint64) (rpi.Action, error) {
	return a.ExecuteFn(name, steps, execs, startTime, endTime)
}
