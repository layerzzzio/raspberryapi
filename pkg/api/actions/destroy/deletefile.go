package destroy

import (
	"github.com/raspibuddy/rpi"
)

// ExecuteDF delete file(s) and returns an action.
func (df *Destroy) ExecuteDF(path string) (rpi.Action, error) {
	// populate execs in the right execution order
	// always start with 1, not with 0
	execs := map[int]rpi.Exec{
		1: df.a.DeleteFile(path),
	}

	return df.dessys.ExecuteDF(execs)
}
