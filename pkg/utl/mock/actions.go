package mock

import (
	"github.com/raspibuddy/rpi"
)

// Actions mock
type Actions struct {
	DeleteFileFn func(path string) rpi.Exec
}

// DeleteFile mock
func (a Actions) DeleteFile(path string) rpi.Exec {
	return a.DeleteFileFn(path)
}
