package destroy

import (
	"fmt"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// ExecuteDF delete file(s) and returns an action.
func (des *Destroy) ExecuteDF(path string) (rpi.Action, error) {
	// populate execs in the right execution order
	// always start with 1, not with 0
	execs := map[int]rpi.Exec{
		1: des.a.DeleteFile(path),
	}

	return des.dessys.ExecuteDF(execs)
}

// ExecuteSUS stop a user session and returns an action.
func (des *Destroy) ExecuteSUS(processname string, processtype string) (rpi.Action, error) {
	// define the execution plan
	// the below execution plan contains only one step
	// this one step has only one substep
	execPlan := map[int](map[int]actions.Func){
		1: {
			// no dependency for this task
			// if there was one, the name of the limiting function
			// should be the dependency value
			1: {
				Name:    actions.KillProcessByName,
				Pointer: des.a.KillProcessByName,
				Argument: actions.KPBN{
					Processname: processname,
					Processtype: processtype,
				},
			},
		},
	}

	return des.dessys.ExecuteSUS(execPlan)
}

// ExecuteKP kill a process and returns an action.
func (des *Destroy) ExecuteKP(pid int) (rpi.Action, error) {
	// populate execs in the right execution order
	// always start with 1, not with 0
	execs := map[int]rpi.Exec{
		1: des.a.KillProcess(fmt.Sprint(pid)),
	}

	return des.dessys.ExecuteKP(execs)
}
