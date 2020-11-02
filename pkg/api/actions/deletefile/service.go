package deletefile

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all DeleteFile application services.
type Service interface {
	Execute(string) (rpi.Action, error)
}

// DeleteFile represents a DeleteFile application service.
type DeleteFile struct {
	delsys DELSYS
	a      Actions
}

// DELSYS represents a DeleteFile repository service.
type DELSYS interface {
	Execute(map[int]rpi.Exec) (rpi.Action, error)
}

// Actions represents the system metrics interface
type Actions interface {
	DeleteFile(string) rpi.Exec
}

// New creates a DELSYS application service instance.
func New(delsys DELSYS, a Actions) *DeleteFile {
	return &DeleteFile{delsys: delsys, a: a}
}
