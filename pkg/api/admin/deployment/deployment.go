package deployment

import (
	"fmt"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteDPTOOL deploys a specific version on the device.
func (d *Deployment) ExecuteDPTOOL(deployType string, url string, version string) (rpi.Action, error) {
	deployScript := "/tmp/deploy_apis.sh"

	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: d.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf(
							"wget -nv %v -O %v",
							url,
							deployScript,
						),
					},
				},
			},
		},
		2: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: d.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf(
							"chmod 755 %v",
							deployScript,
						),
					},
				},
			},
		},
		3: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: d.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf(
							"%v %v %v",
							deployScript,
							deployType,
							version,
						),
					},
				},
			},
		},
		4: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: d.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf(
							"rm -f %v",
							deployScript,
						),
					},
				},
			},
		},
	}

	return d.dsys.ExecuteDPTOOL(plan)
}
