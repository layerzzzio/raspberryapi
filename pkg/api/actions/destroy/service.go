package destroy

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all Destroy application services.
type Service interface {
	ExecuteDF(string) (rpi.Action, error)
	ExecuteDU(string, string) (rpi.Action, error)
	ExecuteKP(int) (rpi.Action, error)
}

// Destroy represents a Destroy application service.
type Destroy struct {
	dessys DESSYS
	a      Actions
}

// DESSYS represents a Destroy repository service.
type DESSYS interface {
	ExecuteDF(map[int]rpi.Exec) (rpi.Action, error)
	ExecuteDU(map[int]rpi.Exec) (rpi.Action, error)
	ExecuteKP(map[int]rpi.Exec) (rpi.Action, error)
}

// Actions represents the system metrics interface
type Actions interface {
	DeleteFile(string) rpi.Exec
	DisconnectUser(string, string) rpi.Exec
	KillProcess(int) rpi.Exec
}

// New creates a DESSYS application service instance.
func New(dessys DESSYS, a Actions) *Destroy {
	return &Destroy{dessys: dessys, a: a}
}
