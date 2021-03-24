package install

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Service represents all Install application services.
type Service interface {
	ExecuteAG(string, string) (rpi.Action, error)
	ExecuteNV(string) (rpi.Action, error)
}

// Install represents a Install application service.
type Install struct {
	inssys INSSYS
	a      Actions
}

// INSSYS represents a Install repository service.
type INSSYS interface {
	ExecuteAG(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteNV(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	ExecuteBashCommand(interface{}) (rpi.Exec, error)
}

// New creates a INSSYS application service instance.
func New(inssys INSSYS, a Actions) *Install {
	return &Install{inssys: inssys, a: a}
}
