package appaction

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Service represents all AppAction application services.
type Service interface {
	ExecuteWOVA(string, string, string, string, string, string) (rpi.Action, error)
}

// AppAction represents a AppAction application service.
type AppAction struct {
	aacsys AACSYS
	a      Actions
	i      Infos
}

// AACSYS represents a AppAction repository service.
type AACSYS interface {
	ExecuteWOVA(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	ExecuteBashCommand(interface{}) (rpi.Exec, error)
	KillProcess(interface{}) (rpi.Exec, error)
}

// Infos represents the infos interface
type Infos interface {
	VPNConfigFiles(string, string, string) []string
	ProcessesPids(string) []string
}

// New creates a INSSYS application service instance.
func New(aacsys AACSYS, a Actions, i Infos) *AppAction {
	return &AppAction{aacsys: aacsys, a: a, i: i}
}
