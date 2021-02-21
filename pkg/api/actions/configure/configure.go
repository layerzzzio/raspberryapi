package configure

import (
	"net/http"

	"github.com/labstack/echo"
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
					actions.EnableOrDisableConfig{
						DirOrFilePath: "/etc/systemd/system/dhcpcd.service.d",
						Action:        action,
					},
				},
			},
		},
	}

	return con.consys.ExecuteWNB(plan)
}

// ExecuteOV enable or disable overscan and returns an action
func (con *Configure) ExecuteOV(action string) (rpi.Action, error) {
	var plan map[int]map[int]actions.Func

	if action == "enable" {
		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      actions.DisableOrEnableOverscan,
					Reference: con.a.DisableOrEnableOverscan,
					Argument: []interface{}{
						actions.EnableOrDisableConfig{
							DirOrFilePath: BootConfig,
							Action:        action,
						},
					},
				},
			},
		}
	} else if action == "disable" {
		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      actions.DisableOrEnableOverscan,
					Reference: con.a.DisableOrEnableOverscan,
					Argument: []interface{}{
						actions.EnableOrDisableConfig{
							DirOrFilePath: BootConfig,
							Action:        action,
						},
					},
				},
			},
			2: {
				1: {
					Name:      actions.CommentOverscan,
					Reference: con.a.CommentOverscan,
					Argument: []interface{}{
						actions.CommentOrUncommentConfig{
							DirOrFilePath: BootConfig,
							Action:        "comment",
						},
					},
				},
			},
		}
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable overscan failed")
	}

	return con.consys.ExecuteOV(plan)
}
