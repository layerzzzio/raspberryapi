package appinstall

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteAG install a package with apt-get
func (ins *AppInstall) ExecuteAG(action string, pkg string) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: ins.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: "dpkg --configure -a",
					},
				},
			},
		},
		2: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: ins.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf("apt-get %v -y %v", action, pkg),
					},
				},
			},
		},
	}

	return ins.inssys.ExecuteAG(plan)
}

// ExecuteWOV install a vpn that works with OVPN
func (ins *AppInstall) ExecuteWOV(
	action string,
	vpnName string,
	url string,
) (rpi.Action, error) {
	var plan map[int](map[int]actions.Func)

	etcDir := fmt.Sprintf("/etc/openvpn/wov_%v", vpnName)

	if action == "install" {
		isOpenVPNInstalled := ins.i.IsDPKGInstalled("openvpn")

		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: ins.a.ExecuteBashCommand,
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
					Reference: ins.a.ExecuteBashCommand,
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
					Reference: ins.a.ExecuteBashCommand,
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

		// if openvpn is not installed, install it
		if !isOpenVPNInstalled {
			plan[4] = map[int]actions.Func{
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: ins.a.ExecuteBashCommand,
					Argument: []interface{}{
						actions.EBC{
							Command: "dpkg --configure -a",
						},
					},
				},
			}

			plan[5] = map[int]actions.Func{
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: ins.a.ExecuteBashCommand,
					Argument: []interface{}{
						actions.EBC{
							Command: "apt-get install -y openvpn",
						},
					},
				},
			}
		}
	} else if action == "purge" {
		plan = map[int](map[int]actions.Func){
			1: {
				1: {
					Name:      actions.ExecuteBashCommand,
					Reference: ins.a.ExecuteBashCommand,
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

	return ins.inssys.ExecuteWOV(plan)
}
