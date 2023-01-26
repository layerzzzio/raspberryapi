package appaction

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteWOVA AppAction a vpn that works with OVPN
func (aac *AppAction) ExecuteWOVA(
	action string,
	vpnName string,
	relativeConfigPath string,
	country string,
	username string,
	password string,
) (rpi.Action, error) {
	var plan = make(map[int](map[int]actions.Func))
	var connect string
	if action == "connect" {
		configFiles := aac.i.VPNConfigFiles(
			vpnName,
			"/etc/openvpn/wov_"+vpnName+"/vpnconfigs/"+relativeConfigPath,
			country,
		)
		randomIndex := rand.Intn(len(configFiles))
		configFile := configFiles[randomIndex]
		connectCommand := fmt.Sprintf(
			"bash -c 'openvpn --config %v --auth-user-pass <(echo -e \"%v\n%v\") &> /tmp/%v_authstatus.log'",
			configFile,
			username,
			password,
			vpnName)
		connect = fmt.Sprintf("nohup %v > /dev/null 2>&1 &", connectCommand)

		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: aac.a.ExecuteBashCommand,
					Argument: []interface{}{
						actions.EBC{
							Command: fmt.Sprintf(
								"if [ -f /tmp/%v_authstatus.log ]; then rm /tmp/%v_authstatus.log ; fi",
								vpnName,
								vpnName,
							),
						},
					},
				},
			},
			2: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: aac.a.ExecuteBashCommand,
					Argument: []interface{}{
						// https://stackoverflow.com/questions/38869427/openvpn-on-linux-passing-username-and-password-in-command-line
						actions.EBC{
							Command: connect,
						},
					},
				},
			},
			3: {
				1: {
					Name:      actions.ConfirmVPNAuthentication,
					Reference: aac.a.ConfirmVPNAuthentication,
					Argument: []interface{}{
						actions.CVPNAUTH{
							Filepath:  fmt.Sprintf("/tmp/%v_authstatus.log", vpnName),
							Timelimit: "15",
						},
					},
				},
			},
			4: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: aac.a.ExecuteBashCommand,
					Argument: []interface{}{
						actions.EBC{
							Command: fmt.Sprintf(
								"if [ -f /tmp/%v_authstatus.log ]; then rm /tmp/%v_authstatus.log ; fi",
								vpnName,
								vpnName,
							),
						},
					},
				},
			},
		}
	} else if action == "disconnect" {
		regex := `openvpn --config\s*.*--auth-user-pass`
		pids := aac.i.ProcessesPids(regex)
		// WARNING
		// if a disconnect command is ran when there are no PID related to openvpn
		// the action's execution that is returned looks empty like that
		// {
		//     "name": "action_vpn_with_ovpn",
		//     "numberOfSteps": 0,
		//     "executions": {},
		//     "exitStatus": 0,
		//     "startTime": 1638108695,
		//     "endTime": 1638108695
		// }
		for k, pid := range pids {
			plan[k+1] = map[int]actions.Func{
				1: {
					Name:      actions.KillProcess,
					Reference: aac.a.KillProcess,
					Argument: []interface{}{
						actions.KP{
							Pid: fmt.Sprint(pid),
						},
					},
				},
			}
		}
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: connect or disconnect vpn with openvpn failed")
	}

	return aac.aacsys.ExecuteWOVA(plan)
}
