package deployment

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Service represents all Deployment application services.
type Service interface {
	ExecuteDPTOOL(string, string, string) (rpi.Action, error)
}

// Deployment represents an Deployment application service.
type Deployment struct {
	dsys DSYS
	a    Actions
}

// DSYS represents a Deployment repository service.
type DSYS interface {
	ExecuteDPTOOL(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	ExecuteBashCommand(interface{}) (rpi.Exec, error)
}

// New creates a Deployment application service instance.
func New(dsys DSYS, a Actions) *Deployment {
	return &Deployment{dsys: dsys, a: a}
}
