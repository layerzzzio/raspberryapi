package configure

import (
	"fmt"
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
					Name:      actions.CommentOrUncommentInFile,
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        "comment",
							DefaultData:   "#startx=",
							Regex:         actions.StartxCameraRegex,
							FunctionName:  "comment_startx",
							AssetFile:     "../assets/config.txt",
						},
					},
				},
			},
			2: {
				1: {
					Name:      actions.CommentOrUncommentInFile,
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        "comment",
							DefaultData:   "#fixup_file=",
							Regex:         actions.FixupFileCameraRegex,
							FunctionName:  "comment_fixup_file",
							AssetFile:     "../assets/config.txt",
						},
					},
				},
			},
			3: {
				1: {
					Name:      actions.DisableOrEnableConfig,
					Reference: con.a.DisableOrEnableConfig,
					Argument: []interface{}{
						actions.EODC{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							Data:          "start_x=1",
							Regex:         actions.DisableOrEnableCameraRegex,
							FunctionName:  actions.DisableOrEnableCameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
			},
			4: {
				1: {
					Name:      actions.SetVariableInConfigFile,
					Reference: con.a.SetVariableInConfigFile,
					Argument: []interface{}{
						actions.SVICF{
							File:      con.i.GetConfigFiles()["bootconfig"].Path,
							Regex:     actions.GpuMemRegex,
							Data:      "gpu_mem=128",
							Threshold: "128",
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
					Name:      actions.CommentOrUncommentInFile,
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        "comment",
							DefaultData:   "#startx=",
							Regex:         actions.StartxCameraRegex,
							FunctionName:  "comment_startx",
							AssetFile:     "../assets/config.txt",
						},
					},
				},
			},
			2: {
				1: {
					Name:      actions.CommentOrUncommentInFile,
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        "comment",
							DefaultData:   "#fixup_file=",
							Regex:         actions.FixupFileCameraRegex,
							FunctionName:  "comment_fixup_file",
							AssetFile:     "../assets/config.txt",
						},
					},
				},
			},
			3: {
				1: {
					Name:      actions.DisableOrEnableConfig,
					Reference: con.a.DisableOrEnableConfig,
					Argument: []interface{}{
						actions.EODC{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        action,
							Data:          "start_x=0",
							Regex:         actions.DisableOrEnableCameraRegex,
							FunctionName:  actions.DisableOrEnableCameraInterface,
							AssetFile:     "../assets/config.txt",
						},
					},
				},
			},
			4: {
				1: {
					Name:      actions.CommentOrUncommentInFile,
					Reference: con.a.CommentOrUncommentInFile,
					Argument: []interface{}{
						actions.COUSLINF{
							DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
							Action:        "comment",
							DefaultData:   "#start_file=",
							Regex:         actions.StartFileCameraRegex,
							FunctionName:  "comment_start_file",
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

// ExecuteSSH enable or disable ssh
func (con *Configure) ExecuteSSH(action string) (rpi.Action, error) {
	var plan map[int]map[int]actions.Func
	var command string

	if action == "enable" {
		command = "ssh-keygen -A && update-rc.d ssh enable && invoke-rc.d ssh start"
	} else if action == "disable" {
		command = "update-rc.d ssh disable && invoke-rc.d ssh stop"
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable ssh failed")
	}

	plan = map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{Command: command},
				},
			},
		},
	}

	return con.consys.ExecuteSSH(plan)
}

// ExecuteVNC enable or disable vnc
func (con *Configure) ExecuteVNC(action string) (rpi.Action, error) {
	var plan map[int]map[int]actions.Func
	var command string

	if action == "enable" {
		command = "systemctl enable vncserver-x11-serviced.service && systemctl start vncserver-x11-serviced.service"
	} else if action == "disable" {
		command = "systemctl disable vncserver-x11-serviced.service && systemctl stop vncserver-x11-serviced.service"
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable vnc failed")
	}

	plan = map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{Command: command},
				},
			},
		},
	}

	return con.consys.ExecuteVNC(plan)
}

// ExecuteSPI enable or disable spi
func (con *Configure) ExecuteSPI(action string) (rpi.Action, error) {
	var plan map[int]map[int]actions.Func
	var data string

	if action == "enable" {
		data = "on"
	} else if action == "disable" {
		data = "off"
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable spi failed")
	}

	blacklist := "/etc/modprobe.d/raspi-blacklist.conf"
	sedBlacklist := "s/^\\(blacklist[[:space:]]*spi[-_]bcm2708\\)/#\\1/"

	plan = map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.DisableOrEnableConfig,
				Reference: con.a.DisableOrEnableConfig,
				Argument: []interface{}{
					actions.EODC{
						DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
						Action:        action,
						Data:          "dtparam=spi=" + data,
						Regex:         actions.DisableOrEnableSPIRegex,
						FunctionName:  actions.DisableOrEnableSPIInterface,
						AssetFile:     "../assets/config.txt",
					},
				},
			},
		},
		2: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf("if ! [ -e %v ]; then touch %v ; fi", blacklist, blacklist),
					},
				},
			},
		},
		3: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf("sed %v -i -e \"%v\"", blacklist, sedBlacklist),
					},
				},
			},
		},
		4: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: "dtparam spi=" + data,
					},
				},
			},
		},
	}

	return con.consys.ExecuteSPI(plan)
}

// ExecuteI2C enable or disable i2c
func (con *Configure) ExecuteI2C(action string) (rpi.Action, error) {
	var plan map[int]map[int]actions.Func
	var data string

	if action == "enable" {
		data = "on"
	} else if action == "disable" {
		data = "off"
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable i2c failed")
	}

	blacklist := "/etc/modprobe.d/raspi-blacklist.conf"
	sedBlacklist := "s/^\\(blacklist[[:space:]]*i2c[-_]bcm2708\\)/#\\1/"

	etcModules := "/etc/modules"
	setEtcModules := "s/^#[[:space:]]*\\(i2c[-_]dev\\)/\\1/"

	plan = map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.DisableOrEnableConfig,
				Reference: con.a.DisableOrEnableConfig,
				Argument: []interface{}{
					actions.EODC{
						DirOrFilePath: con.i.GetConfigFiles()["bootconfig"].Path,
						Action:        action,
						Data:          "dtparam=i2c_arm=" + data,
						Regex:         actions.DisableOrEnableI2CRegex,
						FunctionName:  actions.DisableOrEnableI2CInterface,
						AssetFile:     "../assets/config.txt",
					},
				},
			},
		},
		2: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf("if ! [ -e %v ]; then touch %v ; fi", blacklist, blacklist),
					},
				},
			},
		},
		3: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf("sed %v -i -e \"%v\"", blacklist, sedBlacklist),
					},
				},
			},
		},
		4: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf("sed %v -i -e %v", etcModules, setEtcModules),
					},
				},
			},
		},
		5: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf("if ! grep -q \"^i2c[-_]dev\" %v; then printf \"i2c-dev\n\" >> %v ; fi", etcModules, etcModules),
					},
				},
			},
		},
		6: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf("dtparam i2c_arm=%v", data),
					},
				},
			},
		},
		7: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: con.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: "modprobe i2c-dev",
					},
				},
			},
		},
	}

	return con.consys.ExecuteI2C(plan)
}
