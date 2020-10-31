package deletefile

import (
	"log"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Execute delete file(s) and returns an array of LargestFile model.
func (df *DeleteFile) Execute(path string) (rpi.Action, error) {
	startTime := uint64(time.Now().Unix())

	steps := map[int]string{
		1: actions.DeleteFile,
	}

	execs := make([]rpi.Exec, len(steps))
	execs[0] = df.a.DeleteFile(path)

	if len(execs) != len(steps) {
		log.Fatal("execs elements count different from steps elements count")
	}

	endTime := uint64(time.Now().Unix())

	return df.delsys.Execute(actions.DeleteFile, steps, execs, startTime, endTime)
}
