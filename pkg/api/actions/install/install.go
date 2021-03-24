package install

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteAG install a package with apt-get
func (ins *Install) ExecuteAG(action string, pkg string) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
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

// ExecuteNV install NordVPN
func (ins *Install) ExecuteNV(action string) (rpi.Action, error) {
	var plan map[int](map[int]actions.Func)

	etcDir := "/etc/openvpn/nordvpn"
	nordVPNURl := "https://downloads.nordcdn.com/configs/archives/servers/ovpn.zip"

	if action == "install" {
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
							Command: fmt.Sprintf("if [ ! -f %v/ovpn.zip ] ; then wget %v -P %v ; fi", etcDir, nordVPNURl, etcDir),
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
							Command: fmt.Sprintf("unzip -o %v/ovpn.zip -d %v", etcDir, etcDir),
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

	return ins.inssys.ExecuteNV(plan)
}
