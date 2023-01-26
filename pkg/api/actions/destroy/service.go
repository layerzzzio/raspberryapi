package destroy

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Service represents all Destroy application services.
type Service interface {
	ExecuteDF(string) (rpi.Action, error)
	ExecuteSUS(string, string) (rpi.Action, error)
	ExecuteKP(int) (rpi.Action, error)
}

// Destroy represents a Destroy application service.
type Destroy struct {
	dessys DESSYS
	a      Actions
}

// DESSYS represents a Destroy repository service.
type DESSYS interface {
	ExecuteDF(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteSUS(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteKP(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	DeleteFile(interface{}) (rpi.Exec, error)
	KillProcessByName(interface{}) (rpi.Exec, error)
	KillProcess(interface{}) (rpi.Exec, error)
}

// New creates a DESSYS application service instance.
func New(dessys DESSYS, a Actions) *Destroy {
	return &Destroy{dessys: dessys, a: a}
}
