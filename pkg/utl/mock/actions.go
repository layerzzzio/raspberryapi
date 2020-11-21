package mock

import (
	"github.com/raspibuddy/rpi"
)

// Actions mock
type Actions struct {
	DeleteFileFn        func(path string) rpi.Exec
	KillProcessByNameFn func(terminalname string) rpi.Exec
	KillProcessFn       func(pid string) rpi.Exec
}

// DeleteFile mock
func (a Actions) DeleteFile(path string) rpi.Exec {
	return a.DeleteFileFn(path)
}

// DisconnectUser mock
func (a Actions) KillProcessByName(terminalname string) rpi.Exec {
	return a.KillProcessByNameFn(terminalname)
}

// KillProcess mock
func (a Actions) KillProcess(pid string) rpi.Exec {
	return a.KillProcessFn(pid)
}
