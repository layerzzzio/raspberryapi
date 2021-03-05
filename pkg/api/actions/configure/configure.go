package configure

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/constants"
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
						TargetFile: con.i.GetConfigFiles()["hosts"].Path,
						Data:       hostname,
					},
				},
			},
			2: {
				Name:      actions.ChangeHostnameInHostnameFile,
				Reference: con.a.ChangeHostnameInHostnameFile,
				Argument: []interface{}{
					actions.DataToFile{
						TargetFile: con.i.GetConfigFiles()["hostname"].Path,
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
						DirOrFilePath: constants.DHCPSERVICE,
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
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
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
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
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
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
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

// ExecuteBL enable or disable blanking
func (con *Configure) ExecuteBL(action string) (rpi.Action, error) {
	var plan map[int]map[int]actions.Func

	if action == "enable" || action == "disable" {
		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      actions.DisableOrEnableBlanking,
					Reference: con.a.DisableOrEnableBlanking,
					Argument: []interface{}{
						actions.TargetDestEnableOrDisableConfig{
							TargetDirOrFilePath:      constants.RASPICONFIGX11SERVICE,
							DestinationDirOrFilePath: constants.X11SERVICE,
							Action:                   action,
						},
					},
				},
			},
		}
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable blanking failed")
	}

	return con.consys.ExecuteBL(plan)
}

// ExecuteAUS add user
func (con *Configure) ExecuteAUS(username string, password string) (rpi.Action, error) {

	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.AddUser,
				Reference: con.a.AddUser,
				Argument: []interface{}{
					actions.ADU{
						Username: username,
						Password: password,
					},
				},
			},
		},
	}

	return con.consys.ExecuteAUS(plan)
}

// ExecuteDUS delete user
func (con *Configure) ExecuteDUS(username string) (rpi.Action, error) {

	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.DeleteUser,
				Reference: con.a.DeleteUser,
				Argument: []interface{}{
					actions.ADU{
						Username: username,
					},
				},
			},
		},
	}

	return con.consys.ExecuteDUS(plan)
}
