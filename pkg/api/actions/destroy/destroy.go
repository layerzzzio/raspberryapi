package destroy

import (
	"fmt"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteDF delete file(s) and returns an action.
func (des *Destroy) ExecuteDF(path string) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.DeleteFile,
				Reference: des.a.DeleteFile,
				Argument: []interface{}{
					actions.FileOrDirectory{
						Path: path,
					},
				},
			},
		},
	}

	return des.dessys.ExecuteDF(plan)
}

// ExecuteSUS stop a user session and returns an action.
func (des *Destroy) ExecuteSUS(processname string, processtype string) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.KillProcessByName,
				Reference: des.a.KillProcessByName,
				Argument: []interface{}{
					actions.KPBN{
						Processname: processname,
						Processtype: processtype,
					},
				},
			},
		},
	}

	return des.dessys.ExecuteSUS(plan)
}

// ExecuteKP kill a process and returns an action.
func (des *Destroy) ExecuteKP(pid int) (rpi.Action, error) {
	plan := map[int](map[int]actions.Func){
		1: {
			1: {
				Name:      actions.KillProcess,
				Reference: des.a.KillProcess,
				Argument: []interface{}{
					actions.KP{
						Pid: fmt.Sprint(pid),
					},
				},
			},
		},
	}

	return des.dessys.ExecuteKP(plan)
}
