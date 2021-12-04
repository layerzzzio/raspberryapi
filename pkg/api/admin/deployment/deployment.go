package deployment

import (
	"fmt"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteDPTOOL deploys a specific version on the device.
func (d *Deployment) ExecuteDPTOOL(url string, version string) (rpi.Action, error) {
	prefix := "raspibuddy_deploy"
	releaseName := prefix + "-" + version
	releaseDir := "/usr/bin"

	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: d.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: fmt.Sprintf(
							"rm %v/%v-*",
							releaseDir,
							prefix,
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
							"wget -nv %v/%v -O %v/%v",
							url,
							releaseName,
							releaseDir,
							releaseName,
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
							"chmod 755 %v/%v",
							releaseDir,
							releaseName,
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
							"rm %v/%v ; ln -s %v/%v %v/%v",
							releaseDir,
							prefix,
							releaseDir,
							releaseName,
							releaseDir,
							prefix,
						),
					},
				},
			},
		},
	}

	return d.dsys.ExecuteDPTOOL(plan)
}
