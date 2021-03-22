package install

import (
	"fmt"

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
