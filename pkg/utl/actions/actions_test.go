package actions_test

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/stretchr/testify/assert"
)

var (
	dummypath = "./dummyfile"
)

func TestDeleteFile(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		wantedData rpi.Exec
	}{
		{
			name: "error",
			path: "",
			wantedData: rpi.Exec{
				Name:       "delete_file",
				StartTime:  uint64(time.Now().Unix()),
				EndTime:    uint64(time.Now().Unix()),
				ExitStatus: 1,
				Stdin:      "",
				Stdout:     "",
				Stderr:     "remove : no such file or directory",
			},
		},
		{
			name: "success",
			path: dummypath,
			wantedData: rpi.Exec{
				Name:       "delete_file",
				StartTime:  uint64(time.Now().Unix()),
				EndTime:    uint64(time.Now().Unix()),
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "",
				Stderr:     "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := actions.New()
			createFile(dummypath)
			deletefile := a.DeleteFile(tc.path)
			assert.Equal(t, tc.wantedData, deletefile)
		})
	}
}

func TestKillProcess(t *testing.T) {
	cases := []struct {
		name         string
		convertIssue bool
		pidAlive     bool
		wantedData   rpi.Exec
	}{
		{
			name:         "error pid convertion issue",
			convertIssue: true,
			wantedData: rpi.Exec{
				Name:       "kill_process",
				StartTime:  uint64(time.Now().Unix()),
				EndTime:    uint64(time.Now().Unix()),
				ExitStatus: 1,
				Stdin:      "",
				Stdout:     "",
				Stderr:     "pid is not an int",
			},
		},
		{
			name:     "error killing process",
			pidAlive: false,
			wantedData: rpi.Exec{
				Name:       "kill_process",
				StartTime:  uint64(time.Now().Unix()),
				EndTime:    uint64(time.Now().Unix()),
				ExitStatus: 1,
				Stdin:      "",
				Stdout:     "",
				Stderr:     "os: process already finished",
			},
		},
		{
			name:     "success killing process",
			pidAlive: true,
			wantedData: rpi.Exec{
				Name:       "kill_process",
				StartTime:  uint64(time.Now().Unix()),
				EndTime:    uint64(time.Now().Unix()),
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "",
				Stderr:     "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := actions.New()
			cmd := exec.Command("bash", "sleep 10")
			err := cmd.Start()
			if err != nil {
				t.Fatalf("Failed to start test process: %v", err)
			}

			var largestfiles rpi.Exec

			if tc.convertIssue {
				largestfiles = a.KillProcess("ABC")
			} else {
				if tc.pidAlive {
					// process is still alive
					largestfiles = a.KillProcess(fmt.Sprint(cmd.Process.Pid))
					err = cmd.Wait()
					if err == nil {
						t.Errorf("Test process succeeded, but expected to fail")
					}
				} else {
					// process is dead
					err = cmd.Wait()
					if err == nil {
						t.Errorf("Test process succeeded, but expected to fail")
					}
					largestfiles = a.KillProcess(fmt.Sprint(cmd.Process.Pid))
				}
			}
			assert.Equal(t, tc.wantedData, largestfiles)
		})
	}
}

func TestKillProcessByName(t *testing.T) {
	cases := []struct {
		name             string
		argument         actions.KPBN
		wantedExitStatus uint8
		wantedStderr     string
	}{
		{
			name: "error killing process by its name",
			argument: actions.KPBN{
				Processname: "impossible_process_name",
				Processtype: "dummy",
			},
			wantedExitStatus: 1,
			wantedStderr:     "exit status 1",
		},
		{
			name: "error killing process by its name (terminal)",
			argument: actions.KPBN{
				Processname: "impossible_process_name",
				Processtype: "terminal",
			},
			wantedExitStatus: 1,
			wantedStderr:     "exit status 2",
		},
		{
			name: "error processname is empty",
			argument: actions.KPBN{
				Processname: "",
				Processtype: "terminal",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := actions.New()
			cmd := exec.Command("sh", "-c", "echo")
			err := cmd.Run()
			if err != nil {
				t.Fatalf("Failed to start test process: %v", err)
			}

			killProcessByName, _ := a.KillProcessByName(tc.argument)

			err = cmd.Wait()
			if err == nil {
				t.Errorf("Test process succeeded, but expected to fail")
			}

			assert.Equal(t, tc.wantedExitStatus, killProcessByName.ExitStatus)
			assert.Equal(t, tc.wantedStderr, killProcessByName.Stderr)
		})
	}
}

func createFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func TestFlattenExecPlan(t *testing.T) {
	cases := []struct {
		name       string
		execPlan   map[int](map[int]actions.Func)
		wantedData map[string]rpi.Exec
	}{
		{
			name: "success flatten exec plan",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name: "dummy_11",
					},
					2: {
						Name: "dummy_12",
					},
				},
				2: {
					1: {
						Name: "dummy_21",
					},
					2: {
						Name: "dummy_22",
					},
				},
			},
			wantedData: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"1" + actions.Separator + "2": {},
				"2" + actions.Separator + "1": {},
				"2" + actions.Separator + "2": {},
			},
		},
		{
			name: "success with another example",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:    "funcA",
						Pointer: funcA,
						Argument: []interface{}{
							ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedData: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			flattenExecPlan := actions.FlattenExecPlan(tc.execPlan)
			assert.Equal(t, tc.wantedData, flattenExecPlan)
		})
	}
}

