package mock

import (
	"github.com/raspibuddy/rpi"
)

// Actions mock
type Actions struct {
	DeleteFileFn                   func(arg interface{}) (rpi.Exec, error)
	KillProcessByNameFn            func(arg interface{}) (rpi.Exec, error)
	KillProcessFn                  func(arg interface{}) (rpi.Exec, error)
	ChangeHostnameInHostnameFileFn func(arg interface{}) (rpi.Exec, error)
	ChangeHostnameInHostsFileFn    func(arg interface{}) (rpi.Exec, error)
	ChangePasswordFn               func(arg interface{}) (rpi.Exec, error)
	WaitForNetworkAtBootFn         func(arg interface{}) (rpi.Exec, error)
	DisableOrEnableOverscanFn      func(arg interface{}) (rpi.Exec, error)
	CommentOverscanFn              func(arg interface{}) (rpi.Exec, error)
	DisableOrEnableBlankingFn      func(arg interface{}) (rpi.Exec, error)
}

// DeleteFile mock
func (a Actions) DeleteFile(arg interface{}) (rpi.Exec, error) {
	return a.DeleteFileFn(arg)
}

// KillProcessByName mock
func (a Actions) KillProcessByName(arg interface{}) (rpi.Exec, error) {
	return a.KillProcessByNameFn(arg)
}

// KillProcess mock
func (a Actions) KillProcess(arg interface{}) (rpi.Exec, error) {
	return a.KillProcessFn(arg)
}

// ChangeHostnameInHostnameFile mock
func (a Actions) ChangeHostnameInHostnameFile(arg interface{}) (rpi.Exec, error) {
	return a.ChangeHostnameInHostnameFileFn(arg)
}

// ChangeHostnameInHostsFile mock
func (a Actions) ChangeHostnameInHostsFile(arg interface{}) (rpi.Exec, error) {
	return a.ChangeHostnameInHostsFileFn(arg)
}

// ChangePassword mock
func (a Actions) ChangePassword(arg interface{}) (rpi.Exec, error) {
	return a.ChangePasswordFn(arg)
}

// WaitForNetworkAtBoot mock
func (a Actions) WaitForNetworkAtBoot(arg interface{}) (rpi.Exec, error) {
	return a.WaitForNetworkAtBootFn(arg)
}

// DisableOrEnableOverscan mock
func (a Actions) DisableOrEnableOverscan(arg interface{}) (rpi.Exec, error) {
	return a.DisableOrEnableOverscanFn(arg)
}

// CommentOverscan mock
func (a Actions) CommentOverscan(arg interface{}) (rpi.Exec, error) {
	return a.CommentOverscanFn(arg)
}

func (a Actions) DisableOrEnableBlanking(arg interface{}) (rpi.Exec, error) {
	return a.DisableOrEnableBlankingFn(arg)
}
