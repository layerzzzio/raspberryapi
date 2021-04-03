package appaction

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteWOVA AppAction a vpn that works with OVPN
func (aac *AppAction) ExecuteWOVA(
	action string,
	vpnName string,
	country string,
	username string,
	password string,
) (rpi.Action, error) {
	var plan = make(map[int](map[int]actions.Func))
	var connect string
	if action == "connect" {
		configFiles := aac.i.VPNConfigFiles(
			vpnName,
			"/etc/openvpn/wov_"+vpnName+"/vpnconfigs",
			country,
		)
		randomIndex := rand.Intn(len(configFiles))
		configFile := configFiles[randomIndex]
		connectCommand := fmt.Sprintf(
			"openvpn --config %v --auth-user-pass <(echo -e \"%v\n%v\")",
			configFile,
			username,
			password)
		connect = fmt.Sprintf("nohup %v > /dev/null 2>&1 &", connectCommand)

		plan = map[int](map[int]actions.Func){
			1: {
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
		}
	} else if action == "disconnect" {
		regex := `openvpn --config\s*.*--auth-user-pass`
		pids := aac.i.ProcessesPids(regex)
		for k, pid := range pids {
			plan[1] = map[int]actions.Func{
				k: {
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
