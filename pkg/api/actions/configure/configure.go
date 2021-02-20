package configure

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

var (
	// BootConfig contains the boot configs
	BootConfig = "/boot/config.txt"
)

// ExecuteCH changes hostname and returns an action.
func (con *Configure) ExecuteCH(hostname string) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ChangeHostnameInHostsFile,
				Reference: con.a.ChangeHostnameInHostsFile,
				Argument: []interface{}{
					actions.DataToFile{
						TargetFile: "/etc/hosts",
						Data:       hostname,
					},
				},
			},
			2: {
				Name:      actions.ChangeHostnameInHostnameFile,
				Reference: con.a.ChangeHostnameInHostnameFile,
				Argument: []interface{}{
					actions.DataToFile{
						TargetFile: "/etc/hostname",
						Data:       hostname,
					},
				},
			},
		},
	}

	return con.consys.ExecuteCH(plan)
}

// ExecuteCP changes password and returns an action.
func (con *Configure) ExecuteCP(password string, username string) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ChangePassword,
				Reference: con.a.ChangePassword,
				Argument: []interface{}{
					actions.CP{
						Password: password,
						Username: username,
					},
				},
			},
		},
	}

	return con.consys.ExecuteCP(plan)
}

// ExecuteWNB enable or disable wait for network at boot and returns an action
func (con *Configure) ExecuteWNB(action string) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.WaitForNetworkAtBoot,
				Reference: con.a.WaitForNetworkAtBoot,
				Argument: []interface{}{
					actions.WNB{
						Directory: "/etc/systemd/system/dhcpcd.service.d",
						Action:    action,
					},
				},
			},
		},
	}

	return con.consys.ExecuteWNB(plan)
}

// ExecuteOV enable or disable overscan and returns an action
func (con *Configure) ExecuteOV(action string) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.Overscan,
				Reference: con.a.Overscan,
				Argument: []interface{}{
					actions.WNB{
						Directory: BootConfig,
						Action:    action,
					},
				},
			},
		},
	}

	return con.consys.ExecuteOV(plan)
}
