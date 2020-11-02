package destroy

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all Destroy application services.
type Service interface {
	ExecuteDF(string) (rpi.Action, error)
}

// Destroy represents a Destroy application service.
type Destroy struct {
	dessys DESSYS
	a      Actions
}

// DESSYS represents a Destroy repository service.
type DESSYS interface {
	ExecuteDF(map[int]rpi.Exec) (rpi.Action, error)
}

// Actions represents the system metrics interface
type Actions interface {
	DeleteFile(string) rpi.Exec
}

// New creates a DESSYS application service instance.
func New(dessys DESSYS, a Actions) *Destroy {
	return &Destroy{dessys: dessys, a: a}
}