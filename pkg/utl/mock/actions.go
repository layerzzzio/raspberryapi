package mock

import (
	"github.com/raspibuddy/rpi"
)

// Actions mock
type Actions struct {
	DeleteFileFn     func(path string) rpi.Exec
	DisconnectUserFn func(terminal string, username string) rpi.Exec
	KillProcessFn    func(pid int) rpi.Exec
}

// DeleteFile mock
func (a Actions) DeleteFile(path string) rpi.Exec {
	return a.DeleteFileFn(path)
}

// DisconnectUser mock
func (a Actions) DisconnectUser(terminal string, username string) rpi.Exec {
	return a.DisconnectUserFn(terminal, username)
}

// KillProcess mock
func (a Actions) KillProcess(pid int) rpi.Exec {
	return a.KillProcessFn(pid)
}