func timeSleep(dur1 int, dur2 int) int {
	time.Sleep(time.Duration(dur1+dur2) * time.Second)
	return dur1 + dur2
}

func TestCall(t *testing.T) {
	cases := []struct {
		name       string
		funcName   interface{}
		params     []interface{}
		wantedData int
		wantedErr  error
	}{
		{
			name:      "error params out of index",
			funcName:  timeSleep,
			params:    []interface{}{1, 1, 1},
			wantedErr: errors.New("The number of params is out of index."),
		},
		{
			name:       "success calling function timeSleep",
			funcName:   timeSleep,
			params:     []interface{}{1, 1},
			wantedData: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)
			flattenExecPlan, err := actions.Call(tc.funcName, tc.params, &wg)
			if err != nil {
				assert.Equal(t, tc.wantedErr, err)
			} else {
				assert.Equal(t, tc.wantedData, flattenExecPlan.(int))
			}
		})
	}
}

func TestError(t *testing.T) {
	cases := []struct {
		name       string
		params     *actions.Error
		wantedData string
	}{
		{
			name:       "success",
			params:     &actions.Error{[]string{"dummy"}},
			wantedData: "at least one argument is empty: [dummy]",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := actions.Error{tc.params.Arguments}
			res := err.Error()
			assert.Equal(t, tc.wantedData, res)
		})
	}
}

type ArgFuncA struct {
	Arg0 string
	Arg1 string
}

func funcA(arg interface{}) (rpi.Exec, error) {
	arg0 := arg.(ArgFuncA).Arg0
	arg1 := arg.(ArgFuncA).Arg1

	if arg0 == "" || arg1 == "" {
		return rpi.Exec{}, &actions.Error{[]string{"arg0", "arg1"}}
	}

	stdOut := fmt.Sprintf("%v-%v", arg0, arg1)

	time.Sleep(1 * time.Second)

	res := rpi.Exec{
		Name:       "funcA",
		StartTime:  1,
		EndTime:    2,
		ExitStatus: 0,
		Stdin:      "",
		Stdout:     stdOut,
		Stderr:     "",
	}

	return res, nil
}

type ArgFuncB struct {
	Arg2 string
}

func funcB(arg interface{}) (rpi.Exec, error) {
	arg2 := arg.(ArgFuncB).Arg2

	if arg2 == "" {
		return rpi.Exec{}, &actions.Error{[]string{"arg2"}}
	}

	stdOut := fmt.Sprint(arg2)

	time.Sleep(2 * time.Second)

	res := rpi.Exec{
		Name:       "funcB",
		StartTime:  1,
		EndTime:    2,
		ExitStatus: 0,
		Stdin:      "",
		Stdout:     stdOut,
		Stderr:     "",
	}

	return res, nil
}

func TestExecuteExecPlanNoDependency(t *testing.T) {
	cases := []struct {
		name                 string
		execPlan             map[int](map[int]actions.Func)
		progress             map[string]rpi.Exec
		wantedDataExec       map[string]rpi.Exec
		wantedDataExitStatus uint8
	}{
		{
			name: "success : one parent | one child (funcA)",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:    "funcA",
						Pointer: funcA,
						Argument: []interface{}{
							ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "funcA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
			},
			wantedDataExitStatus: 0,
		},
		{
			name: "success : one parent | one child (funcB)",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:    "funcB",
						Pointer: funcB,
						Argument: []interface{}{
							ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "funcB",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string2",
				},
			},
			wantedDataExitStatus: 0,
		},
		{
			name: "success : one parent | two children",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:    "funcA",
						Pointer: funcA,
						Argument: []interface{}{
							ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
					2: actions.Func{
						Name:    "funcB",
						Pointer: funcB,
						Argument: []interface{}{
							ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"1" + actions.Separator + "2": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "funcA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"1" + actions.Separator + "2": {
					Name:       "funcB",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string2",
				},
			},
			wantedDataExitStatus: 0,
		},
		{
			name: "success : two parents | one child each",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:    "funcA",
						Pointer: funcA,
						Argument: []interface{}{
							ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:    "funcB",
						Pointer: funcB,
						Argument: []interface{}{
							ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"2" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "funcA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"2" + actions.Separator + "1": {
					Name:       "funcB",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string2",
				},
			},
			wantedDataExitStatus: 0,
		},
		{
			name: "success : two parents | two child each",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:    "funcA",
						Pointer: funcA,
						Argument: []interface{}{
							ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
					2: actions.Func{
						Name:    "funcA",
						Pointer: funcA,
						Argument: []interface{}{
							ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:    "funcB",
						Pointer: funcB,
						Argument: []interface{}{
							ArgFuncB{
								Arg2: "string2",
							},
						},
					},
					2: actions.Func{
						Name:    "funcB",
						Pointer: funcB,
						Argument: []interface{}{
							ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"1" + actions.Separator + "2": {},
				"2" + actions.Separator + "1": {},
				"2" + actions.Separator + "2": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "funcA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"1" + actions.Separator + "2": {
					Name:       "funcA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"2" + actions.Separator + "1": {
					Name:       "funcB",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string2",
				},
				"2" + actions.Separator + "2": {
					Name:       "funcB",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string2",
				},
			},
			wantedDataExitStatus: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			exec, exitStatus := actions.ExecuteExecPlan(tc.execPlan, tc.progress)
			assert.Equal(t, tc.wantedDataExec, exec)
			assert.Equal(t, tc.wantedDataExitStatus, exitStatus)
		})
	}
}
