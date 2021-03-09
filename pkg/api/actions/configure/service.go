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
	ExecuteAUS(string, string) (rpi.Action, error)
	ExecuteDUS(string) (rpi.Action, error)
	ExecuteCA(string) (rpi.Action, error)
	ExecuteSSH(string) (rpi.Action, error)
	ExecuteVNC(string) (rpi.Action, error)
	ExecuteSPI(string) (rpi.Action, error)
	ExecuteI2C(string) (rpi.Action, error)
	ExecuteONW(string) (rpi.Action, error)
	ExecuteRG(string) (rpi.Action, error)
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
	ExecuteAUS(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteDUS(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteCA(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteSSH(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteVNC(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteSPI(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteI2C(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteONW(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteRG(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// Actions represents the actions interface
type Actions interface {
	ChangeHostnameInHostnameFile(interface{}) (rpi.Exec, error)
	ChangeHostnameInHostsFile(interface{}) (rpi.Exec, error)
	ChangePassword(interface{}) (rpi.Exec, error)
	WaitForNetworkAtBoot(interface{}) (rpi.Exec, error)
	CommentOverscan(interface{}) (rpi.Exec, error)
	DisableOrEnableBlanking(interface{}) (rpi.Exec, error)
	AddUser(interface{}) (rpi.Exec, error)
	DeleteUser(interface{}) (rpi.Exec, error)
	DisableOrEnableConfig(interface{}) (rpi.Exec, error)
	CommentOrUncommentInFile(interface{}) (rpi.Exec, error)
	SetVariableInConfigFile(interface{}) (rpi.Exec, error)
	ExecuteBashCommand(interface{}) (rpi.Exec, error)
	DisableOrEnableRemoteGpio(interface{}) (rpi.Exec, error)
}

// Infos represents the infos interface
type Infos interface {
	GetConfigFiles() map[string]rpi.ConfigFileDetails
}

// New creates a CONSYS application service instance.
func New(consys CONSYS, a Actions, i Infos) *Configure {
	return &Configure{consys: consys, a: a, i: i}
}
