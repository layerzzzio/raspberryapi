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
					Reference: con.a.DisableOrEnableConfig,
					Argument: []interface{}{
						actions.EODC{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							Data:          "disable_overscan=0",
							Regex:         actions.DisableOrEnableOverscanRegex,
							FunctionName:  actions.DisableOrEnableOverscan,
							AssetFile:     "../assets/config.txt",
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
					Reference: con.a.DisableOrEnableConfig,
					Argument: []interface{}{
						actions.EODC{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							Data:          "#disable_overscan=1",
							Regex:         actions.DisableOrEnableOverscanRegex,
							FunctionName:  actions.DisableOrEnableOverscan,
							AssetFile:     "../assets/config.txt",
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

// ExecuteCA disables or enables camera interface
func (con *Configure) ExecuteCA(action string) (rpi.Action, error) {
	var plan map[int]map[int]actions.Func

	if action == "enable" {
		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      "comment_startx",
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							DefaultData:   "#startx=",
							Regex:         actions.StartxCameraRegex,
							FunctionName:  actions.CameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
				2: {
					Name:      "comment_fixup_file",
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							DefaultData:   "#fixup_file=",
							Regex:         actions.FixupFileCameraRegex,
							FunctionName:  actions.CameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
				3: {
					Name:      actions.DisableOrEnableCameraInterface,
					Reference: con.a.DisableOrEnableConfig,
					Argument: []interface{}{
						actions.EODC{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							Data:          "start_x=1",
							Regex:         actions.DisableOrEnableCameraRegex,
							FunctionName:  actions.CameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
				4: {
					Name:      "set_gpu_mem_for_camera",
					Reference: con.a.SetVariableInConfigFile,
					Argument: []interface{}{
						actions.SVICF{
							File:      con.i.GetConfigFiles()["bootconfig"].Path,
							Regex:     actions.GpuMemCameraRegex,
							Data:      "gpu_mem=128",
							AssetFile: "../assets/config.txt",
						},
					},
				},
			},
		}
	} else if action == "disable" {
		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      "comment_startx",
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							DefaultData:   "#startx=",
							Regex:         actions.StartxCameraRegex,
							FunctionName:  actions.CameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
				2: {
					Name:      "comment_fixup_file",
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							DefaultData:   "#fixup_file=",
							Regex:         actions.FixupFileCameraRegex,
							FunctionName:  actions.CameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
				3: {
					Name:      actions.DisableOrEnableCameraInterface,
					Reference: con.a.DisableOrEnableConfig,
					Argument: []interface{}{
						actions.EODC{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							Data:          "start_x=0",
							Regex:         actions.DisableOrEnableCameraRegex,
							FunctionName:  actions.CameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
				4: {
					Name:      "comment_start_file",
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							DefaultData:   "#start_file=",
							Regex:         actions.StartFileCameraRegex,
							FunctionName:  actions.CameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
			},
		}
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable camera failed")
	}

	return con.consys.ExecuteCA(plan)
}
