package deployment

import (
	"fmt"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteDP deploys a specific version on the device.
func (d *Deployment) ExecuteDP(version string) (rpi.Action, error) {
	url := "https://api.github.com/repos/raspibuddy/rpi-release/releases/latest"
	
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: d.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf(
							"curl -s %v | grep \"%v\" | cut -d : -f 2,3 | tr -d \" | wget -qi -",
							url,
							version,
						),
					},
				},
			},
		},
	}
	return d.dsys.ExecuteDP(plan)
}
