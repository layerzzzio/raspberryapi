package configure

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Service represents all Configure application services.
type Service interface {
	ExecuteCH(string) (rpi.Action, error)
	ExecuteCP(string, string) (rpi.Action, error)
	ExecuteWNB(string) (rpi.Action, error)
	ExecuteOV(string) (rpi.Action, error)
	ExecuteBL(string) (rpi.Action, error)
}

// Configure represents a Configure application service.
type Configure struct {
	consys CONSYS
	a      Actions
	i      Infos
}

// CONSYS represents a Configure repository service.
type CONSYS interface {
	ExecuteCH(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteCP(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteWNB(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteOV(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteBL(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	ChangeHostnameInHostnameFile(interface{}) (rpi.Exec, error)
	ChangeHostnameInHostsFile(interface{}) (rpi.Exec, error)
	ChangePassword(interface{}) (rpi.Exec, error)
	WaitForNetworkAtBoot(interface{}) (rpi.Exec, error)
	DisableOrEnableOverscan(interface{}) (rpi.Exec, error)
	CommentOverscan(interface{}) (rpi.Exec, error)
	DisableOrEnableBlanking(interface{}) (rpi.Exec, error)
}

// Infos represents the infos interface
type Infos interface {
	GetConfigFiles() map[string]rpi.ConfigFileDetails
}

// New creates a CONSYS application service instance.
func New(consys CONSYS, a Actions, i Infos) *Configure {
	return &Configure{consys: consys, a: a, i: i}
}
