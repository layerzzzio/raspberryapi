package mock

import (
	"github.com/raspibuddy/rpi"
)

// Actions mock
type Actions struct {
	DeleteFileFn        func(path string) rpi.Exec
	KillProcessByNameFn func(processname string, processtype string) rpi.Exec
	KillProcessFn       func(pid string) rpi.Exec
}

// DeleteFile mock
func (a Actions) DeleteFile(path string) rpi.Exec {
	return a.DeleteFileFn(path)
}

// KillProcessByName mock
func (a Actions) KillProcessByName(processname string, processtype string) rpi.Exec {
	return a.KillProcessByNameFn(processname, processtype)
}

// KillProcess mock
func (a Actions) KillProcess(pid string) rpi.Exec {
	return a.KillProcessFn(pid)
}
