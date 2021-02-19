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
}

// Configure represents a Configure application service.
type Configure struct {
	consys CONSYS
	a      Actions
}

// CONSYS represents a Configure repository service.
type CONSYS interface {
	ExecuteCH(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteCP(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteWNB(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	ChangeHostnameInHostnameFile(interface{}) (rpi.Exec, error)
	ChangeHostnameInHostsFile(interface{}) (rpi.Exec, error)
	ChangePassword(interface{}) (rpi.Exec, error)
	WaitForNetworkAtBoot(interface{}) (rpi.Exec, error)
}

// New creates a CONSYS application service instance.
func New(consys CONSYS, a Actions) *Configure {
	return &Configure{consys: consys, a: a}
}
