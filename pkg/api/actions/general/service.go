package general

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Service represents all General application services.
type Service interface {
	ExecuteRBS(string) (rpi.Action, error)
}

// General represents a General application service.
type General struct {
	gensys GENSYS
	a      Actions
}

// GENSYS represents a General repository service.
type GENSYS interface {
	ExecuteRBS(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	ExecuteBashCommand(interface{}) (rpi.Exec, error)
}

// New creates a GENSYS application service instance.
func New(dessys GENSYS, a Actions) *General {
	return &General{gensys: dessys, a: a}
}
