package appaction

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteWOVA AppAction a vpn that works with OVPN
func (aac *AppAction) ExecuteWOVA(
	action string,
	vpnName string,
	url string,
) (rpi.Action, error) {
	var plan map[int](map[int]actions.Func)

	etcDir := fmt.Sprintf("/etc/openvpn/wov_%v", vpnName)

	if action == "install" {
		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: aac.a.ExecuteBashCommand,
					Argument: []interface{}{
						actions.EBC{
							Command: fmt.Sprintf("mkdir -p %v", etcDir),
						},
					},
				},
			},
			2: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: aac.a.ExecuteBashCommand,
					Argument: []interface{}{
						actions.EBC{
							Command: fmt.Sprintf(
								"if [ ! -f %v/vpnconfigs.zip ] ; then wget -cO %v/vpnconfigs.zip %v ; fi",
								etcDir,
								etcDir,
								url,
							),
						},
					},
				},
			},
			3: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: aac.a.ExecuteBashCommand,
					Argument: []interface{}{
						actions.EBC{
							Command: fmt.Sprintf(
								"unzip -o %v/vpnconfigs.zip -d %v/vpnconfigs",
								etcDir,
								etcDir,
							),
						},
					},
				},
			},
		}
	} else if action == "purge" {
		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: aac.a.ExecuteBashCommand,
					Argument: []interface{}{
						actions.EBC{
							Command: fmt.Sprintf("rm -rRf %v", etcDir),
						},
					},
				},
			},
		}
	} else {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: install or purge nordvpn failed")
	}

	return aac.aacsys.ExecuteWOVA(plan)
}
