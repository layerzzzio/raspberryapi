package general

import (
	"fmt"
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

// ExecuteSASO starts/stops a service and returns an action.
func (gen *General) ExecuteSASO(action string, service string) (rpi.Action, error) {
	command := fmt.Sprintf("systemctl %v %v", action, service)

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

	return gen.gensys.ExecuteSASO(plan)
}
