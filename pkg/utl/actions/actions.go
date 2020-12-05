package actions

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/raspibuddy/rpi"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

var (
	// Separator separates parent and child execution
	Separator = "<|>"

	// DeleteFile is the name of the delete file exec
	DeleteFile = "delete_file"

	// KillProcess is the name of the kill process exec
	KillProcess = "kill_process"

	// KillProcessByName is the name of the kill process by name exec
	KillProcessByName = "kill_process_by_name"

	// StopUserSession is the name of the disconnect user action
	StopUserSession = "stop_user_sessions"
)

// Service represents several system scripts.
type Service struct{}

// Actions represents multiple system related action scripts.
type Actions interface{}

// New creates a service instance.
func New() *Service {
	return &Service{}
}

// Error is returned by Actions when the argument evaluation fails
type Error struct {
	// Name is the file name for which the error occurred.
	Arguments []string
}

func (e *Error) Error() string {
	return fmt.Sprintf("at least one argument is empty: %v", e.Arguments)
}

// Func represents a function to be called by function Call
// attribute Arguments content should be ordered
type Func struct {
	Name     string
	Pointer  interface{}
	Argument []interface{}
	// Example: "1" + action.Separator + "2" = "1<|>2"
	// Why not another function name ?
	// Reason : ensure uniqueness of the dependency
	Dependency []string
}

// KillProcess kill a given process
func (s Service) KillProcess(pid string) rpi.Exec {
	var stdErr string
	startTime := uint64(time.Now().Unix())
	pidNum, err := strconv.Atoi(pid)
	if err != nil {
		return rpi.Exec{
			Name:       KillProcess,
			StartTime:  startTime,
			EndTime:    uint64(time.Now().Unix()),
			ExitStatus: uint8(1),
			Stderr:     "pid is not an int",
		}
	} else {
		exitStatus := 0
		ps, _ := os.FindProcess(pidNum)
		e := ps.Kill()

		if e != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(e)
		}
		// execution end time
		endTime := uint64(time.Now().Unix())
		return rpi.Exec{
			Name:       KillProcess,
			StartTime:  startTime,
			EndTime:    endTime,
			ExitStatus: uint8(exitStatus),
			Stderr:     stdErr,
		}
	}
}

type KPBN struct {
	Processname string
	Processtype string
}

// KillProcessByName disconnect a user from an active tty from the current host
func (s Service) KillProcessByName(arg interface{}) (rpi.Exec, error) {
	processname := arg.(KPBN).Processname
	processtype := arg.(KPBN).Processtype

	if processname == "" || processtype == "" {
		return rpi.Exec{}, &Error{[]string{"processname", "processtype"}}
	}

	startTime := uint64(time.Now().Unix())

	exitStatus := 0
	var stdErr string

	var err error
	if processtype == "terminal" {
		_, err = exec.Command("sh", "-c", "pkill -t "+processname).Output()
	} else {
		_, err = exec.Command("sh", "-c", "pkill "+processname).Output()
	}

	if err != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(err)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       KillProcessByName,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// DeleteFile deletes a file or (empty) directory
func (s Service) DeleteFile(path string) rpi.Exec {
	// execution start time
	startTime := uint64(time.Now().Unix())

	exitStatus := 0
	var stdErr string
	e := os.Remove(path)
	if e != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(e)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       DeleteFile,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}
}

func Call(funcName interface{}, params []interface{}, wg *sync.WaitGroup) (result interface{}, err error) {
	defer wg.Done()

	f := reflect.ValueOf(funcName)
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	// var res []reflect.Value
	res := f.Call(in)
	result = res[0].Interface()
	return
}

func FlattenExecPlan(execPlan map[int](map[int]Func)) map[string]rpi.Exec {
	progress := map[string]rpi.Exec{}
	for kp, parentExec := range execPlan {
		parentIndex := fmt.Sprint(kp)
		for kc := range parentExec {
			childIndex := Separator + fmt.Sprint(kc)
			index := parentIndex + childIndex
			progress[index] = rpi.Exec{}
		}
	}

	return progress
}

func ExecuteExecPlan(execPlan map[int](map[int]Func), progress map[string]rpi.Exec) (map[string]rpi.Exec, uint8) {
	var exitStatus uint8
	var index string

	for kp, parentExec := range execPlan {
		index = fmt.Sprint(kp)
		var wg sync.WaitGroup

		for kc, childExec := range parentExec {
			wg.Add(1)

			i, err := strconv.Atoi(index)
			if err != nil {
				panic("cannot convert index to integer")
			}

			// adding arguments here in absolutely essentials
			// it allows the params to be sorted in the right order
			// arguments = append(arguments, childExec.Argument...)
			if len(childExec.Dependency) > 0 && i > 1 {
				for _, dep := range childExec.Dependency {
					childExec.Argument = append(childExec.Argument, progress[dep].Stdout)
				}
			}

			var res interface{}
			var errC error

			go func(childExec Func, kc int) {
				res, errC = Call(childExec.Pointer, childExec.Argument, &wg)
				if errC != nil {
					progress[index+Separator+fmt.Sprint(kc)] = rpi.Exec{
						Name:       childExec.Name,
						ExitStatus: 1,
						Stderr:     fmt.Sprint(errC),
					}
				} else {
					progress[index+Separator+fmt.Sprint(kc)] = res.(rpi.Exec)
				}
			}(childExec, kc)

		}

		wg.Wait()

		for _, v := range progress {
			if v.ExitStatus != 0 {
				exitStatus = 1
				break
			}
		}
	}

	return progress, exitStatus
}
