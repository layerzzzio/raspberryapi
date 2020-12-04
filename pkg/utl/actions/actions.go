package actions

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/shomali11/parallelizer"
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

func Call(funcName interface{}, params []interface{}) (result interface{}, err error) {
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

// Parallelize parallelizes the function calls
func Parallelize(functions []func()) error {
	return ParallelizeContext(context.Background(), functions...)
}

// ParallelizeTimeout parallelizes the function calls with a timeout
func ParallelizeTimeout(timeout time.Duration, functions []func()) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return ParallelizeContext(ctx, functions...)
}

// ParallelizeContext parallelizes the function calls with a context
func ParallelizeContext(ctx context.Context, functions ...func()) error {
	group := parallelizer.NewGroup()
	for _, function := range functions {
		group.Add(function)
	}

	return group.Wait(parallelizer.WithContext(ctx))
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
	var arguments []interface{}

	for kp, parentExec := range execPlan {
		index = fmt.Sprint(kp)
		childExecFunc := make([]func(), 0, len(parentExec))

		for kc, childExec := range parentExec {
			fmt.Println("----------------> " + childExec.Name)
			fmt.Println("----------------> " + fmt.Sprint(kc))

			// creating the func that will be added to the WaitGroup
			funChild := func() {
				i, err := strconv.Atoi(index)
				if err != nil {
					panic("cannot convert index to integer")
				}

				fmt.Println("----------------> " + fmt.Sprint(kc))

				// adding arguments here in absolutely essentials
				// it allows the params to be sorted in the right order
				arguments = append(arguments, childExec.Argument...)

				// intermediate arguments can only be extracted after first step is completed
				// and only if the current function has one or multiple dependencies
				// dep should be formatted the same way as index+Separator+fmt.Sprint(kc)
				if len(childExec.Dependency) > 0 && i > 1 {
					for _, dep := range childExec.Dependency {
						arguments = append(arguments, progress[dep].Stdout)
					}
				}

				fmt.Println("-------------------------------- " + childExec.Name)
				fmt.Println(childExec.Argument)
				fmt.Println(childExec.Name)
				fmt.Println(childExec.Pointer)
				fmt.Println(arguments)

				res, _ := Call(childExec.Pointer, arguments)
				resExec := res.(rpi.Exec)
				exitStatus = resExec.ExitStatus
				progress[index+Separator+fmt.Sprint(kc)] = resExec
			}
			childExecFunc = append(childExecFunc, funChild)
			// index and intermediateRes are re-initialized at every turn
			arguments = nil

			if exitStatus != 0 {
				break
			}
		}

		Parallelize(childExecFunc)

		if exitStatus != 0 {
			break
		}
	}

	return progress, exitStatus
}
