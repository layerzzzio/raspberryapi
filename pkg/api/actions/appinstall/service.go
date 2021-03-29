package appinstall

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Service represents all AppInstall application services.
type Service interface {
	ExecuteAG(string, string) (rpi.Action, error)
	ExecuteWOV(string, string, string) (rpi.Action, error)
}

// AppInstall represents a AppInstall application service.
type AppInstall struct {
	inssys INSSYS
	a      Actions
}

// INSSYS represents a AppInstall repository service.
type INSSYS interface {
	ExecuteAG(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteWOV(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	ExecuteBashCommand(interface{}) (rpi.Exec, error)
}

// New creates a INSSYS application service instance.
func New(inssys INSSYS, a Actions) *AppInstall {
	return &AppInstall{inssys: inssys, a: a}
}
