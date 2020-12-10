package mock

import (
	"github.com/raspibuddy/rpi"
)

// Actions mock
type Actions struct {
	DeleteFileFn        func(arg interface{}) (rpi.Exec, error)
	KillProcessByNameFn func(arg interface{}) (rpi.Exec, error)
	KillProcessFn       func(arg interface{}) (rpi.Exec, error)
}

// DeleteFile mock
func (a Actions) DeleteFile(arg interface{}) (rpi.Exec, error) {
	return a.DeleteFileFn(arg)
}

// KillProcessByName mock
func (a Actions) KillProcessByName(arg interface{}) (rpi.Exec, error) {
	return a.KillProcessByNameFn(arg)
}

// KillProcess mock
func (a Actions) KillProcess(arg interface{}) (rpi.Exec, error) {
	return a.KillProcessFn(arg)
}
