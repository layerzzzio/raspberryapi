package general

import (
	"strings"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteRBS reboot/shutdown and returns an action.
func (gen *General) ExecuteRBS(option string) (rpi.Action, error) {
	command := "shutdown --poweroff now"

	if strings.ToLower(option) == "reboot" {
		command = "reboot now"
	}

	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.ExecuteBashCommand,
				Reference: gen.a.ExecuteBashCommand,
				Argument: []interface{}{
					actions.EBC{
						Command: command,
					},
				},
			},
		},
	}

	return gen.gensys.ExecuteRBS(plan)
}
