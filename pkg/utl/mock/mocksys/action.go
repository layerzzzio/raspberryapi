package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Action mock
type Action struct {
	ExecuteDFFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteSUSFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteKPFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteCHFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteCPFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteWNBFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteOVFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteBLFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteAUSFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteDUSFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteCAFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteSSHFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteVNCFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteSPIFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteI2CFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteONWFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteRGFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteUPDFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteUPGFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteUPDGFn   func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteWCFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteAGFn     func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteWOVFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteWOVAFn   func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteRBSFn    func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteDPTOOLFn func(map[int](map[int]actions.Func)) (rpi.Action, error)
	ExecuteSASOFn   func(map[int](map[int]actions.Func)) (rpi.Action, error)
}

// ExecuteDF mock
func (a *Action) ExecuteDF(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteDFFn(plan)
}

// ExecuteSUS mock
func (a *Action) ExecuteSUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteSUSFn(plan)
}

// ExecuteKP mock
func (a *Action) ExecuteKP(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteKPFn(plan)
}

// ExecuteCH mock
func (a *Action) ExecuteCH(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteCHFn(plan)
}

// ExecuteCP mock
func (a *Action) ExecuteCP(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteCPFn(plan)
}

// ExecuteWNB mock
func (a *Action) ExecuteWNB(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteWNBFn(plan)
}

// ExecuteOV mock
func (a *Action) ExecuteOV(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteOVFn(plan)
}

// ExecuteBL mock
func (a *Action) ExecuteBL(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteBLFn(plan)
}

// ExecuteAUS mock
func (a *Action) ExecuteAUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteAUSFn(plan)
}

// ExecuteDUS mock
func (a *Action) ExecuteDUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteDUSFn(plan)
}

// ExecuteCA mock
func (a *Action) ExecuteCA(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteCAFn(plan)
}

// ExecuteSSH mock
func (a *Action) ExecuteSSH(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteSSHFn(plan)
}

// ExecuteVNC mock
func (a *Action) ExecuteVNC(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteVNCFn(plan)
}

// ExecuteSPI mock
func (a *Action) ExecuteSPI(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteSPIFn(plan)
}

// ExecuteI2C mock
func (a *Action) ExecuteI2C(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteI2CFn(plan)
}

// ExecuteONW mock
func (a *Action) ExecuteONW(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteONWFn(plan)
}

// ExecuteRG mock
func (a *Action) ExecuteRG(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteRGFn(plan)
}

// ExecuteUPD mock
func (a *Action) ExecuteUPD(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteUPDFn(plan)
}

// ExecuteUPG mock
func (a *Action) ExecuteUPG(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteUPGFn(plan)
}

// ExecuteUPDG mock
func (a *Action) ExecuteUPDG(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteUPDGFn(plan)
}

// ExecuteWC mock
func (a *Action) ExecuteWC(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteWCFn(plan)
}

// ExecuteAG mock
func (a *Action) ExecuteAG(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteAGFn(plan)
}

// ExecuteWOV mock
func (a *Action) ExecuteWOV(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteWOVFn(plan)
}

// ExecuteWOVA mock
func (a *Action) ExecuteWOVA(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteWOVAFn(plan)
}

// ExecuteRBS mock
func (a *Action) ExecuteRBS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteRBSFn(plan)
}

// ExecuteDPTOOL mock
func (a *Action) ExecuteDPTOOL(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteDPTOOLFn(plan)
}

// ExecuteSASO mock
func (a *Action) ExecuteSASO(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	return a.ExecuteSASOFn(plan)
}
