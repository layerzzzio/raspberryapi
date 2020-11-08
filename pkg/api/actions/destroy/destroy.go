package destroy

import (
	"github.com/raspibuddy/rpi"
)

// ExecuteDF delete file(s) and returns an action.
func (des *Destroy) ExecuteDF(path string) (rpi.Action, error) {
	// populate execs in the right execution order
	// always start with 1, not with 0
	execs := map[int]rpi.Exec{
		1: des.a.DeleteFile(path),
	}

	return des.dessys.ExecuteDF(execs)
}

// ExecuteDU disconnect user and returns an action.
func (des *Destroy) ExecuteDU(terminal string, username string) (rpi.Action, error) {
	// populate execs in the right execution order
	// always start with 1, not with 0
	execs := map[int]rpi.Exec{
		1: des.a.DisconnectUser(terminal, username),
	}

	return des.dessys.ExecuteDU(execs)
}

// ExecuteKP kill a process and returns an action.
func (des *Destroy) ExecuteKP(pid int) (rpi.Action, error) {
	// populate execs in the right execution order
	// always start with 1, not with 0
	execs := map[int]rpi.Exec{
		1: des.a.KillProcess(pid),
	}

	return des.dessys.ExecuteKP(execs)
}