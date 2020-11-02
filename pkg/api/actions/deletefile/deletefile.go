package deletefile

import (
	"github.com/raspibuddy/rpi"
)

// Execute delete file(s) and returns an array of LargestFile model.
func (df *DeleteFile) Execute(path string) (rpi.Action, error) {
	// populate execs in the right execution order
	// always start with 1, not with 0
	execs := map[int]rpi.Exec{
		1: df.a.DeleteFile(path),
	}

	return df.delsys.Execute(execs)
}
