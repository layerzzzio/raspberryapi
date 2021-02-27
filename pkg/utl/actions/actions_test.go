package actions_test

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
	"github.com/shirou/gopsutil/host"
	"github.com/stretchr/testify/assert"
)

var (
	dummyfilepath      = "./dummyfile"
	dummydirectorypath = "./dummydirectory"
)

func TestDeleteFile(t *testing.T) {
	cases := []struct {
		name             string
		argument         interface{}
		wantedExitStatus uint8
		wantedStderr     string
		wantedErr        error
	}{
		{
			name:             "error : no such file or directory",
			argument:         actions.FileOrDirectory{Path: ""},
			wantedExitStatus: 1,
			wantedStderr:     "remove : no such file or directory",
			wantedErr:        nil,
		},
		{
			name: "error : too many arguments",
			argument: []actions.OtherParams{
				{Value: map[string]string{"path": dummyfilepath}},
				{Value: map[string]string{"dummy": dummyfilepath}},
			},
			wantedExitStatus: 1,
			wantedStderr:     "",
			wantedErr:        &actions.Error{Arguments: []string{"path"}},
		},
		{
			name: "success",
			argument: actions.OtherParams{
				Value: map[string]string{
					"path": dummyfilepath,
				},
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name:             "success",
			argument:         actions.FileOrDirectory{Path: dummyfilepath},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := actions.New()
			test_utl.CreateFile(dummyfilepath)
			deletefile, err := a.DeleteFile(tc.argument)
			assert.Equal(t, tc.wantedExitStatus, deletefile.ExitStatus)
			assert.Equal(t, tc.wantedStderr, deletefile.Stderr)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestKillProcess(t *testing.T) {
	cases := []struct {
		name             string
		convertIssue     bool
		pidAlive         bool
		otherParam       bool
		argument         interface{}
		wantedExitStatus uint8
		wantedStderr     string
		wantedErr        error
	}{
		{
			name:             "error too many arguments",
			convertIssue:     true,
			wantedExitStatus: 1,
			argument:         []actions.KP{{Pid: "ABC"}, {Pid: "ABC"}},
			wantedStderr:     "",
			wantedErr:        &actions.Error{Arguments: []string{"pid"}},
		},
		{
			name:             "error pid convertion issue",
			convertIssue:     true,
			wantedExitStatus: 1,
			argument:         actions.KP{Pid: "ABC"},
			wantedStderr:     "pid is not an int",
			wantedErr:        nil,
		},
		{
			name:             "error killing process",
			convertIssue:     false,
			pidAlive:         false,
			wantedExitStatus: 1,
			wantedStderr:     "os: process already finished",
			wantedErr:        nil,
		},
		{
			name:             "success killing process other params",
			convertIssue:     false,
			pidAlive:         true,
			otherParam:       true,
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name:             "success killing process",
			convertIssue:     false,
			pidAlive:         true,
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
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
				largestfiles, err = a.KillProcess(tc.argument)
			} else {
				if tc.pidAlive {
					// process is still alive
					if tc.otherParam {
						largestfiles, err = a.KillProcess(actions.OtherParams{
							Value: map[string]string{
								"pid": fmt.Sprint(cmd.Process.Pid),
							},
						})
						err := cmd.Wait()
						fmt.Println(err)
					} else {
						largestfiles, err = a.KillProcess(actions.KP{Pid: fmt.Sprint(cmd.Process.Pid)})
					}
				} else {
					// process is dead
					err = cmd.Wait()
					if err == nil {
						t.Errorf("Test process succeeded, but expected to fail")
					}
					largestfiles, err = a.KillProcess(actions.KP{Pid: fmt.Sprint(cmd.Process.Pid)})
				}
			}
			assert.Equal(t, tc.wantedExitStatus, largestfiles.ExitStatus)
			assert.Equal(t, tc.wantedStderr, largestfiles.Stderr)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestKillProcessByName(t *testing.T) {
	cases := []struct {
		name             string
		argument         interface{}
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
			wantedExitStatus: 1,
			wantedStderr:     "exit status 2",
		},
		{
			name: "error processtype is empty",
			argument: actions.KPBN{
				Processname: "",
				Processtype: "terminal",
			},
			wantedExitStatus: 1,
			wantedStderr:     "exit status 2",
		},
		{
			name: "error otherParam",
			argument: actions.OtherParams{
				Value: map[string]string{
					"processname": "impossible_process",
					"processtype": "terminal",
				},
			},
			wantedExitStatus: 1,
			wantedStderr:     "exit status 1",
		},
		{
			name:             "error wrong type",
			argument:         "dummy",
			wantedExitStatus: 1,
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
			// assert.Equal(t, tc.wantedStderr, killProcessByName.Stderr)
		})
	}
}
func TestFlattenPlan(t *testing.T) {
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
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
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
			FlattenPlan := actions.FlattenPlan(tc.execPlan)
			assert.Equal(t, tc.wantedData, FlattenPlan)
		})
	}
}

func TestCall(t *testing.T) {
	cases := []struct {
		name       string
		funcName   interface{}
		params     []interface{}
		wantedData rpi.Exec
		wantedErr  error
	}{
		{
			name:      "error params out of index",
			funcName:  test_utl.FuncB,
			params:    []interface{}{test_utl.ArgFuncB{Arg2: "string2"}, "dummy"},
			wantedErr: errors.New("The number of params is out of index."),
		},
		{
			name:     "success calling function FuncB",
			funcName: test_utl.FuncB,
			params:   []interface{}{test_utl.ArgFuncB{Arg2: "string2"}},
			wantedData: rpi.Exec{
				Name:       "FuncB",
				StartTime:  1,
				EndTime:    2,
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "string2",
				Stderr:     "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			call, err := actions.Call(tc.funcName, tc.params)
			fmt.Println(call)
			fmt.Println(err)
			if err != nil {
				assert.Equal(t, tc.wantedErr, err)
			} else {
				assert.Equal(t, tc.wantedData, call.(rpi.Exec))
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

func TestExecutePlanWithoutDependency(t *testing.T) {
	cases := []struct {
		name                 string
		execPlan             map[int](map[int]actions.Func)
		progress             map[string]rpi.Exec
		timeExpected         int
		wantedDataExec       map[string]rpi.Exec
		wantedDataExitStatus uint8
	}{
		{
			name: "success : one parent | one child (test_utl.FuncA)",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			timeExpected: 2,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncA",
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
			name: "success : one parent | one child (test_utl.FuncB)",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:      "FuncB",
						Reference: test_utl.FuncB,
						Argument: []interface{}{
							test_utl.ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			timeExpected: 1,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncB",
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
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
					2: actions.Func{
						Name:      "FuncB",
						Reference: test_utl.FuncB,
						Argument: []interface{}{
							test_utl.ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			timeExpected: 2,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"1" + actions.Separator + "2": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"1" + actions.Separator + "2": {
					Name:       "FuncB",
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
			name: "error : 4 parents | one child each | abort plan at step 2",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:      "FuncB",
						Reference: test_utl.FuncB,
						// too many arguments forcing an error
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
				3: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
				4: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			timeExpected: 2,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"2" + actions.Separator + "1": {},
				"3" + actions.Separator + "1": {},
				"4" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"2" + actions.Separator + "1": {
					Name:       "FuncB",
					StartTime:  0,
					EndTime:    0,
					ExitStatus: 1,
					Stdin:      "",
					Stderr:     "The number of params is out of index.",
					Stdout:     "",
				},
				"3" + actions.Separator + "1": {},
				"4" + actions.Separator + "1": {},
			},
			wantedDataExitStatus: 1,
		},
		{
			name: "success : two parents | one child each",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:      "FuncB",
						Reference: test_utl.FuncB,
						Argument: []interface{}{
							test_utl.ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			timeExpected: 3,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"2" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"2" + actions.Separator + "1": {
					Name:       "FuncB",
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
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
					2: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:      "FuncB",
						Reference: test_utl.FuncB,
						Argument: []interface{}{
							test_utl.ArgFuncB{
								Arg2: "string2",
							},
						},
					},
					2: actions.Func{
						Name:      "FuncB",
						Reference: test_utl.FuncB,
						Argument: []interface{}{
							test_utl.ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			timeExpected: 3,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"1" + actions.Separator + "2": {},
				"2" + actions.Separator + "1": {},
				"2" + actions.Separator + "2": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"1" + actions.Separator + "2": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"2" + actions.Separator + "1": {
					Name:       "FuncB",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string2",
				},
				"2" + actions.Separator + "2": {
					Name:       "FuncB",
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

	counter := 1
	for {
		fmt.Println("===========> Testing round " + fmt.Sprint(counter))
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				start := int(time.Now().Unix())
				exec, exitStatus := actions.ExecutePlan(tc.execPlan, tc.progress)
				fmt.Println("duration: " + fmt.Sprint(int(time.Now().Unix())-start))
				fmt.Println("timeExpected: " + fmt.Sprint(tc.timeExpected))
				assert.Equal(t, tc.wantedDataExec, exec)
				assert.Equal(t, tc.wantedDataExitStatus, exitStatus)
			})
		}
		if counter == 1 {
			break
		}
		counter += 1
	}
}

func TestExecutePlanWithDependency(t *testing.T) {
	cases := []struct {
		name                 string
		execPlan             map[int](map[int]actions.Func)
		progress             map[string]rpi.Exec
		timeExpected         int
		wantedDataExec       map[string]rpi.Exec
		wantedDataExitStatus uint8
	}{
		{
			name: "success : two parents | one child each | argument from previous step",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:      "FuncC",
						Reference: test_utl.FuncC,
						Dependency: actions.OtherParams{
							Value: map[string]string{
								"arg3": "1" + actions.Separator + "1",
							},
						},
					},
				},
			},
			timeExpected: 3,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"2" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"2" + actions.Separator + "1": {
					Name:       "FuncC",
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
			name: "error : two parents | one child each | argument from previous and current step",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:      "FuncC",
						Reference: test_utl.FuncC,
						Argument: []interface{}{
							test_utl.ArgFuncC{
								Arg3: "string3",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Dependency: actions.OtherParams{
							Value: map[string]string{
								"arg0": "1" + actions.Separator + "1",
								"arg1": "string1",
							},
						},
					},
				},
				3: {
					1: actions.Func{
						Name:      "FuncC",
						Reference: test_utl.FuncC,
						// too many arguments: forcing error
						Argument: []interface{}{
							test_utl.ArgFuncC{
								Arg3: "string3",
							},
							test_utl.ArgFuncC{
								Arg3: "string3",
							},
						},
					},
				},
				4: {
					1: actions.Func{
						Name:      "FuncC",
						Reference: test_utl.FuncC,
						Argument: []interface{}{
							test_utl.ArgFuncC{
								Arg3: "string3",
							},
						},
					},
				},
			},
			timeExpected: 3,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"2" + actions.Separator + "1": {},
				"3" + actions.Separator + "1": {},
				"4" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncC",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string3",
				},
				"2" + actions.Separator + "1": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string3-string1",
				},
				"3" + actions.Separator + "1": {
					Name:       "FuncC",
					StartTime:  0,
					EndTime:    0,
					ExitStatus: 1,
					Stdin:      "",
					Stderr:     "The number of params is out of index.",
					Stdout:     "",
				},
				"4" + actions.Separator + "1": {},
			},
			wantedDataExitStatus: 1,
		},
		{
			name: "success : two parents | one child each | argument from previous and current step",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:      "FuncC",
						Reference: test_utl.FuncC,
						Argument: []interface{}{
							test_utl.ArgFuncC{
								Arg3: "string3",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Dependency: actions.OtherParams{
							Value: map[string]string{
								"arg0": "1" + actions.Separator + "1",
								"arg1": "string1",
							},
						},
					},
				},
			},
			timeExpected: 3,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"2" + actions.Separator + "1": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncC",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string3",
				},
				"2" + actions.Separator + "1": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string3-string1",
				},
			},
			wantedDataExitStatus: 0,
		},
		{
			name: "success : two parents | two children each | argument from previous step",
			execPlan: map[int](map[int]actions.Func){
				1: {
					1: actions.Func{
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
					2: actions.Func{
						Name:      "FuncB",
						Reference: test_utl.FuncB,
						Argument: []interface{}{
							test_utl.ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
				2: {
					1: actions.Func{
						Name:      "FuncC",
						Reference: test_utl.FuncC,
						Dependency: actions.OtherParams{
							Value: map[string]string{
								"arg3": "1" + actions.Separator + "1",
							},
						},
					},
					2: actions.Func{
						Name:      "FuncC",
						Reference: test_utl.FuncC,
						Dependency: actions.OtherParams{
							Value: map[string]string{
								"arg3": "1" + actions.Separator + "2",
							},
						},
					},
				},
			},
			timeExpected: 3,
			progress: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {},
				"1" + actions.Separator + "2": {},
				"2" + actions.Separator + "1": {},
				"2" + actions.Separator + "2": {},
			},
			wantedDataExec: map[string]rpi.Exec{
				"1" + actions.Separator + "1": {
					Name:       "FuncA",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"1" + actions.Separator + "2": {
					Name:       "FuncB",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string2",
				},
				"2" + actions.Separator + "1": {
					Name:       "FuncC",
					StartTime:  1,
					EndTime:    2,
					ExitStatus: 0,
					Stdin:      "",
					Stderr:     "",
					Stdout:     "string0-string1",
				},
				"2" + actions.Separator + "2": {
					Name:       "FuncC",
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

	counter := 1
	for {
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				start := int(time.Now().Unix())
				exec, exitStatus := actions.ExecutePlan(tc.execPlan, tc.progress)
				fmt.Println("duration: " + fmt.Sprint(int(time.Now().Unix())-start))
				fmt.Println("timeExpected: " + fmt.Sprint(tc.timeExpected))
				assert.Equal(t, tc.wantedDataExec, exec)
				assert.Equal(t, tc.wantedDataExitStatus, exitStatus)
			})
		}
		if counter == 1 {
			break
		}
		counter += 1
	}
}

func TestBackupFile(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		perm       uint32
		wantedData error
	}{
		{
			name:       "success file does not exist",
			path:       "./dummyfile",
			perm:       0755,
			wantedData: nil,
		},
		{
			name:       "success file exists",
			path:       "./TestBackupFile",
			perm:       0755,
			wantedData: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name != "success file does not exist" {
				// create file
				file, err := os.Create(tc.path)
				if err != nil {
					log.Fatal(err)
				}

				// add text and close
				fmt.Fprintln(file, "hey_man")
				file.Close()

				// backup file
				backupFile := actions.BackupFile(tc.path, tc.perm)

				// open backup file and read the content
				fileBak, err := os.Open(tc.path + ".bak")
				if err != nil {
					log.Fatal(err)
				}

				var readLines = []string{}
				scanner := bufio.NewScanner(fileBak)
				for scanner.Scan() {
					readLines = append(readLines, scanner.Text())
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}

				fileBak.Close()

				// remove original file
				e := os.Remove(tc.path)
				if e != nil {
					fmt.Println(e)
				}

				// remove backup file
				e = os.Remove(tc.path + ".bak")
				if e != nil {
					fmt.Println(e)
				}

				assert.Equal(t, tc.wantedData, backupFile)
				// backup file content should be equal to original file
				assert.Equal(t, "hey_man", strings.Join(readLines, ""))
			}
			backupFile := actions.BackupFile(tc.path, tc.perm)
			assert.Equal(t, tc.wantedData, backupFile)
		})
	}
}

func TestApplyPermissionsToFile(t *testing.T) {
	cases := []struct {
		name             string
		path             string
		perm             uint32
		isChmodingFailed bool
		wantedData       error
		wantedPerm       os.FileMode
	}{
		{
			name:       "perm does not match regex",
			path:       "./dummyfile",
			perm:       1755,
			wantedData: nil,
			wantedPerm: os.FileMode(0644),
		},
		{
			name:       "success: simple perm matches regex",
			path:       "./dummyfile",
			perm:       0755,
			wantedData: nil,
			wantedPerm: os.FileMode(0755),
		},
		{
			name:       "success: another simple perm matches regex",
			path:       "./dummyfile",
			perm:       0100,
			wantedData: nil,
			wantedPerm: os.FileMode(0100),
		},
		{
			name:       "edge case: perm is 0",
			path:       "./dummyfile",
			perm:       0,
			wantedData: nil,
			wantedPerm: os.FileMode(0644),
		},
		{
			name:       "edge case: perm is 00",
			path:       "./dummyfile",
			perm:       00,
			wantedData: nil,
			wantedPerm: os.FileMode(0644),
		},
		{
			name:       "edge case: perm is 000",
			path:       "./dummyfile",
			perm:       000,
			wantedData: nil,
			wantedPerm: os.FileMode(0644),
		},
		{
			name:       "edge case: perm is 0000",
			path:       "./dummyfile",
			perm:       0000,
			wantedData: nil,
			wantedPerm: os.FileMode(0644),
		},
		{
			name:             "regex matches but chmoding failed",
			path:             "./dummyfile",
			perm:             0755,
			isChmodingFailed: true,
			wantedData:       fmt.Errorf("chmoding file failed"),
			wantedPerm:       os.FileMode(0000),
		},
		{
			name:             "regex does not match and chmoding failed",
			path:             "./dummyfile",
			perm:             1755,
			isChmodingFailed: true,
			wantedData:       fmt.Errorf("chmoding default file permissions failed"),
			wantedPerm:       os.FileMode(0000),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var applyPerm error
			var filePerm os.FileMode

			if tc.isChmodingFailed {
				applyPerm = actions.ApplyPermissionsToFile(tc.path, tc.perm)
			} else {
				// create file with perm 0666
				file, err := os.Create(tc.path)
				if err != nil {
					log.Fatal(err)
				}

				// add text and close
				fmt.Fprintln(file, "hey_man")
				file.Close()

				// apply permissions to file
				applyPerm = actions.ApplyPermissionsToFile(tc.path, tc.perm)

				// check the perm after applying them
				info, err := os.Stat(tc.path)
				if err != nil {
					log.Fatal(err)
				}
				filePerm = info.Mode()

				// remove file
				e := os.Remove(tc.path)
				if e != nil {
					log.Fatal(err)
				}
			}
			assert.Equal(t, tc.wantedData, applyPerm)
			assert.Equal(t, tc.wantedPerm, filePerm)
		})
	}
}

func TestOverwriteToFile(t *testing.T) {
	cases := []struct {
		name       string
		args       actions.WriteToFileArg
		isSuccess  bool
		wantedData error
	}{
		{
			name: "success with multiline",
			args: actions.WriteToFileArg{
				File:      "./test_write_to_file",
				Data:      []string{"text_1", "text_2", "text_3"},
				Multiline: true,
			},
			isSuccess:  true,
			wantedData: nil,
		},
		{
			name: "success not multiline",
			args: actions.WriteToFileArg{
				File:      "./test_write_to_file",
				Data:      []string{"text_1", "text_2", "text_3"},
				Multiline: false,
			},
			isSuccess:  true,
			wantedData: nil,
		},
		{
			name: "success permissions not nill",
			args: actions.WriteToFileArg{
				File:        "./test_write_to_file",
				Data:        []string{"text_1", "text_2", "text_3"},
				Multiline:   false,
				Permissions: 0755,
			},
			isSuccess:  true,
			wantedData: nil,
		},
		{
			name: "failure creating file",
			args: actions.WriteToFileArg{
				File:      "",
				Data:      []string{"text_1", "text_2", "text_3"},
				Multiline: false,
			},
			wantedData: fmt.Errorf("creating and opening file failed"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			overwriteToFile := actions.OverwriteToFile(tc.args)

			if tc.isSuccess {
				readLines, err := infos.New().ReadFile(tc.args.File)
				if err != nil {
					log.Fatal(err)
				}

				e := os.Remove(tc.args.File)
				if e != nil {
					fmt.Println(e)
				}

				// assert statements
				if tc.args.Multiline {
					assert.Equal(t, tc.args.Data, readLines)
				} else {
					assert.Equal(t, strings.Join(tc.args.Data, ""), strings.Join(readLines, ""))
				}
			}
			assert.Equal(t, tc.wantedData, overwriteToFile)
		})
	}
}

func TestGetReplaceType(t *testing.T) {
	cases := []struct {
		name       string
		repType    actions.ReplaceType
		wantedData *string
		wantedErr  error
	}{
		{
			name: "error: only one replace type required",
			repType: actions.ReplaceType{
				&actions.AllOccurrences{Occurrence: "occ", NewData: "new_data"},
				&actions.EntireLine{NewData: "new_data"},
			},
			wantedData: nil,
			wantedErr:  fmt.Errorf("only one replace type allowed"),
		},
		{
			name:       "error: at least one replace type required",
			repType:    actions.ReplaceType{},
			wantedData: nil,
			wantedErr:  fmt.Errorf("at least one replace type required"),
		},
		{
			name: "success: all_occurrences",
			repType: actions.ReplaceType{
				&actions.AllOccurrences{Occurrence: "occ", NewData: "new_data"},
				nil,
			},
			wantedData: &actions.RepTypeAllOccurrences,
			wantedErr:  nil,
		},
		{
			name: "success: entire_line",
			repType: actions.ReplaceType{
				nil,
				&actions.EntireLine{NewData: "new_data"},
			},
			wantedData: &actions.RepTypeEntireLine,
			wantedErr:  nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			repType, err := actions.GetReplaceType(tc.repType)
			assert.Equal(t, tc.wantedErr, err)
			assert.Equal(t, tc.wantedData, repType)
		})
	}
}

func TestReplaceLineInFile(t *testing.T) {
	cases := []struct {
		name          string
		args          actions.ReplaceLineInFileArg
		isSuccess     bool
		originalLines []string
		addLines      []string
		wantedLines   []string
		wantedData    error
	}{
		{
			name:      "success with replace type all_occurrences",
			isSuccess: true,
			args: actions.ReplaceLineInFileArg{
				File:        "./test_write_to_file",
				Permissions: 0755,
				Regex:       actions.HostnameChangeInHostsRegex,
				ReplaceType: actions.ReplaceType{
					&actions.AllOccurrences{
						Occurrence: "raspberrypi",
						NewData:    "new_hostname",
					},
					nil,
				},
				HasUniqueLines: true,
				ToAddIfNoMatch: []string{"127.0.1.1		new_hostname"},
			},
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
			},
			addLines: []string{
				"     127.0.1.1		raspberrypi",
			},
			wantedData: nil,
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"127.0.1.1		new_hostname",
			},
		},
		{
			name:      "success with replace type entire_line",
			isSuccess: true,
			args: actions.ReplaceLineInFileArg{
				File:        "./test_write_to_file",
				Permissions: 0755,
				Regex:       actions.HostnameChangeInHostsRegex,
				ReplaceType: actions.ReplaceType{
					nil,
					&actions.EntireLine{
						NewData: "new_hostname",
					},
				},
				HasUniqueLines: true,
				ToAddIfNoMatch: []string{"new_hostname"},
			},
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
			},
			addLines: []string{
				"        127.0.1.1		raspberrypi #random comment",
			},
			wantedData: nil,
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"new_hostname",
			},
		},
		{
			name:      "success: more matches than wanted",
			isSuccess: true,
			args: actions.ReplaceLineInFileArg{
				File:        "./test_write_to_file",
				Permissions: 0755,
				Regex:       actions.HostnameChangeInHostsRegex,
				ReplaceType: actions.ReplaceType{
					nil,
					&actions.EntireLine{
						NewData: "new_hostname",
					},
				},
				HasUniqueLines: false,
				ToAddIfNoMatch: []string{"new_hostname"},
			},
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
			},
			addLines: []string{
				"127.0.1.1		raspberrypi #random comment",
				"127.0.1.1		raspberrypi #random comment",
			},
			wantedData: nil,
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"new_hostname", // the first match
				"new_hostname", // the second match
				"new_hostname", // the default value to add
			},
		},
		{
			name:      "success but no replacement because no match",
			isSuccess: true,
			args: actions.ReplaceLineInFileArg{
				File:        "./test_write_to_file",
				Permissions: 0755,
				Regex:       actions.HostnameChangeInHostsRegex,
				ReplaceType: actions.ReplaceType{
					nil,
					&actions.EntireLine{
						NewData: "new_hostname",
					},
				},
			},
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
			},
			addLines: []string{
				"0.0.0.0		raspberrypi",
			},
			wantedData: nil,
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"0.0.0.0		raspberrypi",
			},
		},
		{
			name:      "success: no match but add default data & duplicates okay",
			isSuccess: true,
			args: actions.ReplaceLineInFileArg{
				File:        "./test_write_to_file",
				Permissions: 0755,
				Regex:       actions.HostnameChangeInHostsRegex,
				ReplaceType: actions.ReplaceType{
					nil,
					&actions.EntireLine{
						NewData: "new_hostname",
					},
				},
				ToAddIfNoMatch: []string{"127.0.1.1_xxx_new_hostname"},
			},
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
			},
			addLines: []string{
				"0.0.0.0		raspberrypi",
			},
			wantedData: nil,
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"0.0.0.0		raspberrypi",
				"127.0.1.1_xxx_new_hostname",
			},
		},
		{
			name:      "success: no match but add default data & uniques lines",
			isSuccess: true,
			args: actions.ReplaceLineInFileArg{
				File:        "./test_write_to_file",
				Permissions: 0755,
				Regex:       actions.HostnameChangeInHostsRegex,
				ReplaceType: actions.ReplaceType{
					nil,
					&actions.EntireLine{
						NewData: "new_hostname",
					},
				},
				ToAddIfNoMatch: []string{"127.0.1.1_xxx_new_hostname", "127.0.1.1_xxx_new_hostname"},
				HasUniqueLines: true,
			},

			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
			},
			addLines: []string{
				"0.0.0.0		raspberrypi",
				"0.0.0.0		raspberrypi",
			},
			wantedData: nil,
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"0.0.0.0		raspberrypi",
				"127.0.1.1_xxx_new_hostname",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isSuccess {
				// create and populate file
				if err := actions.OverwriteToFile(actions.WriteToFileArg{
					File:        tc.args.File,
					Data:        append(tc.originalLines, tc.addLines...),
					Multiline:   true,
					Permissions: 0755,
				}); err != nil {
					log.Fatal(err)
				}

				// replace line in file
				replaceLineInFile := actions.ReplaceLineInFile(tc.args)

				// read the new line
				readLines, err := infos.New().ReadFile(tc.args.File)
				if err != nil {
					log.Fatal(err)
				}

				// fmt.Println(readLines)
				// --> replace all occurrences:
				// [dummy line 1  dummy line 2 127.0.1.1  127.0.1.1		new_hostname  127.0.1.1		new_hostname  ]
				// --> replace entire line:
				// [dummy line 1  dummy line 2 127.0.1.1  new_hostname new_hostname ]

				if e := os.Remove(tc.args.File); e != nil {
					fmt.Println(e)
				}

				assert.Equal(t, tc.wantedLines, readLines)
				assert.Equal(t, tc.wantedData, replaceLineInFile)
			}
		})
	}
}

func TestChangeHostnameInHostnameFile(t *testing.T) {
	cases := []struct {
		name             string
		argument         interface{}
		isSuccess        bool
		originalLines    []string
		addLines         []string
		wantedLines      []string
		wantedExitStatus uint8
		wantedStderr     string
		wantedErr        error
	}{
		{
			name: "error : no such file or directory",
			argument: actions.DataToFile{
				TargetFile: "",
				Data:       "dummydata",
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "creating and opening file failed",
			wantedErr:        nil,
		},
		{
			name: "error : too many arguments",
			argument: []actions.OtherParams{
				{Value: map[string]string{"targetFile": dummyfilepath}},
				{Value: map[string]string{"hostname": dummyfilepath}},
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "",
			wantedErr:        &actions.Error{Arguments: []string{"hostname", "targetFile"}},
		},
		{
			name: "success with otherParams",
			argument: actions.OtherParams{
				Value: map[string]string{
					"targetFile": dummyfilepath,
					"hostname":   "new_hostname",
				},
			},
			isSuccess: true,
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
			},
			addLines: []string{
				"    127.0.1.1		raspberrypi",
			},
			wantedLines:      []string{"new_hostname"},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with regular params",
			argument: actions.DataToFile{
				TargetFile: dummyfilepath,
				Data:       "new_hostname",
			},
			isSuccess: true,
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
			},
			addLines: []string{
				"    127.0.1.1		raspberrypi",
			},
			wantedLines:      []string{"new_hostname"},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var chHostnameInHostnameFile rpi.Exec
			var err error
			a := actions.New()

			if tc.isSuccess {
				// create and populate file
				if err := actions.OverwriteToFile(actions.WriteToFileArg{
					File:        dummyfilepath,
					Data:        append(tc.originalLines, tc.addLines...),
					Multiline:   true,
					Permissions: 0755,
				}); err != nil {
					log.Fatal(err)
				}

				chHostnameInHostnameFile, err = a.ChangeHostnameInHostnameFile(tc.argument)

				// read the new line and delete
				readLines, err := infos.New().ReadFile(dummyfilepath)
				if err != nil {
					log.Fatal(err)
				}

				if e := os.Remove(dummyfilepath); e != nil {
					fmt.Println(e)
				}

				assert.Equal(t, tc.wantedLines, readLines)
			} else {
				chHostnameInHostnameFile, err = a.ChangeHostnameInHostnameFile(tc.argument)
			}

			assert.Equal(t, tc.wantedExitStatus, chHostnameInHostnameFile.ExitStatus)
			assert.Equal(t, tc.wantedStderr, chHostnameInHostnameFile.Stderr)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestChangeHostnameInHostsFile(t *testing.T) {
	info, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}

	cases := []struct {
		name             string
		argument         interface{}
		isSuccess        bool
		createFromAsset  bool
		originalLines    []string
		addLine          string
		wantedLines      []string
		wantedExitStatus uint8
		wantedStderr     string
		wantedErr        error
	}{
		{
			name: "error : no such file or directory",
			argument: actions.DataToFile{
				TargetFile: "",
				Data:       "dummydata",
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "creating and opening file failed",
			wantedErr:        nil,
		},
		{
			name: "error : too many arguments",
			argument: []actions.OtherParams{
				{Value: map[string]string{"targetFile": dummyfilepath}},
				{Value: map[string]string{"hostname": dummyfilepath}},
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "",
			wantedErr:        &actions.Error{Arguments: []string{"hostname", "targetFile"}},
		},
		{
			name: "success with otherParams",
			argument: actions.OtherParams{
				Value: map[string]string{
					"targetFile": dummyfilepath,
					"hostname":   "new_hostname",
				},
			},
			isSuccess: true,
			addLine: "     127.0.1.1		",
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"yessss man",
			},
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"yessss man",
				"127.0.1.1		new_hostname",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with regular params",
			argument: actions.DataToFile{
				TargetFile: dummyfilepath,
				Data:       "new_hostname",
			},
			isSuccess: true,
			addLine: "127.0.1.1		",
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"yessss man",
			},
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"yessss man",
				"127.0.1.1		new_hostname",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success: change hostname from asset",
			argument: actions.DataToFile{
				TargetFile: dummyfilepath,
				Data:       "new_hostname",
			},
			isSuccess:       true,
			createFromAsset: true,
			wantedLines: []string{
				"127.0.0.1	localhost",
				"",
				"::1		localhost ip6-localhost ip6-loopback",
				"",
				"ff02::1		ip6-allnodes",
				"",
				"ff02::2		ip6-allrouters",
				"127.0.1.1		new_hostname",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success: file does not exist",
			argument: actions.DataToFile{
				TargetFile: dummyfilepath,
				Data:       "new_hostname",
			},
			isSuccess: true,
			addLine: "127.0.1.1		",
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"yessss man",
			},
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"yessss man",
				"127.0.1.1		new_hostname",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with no match",
			argument: actions.DataToFile{
				TargetFile: dummyfilepath,
				Data:       "new_hostname",
			},
			isSuccess: true,
			addLine: "XXX.0.1.1		",
			originalLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"yessss man",
			},
			wantedLines: []string{
				"dummy line 1",
				"dummy line 2 127.0.1.1",
				"yessss man",
				"XXX.0.1.1		" + info.Hostname,
				"127.0.1.1		new_hostname",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var chHostnameInHostnameFile rpi.Exec
			var err error
			a := actions.New()

			if tc.isSuccess {
				// create and populate file
				if tc.createFromAsset == false {
					if err := actions.OverwriteToFile(actions.WriteToFileArg{
						File:        dummyfilepath,
						Data:        append(tc.originalLines, []string{tc.addLine + info.Hostname}...),
						Multiline:   true,
						Permissions: 0755,
					}); err != nil {
						log.Fatal(err)
					}
				}

				chHostnameInHostnameFile, err = a.ChangeHostnameInHostsFile(tc.argument)
				if err != nil {
					log.Fatal(err)
				}

				// read the new line and delete
				readLines, err := infos.New().ReadFile(dummyfilepath)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(readLines)

				if e := os.Remove(dummyfilepath); e != nil {
					fmt.Println(e)
				}

				assert.Equal(t, tc.wantedLines, readLines)
			} else {
				chHostnameInHostnameFile, err = a.ChangeHostnameInHostsFile(tc.argument)
			}

			assert.Equal(t, tc.wantedExitStatus, chHostnameInHostnameFile.ExitStatus)
			assert.Equal(t, tc.wantedStderr, chHostnameInHostnameFile.Stderr)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestAddLinesEndOfFile(t *testing.T) {
	cases := []struct {
		name       string
		arg        actions.WriteToFileArg
		isSuccess  bool
		wantedData []string
		wantedErr  error
	}{
		{
			name: "success",
			arg: actions.WriteToFileArg{
				File:      dummyfilepath,
				Data:      []string{"dummy line 4", "dummy line 5"},
				Multiline: true,
			},
			isSuccess: true,
			wantedData: []string{
				"dummy line 1",
				"dummy line 2",
				"dummy line 3",
				"dummy line 4",
				"dummy line 5",
			},
			wantedErr: nil,
		},
		{
			name: "error: file not found",
			arg: actions.WriteToFileArg{
				File:      dummyfilepath + "xxxx",
				Data:      []string{"dummy line 4", "dummy line 5"},
				Multiline: true,
			},
			isSuccess: false,
			wantedErr: fmt.Errorf("reading file failed"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			if tc.isSuccess {
				// create and populate file
				if err := actions.OverwriteToFile(actions.WriteToFileArg{
					File: dummyfilepath,
					Data: []string{
						"dummy line 1",
						"dummy line 2",
						"dummy line 3",
					},
					Multiline:   true,
					Permissions: 0755,
				}); err != nil {
					fmt.Print(err)
				}

				if err = actions.AddLinesEndOfFile(tc.arg); err != nil {
					fmt.Print(err)
				}

				// read the new line and delete
				readLines, err := infos.New().ReadFile(dummyfilepath)
				if err != nil {
					fmt.Print(err)
				}

				if err = os.Remove(dummyfilepath); err != nil {
					fmt.Print(err)
				}

				assert.Equal(t, tc.wantedData, readLines)

			} else {
				if err = actions.AddLinesEndOfFile(tc.arg); err != nil {
					fmt.Print(err)
				}
			}
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestWaitForNetworkAtBoot(t *testing.T) {
	cases := []struct {
		name             string
		argument         interface{}
		isSuccess        bool
		enable           bool
		wantedData       []string
		wantedErr        error
		wantedExitStatus uint8
		wantedStderr     string
	}{
		{
			name: "error : no such file or directory (enable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: "",
				Action:        actions.Enable,
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "creating and opening file failed",
			wantedErr:        nil,
		},
		{
			name: "error : no such file or directory (disable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: "",
				Action:        actions.Disable,
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "remove /wait.conf: no such file or directory",
			wantedErr:        nil,
		},
		{
			name: "error : bad action type",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummydirectorypath,
				Action:        "dummyactiontype",
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "bad action type",
			wantedErr:        nil,
		},
		{
			name: "error : too many arguments",
			argument: []actions.OtherParams{
				{Value: map[string]string{"directory": dummydirectorypath}},
				{Value: map[string]string{"action": "dummyaction"}},
				{Value: map[string]string{"dummyextraarg": "dummyextraarg"}},
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "",
			wantedErr:        &actions.Error{Arguments: []string{"directory", "action"}},
		},
		{
			name: "success enabling with otherParams",
			argument: actions.OtherParams{
				Value: map[string]string{
					"directory": dummydirectorypath,
					"action":    actions.Enable,
				},
			},
			isSuccess: true,
			enable:    true,
			wantedData: []string{
				"[Service]",
				"ExecStart=",
				"ExecStart=/usr/lib/dhcpcd5/dhcpcd -q -w",
			},
			wantedErr:        nil,
			wantedExitStatus: 0,
			wantedStderr:     "",
		},
		{
			name: "success enabling with regular params",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummydirectorypath,
				Action:        actions.Enable,
			},
			isSuccess: true,
			enable:    true,
			wantedData: []string{
				"[Service]",
				"ExecStart=",
				"ExecStart=/usr/lib/dhcpcd5/dhcpcd -q -w",
			},
			wantedErr:        nil,
			wantedExitStatus: 0,
			wantedStderr:     "",
		},
		{
			name: "success disable with otherParams",
			argument: actions.OtherParams{
				Value: map[string]string{
					"directory": dummydirectorypath,
					"action":    actions.Disable,
				},
			},
			isSuccess:        true,
			enable:           false,
			wantedErr:        nil,
			wantedExitStatus: 0,
			wantedStderr:     "",
		},
		{
			name: "success disable with regular params",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummydirectorypath,
				Action:        actions.Disable,
			},
			isSuccess:        true,
			enable:           false,
			wantedErr:        nil,
			wantedExitStatus: 0,
			wantedStderr:     "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var waitForNetworkAtBoot rpi.Exec
			var err error
			a := actions.New()

			if tc.isSuccess {
				if tc.enable == false {
					_ = os.MkdirAll(dummydirectorypath, 0755)
					if err := actions.OverwriteToFile(actions.WriteToFileArg{
						File:      "./" + dummydirectorypath + "/wait.conf",
						Data:      []string{"dummydata"},
						Multiline: true,
					}); err != nil {
						fmt.Println(err)
					}
				}

				waitForNetworkAtBoot, err = a.WaitForNetworkAtBoot(tc.argument)
				if err != nil {
					fmt.Println(err)
				}

				if tc.enable {
					// read the new line and delete
					readLines, err := infos.New().ReadFile("./" + dummydirectorypath + "/wait.conf")
					if err != nil {
						fmt.Println(err)
					}

					// fmt.Println(readLines)

					// assert statements
					assert.Equal(t, tc.wantedData, readLines)
				}

				if err = os.RemoveAll(dummydirectorypath); err != nil {
					fmt.Println(err)
				}

			} else {
				waitForNetworkAtBoot, err = a.WaitForNetworkAtBoot(tc.argument)
				if e := os.RemoveAll(dummydirectorypath); err != nil {
					fmt.Println(e)
				}
			}

			assert.Equal(t, tc.wantedExitStatus, waitForNetworkAtBoot.ExitStatus)
			assert.Equal(t, tc.wantedStderr, waitForNetworkAtBoot.Stderr)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestRemoveDuplicateStrings(t *testing.T) {
	cases := []struct {
		name       string
		strSlice   []string
		wantedData []string
	}{
		{
			name:       "success: one duplicate",
			strSlice:   []string{"A", "B", "B", "C", "D", "E"},
			wantedData: []string{"A", "B", "C", "D", "E"},
		},
		{
			name:       "success: five duplicates",
			strSlice:   []string{"A", "B", "B", "B", "B", "B", "B", "C", "D", "E"},
			wantedData: []string{"A", "B", "C", "D", "E"},
		},
		{
			name:       "success: different duplicates in the middle",
			strSlice:   []string{"A", "B", "B", "B", "B", "B", "B", "C", "D", "D", "D", "E"},
			wantedData: []string{"A", "B", "C", "D", "E"},
		},
		{
			name:       "success: different duplicates at the beginning, in the middle and at the end",
			strSlice:   []string{"A", "A", "B", "B", "B", "B", "B", "B", "C", "D", "D", "D", "E", "E"},
			wantedData: []string{"A", "B", "C", "D", "E"},
		},
		{
			name:       "success: np duplicates",
			strSlice:   []string{"A", "B", "C", "D", "E"},
			wantedData: []string{"A", "B", "C", "D", "E"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dup := actions.RemoveDuplicateStrings(tc.strSlice)
			assert.Equal(t, tc.wantedData, dup)
		})
	}
}

func TestDisableOrEnableOverscan(t *testing.T) {
	cases := []struct {
		name             string
		argument         interface{}
		isSuccess        bool
		createFromAsset  bool
		originalLines    []string
		addLines         []string
		wantedLines      []string
		wantedExitStatus uint8
		wantedStderr     string
		wantedErr        error
	}{
		{
			name: "error : no such file or directory (enable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: "",
				Action:        "enable",
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "creating and opening file failed",
			wantedErr:        nil,
		},
		{
			name: "error : no such file or directory (disable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: "",
				Action:        "disable",
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "creating and opening file failed",
			wantedErr:        nil,
		},
		{
			name: "error : too many arguments",
			argument: []actions.OtherParams{
				{Value: map[string]string{"path": dummyfilepath}},
				{Value: map[string]string{"action": "enable"}},
				{Value: map[string]string{"dummyarg": "dummyargvalue"}},
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "",
			wantedErr:        &actions.Error{Arguments: []string{"path", "action"}},
		},
		{
			name: "error : action not right",
			argument: actions.OtherParams{
				Value: map[string]string{
					"path":   dummyfilepath,
					"action": "enable-xxx",
				},
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "bad action type",
			wantedErr:        nil,
		},
		{
			name: "success with otherParams (enable)",
			argument: actions.OtherParams{
				Value: map[string]string{
					"path":   dummyfilepath,
					"action": "enable",
				},
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
			},
			addLines: []string{
				"     # disable_overscan = 1",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"disable_overscan=0",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with regular params (enable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "enable",
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
			},
			addLines: []string{
				"  #           disable_overscan      = 1 #random comment",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"disable_overscan=0",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success but no match (enable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "enable",
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
			},
			addLines: []string{
				"  #   #        disable_overscan      = 1 #random comment",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"  #   #        disable_overscan      = 1 #random comment",
				"disable_overscan=0",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with otherParams (disable)",
			argument: actions.OtherParams{
				Value: map[string]string{
					"path":   dummyfilepath,
					"action": "disable",
				},
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
			},
			addLines: []string{
				"    disable_overscan  =   0",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"#disable_overscan=1",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with regular params (disable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "disable",
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
			},
			addLines: []string{
				"    disable_overscan  =   0 # random comment",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"#disable_overscan=1",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with regular params but no match(disable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "disable",
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
			},
			addLines: []string{
				"  ABC  disable_overscan  =   0 # random comment",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"  ABC  disable_overscan  =   0 # random comment",
				"#disable_overscan=1",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with regular params (disable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "disable",
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
			},
			addLines: []string{
				"    disable_overscan  =   0 # random comment",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"#disable_overscan=1",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success: created from asset (disable)",
			argument: actions.EnableOrDisableConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "disable",
			},
			isSuccess:       true,
			createFromAsset: true,
			wantedLines: []string{
				"# For more options and information see",
				"# http://rpf.io/configtxt",
				"# Some settings may impact device functionality. See link above for details",
				"",
				"# uncomment if you get no picture on HDMI for a default \"safe\" mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"#disable_overscan=1",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"#overscan_right=16",
				"#overscan_top=16",
				"#overscan_bottom=16",
				"# uncomment to force a console size. By default it will be display's size minus",
				"# overscan.",
				"#framebuffer_width=1280",
				"#framebuffer_height=720",
				"# uncomment if hdmi display is not detected and composite is being output",
				"#hdmi_force_hotplug=1",
				"# uncomment to force a specific HDMI mode (this will force VGA)",
				"#hdmi_group=1",
				"#hdmi_mode=1",
				"# uncomment to force a HDMI mode rather than DVI. This can make audio work in",
				"# DMT (computer monitor) modes",
				"#hdmi_drive=2",
				"# uncomment to increase signal to HDMI, if you have interference, blanking, or",
				"# no display",
				"#config_hdmi_boost=4",
				"# uncomment for composite PAL",
				"#sdtv_mode=2",
				"#uncomment to overclock the arm. 700 MHz is the default.",
				"#arm_freq=800",
				"# Uncomment some or all of these to enable the optional hardware interfaces",
				"#dtparam=i2c_arm=on",
				"#dtparam=i2s=on",
				"#dtparam=spi=on",
				"# Uncomment this to enable infrared communication.",
				"#dtoverlay=gpio-ir,gpio_pin=17",
				"#dtoverlay=gpio-ir-tx,gpio_pin=18",
				"# Additional overlays and parameters are documented /boot/overlays/README",
				"# Enable audio (loads snd_bcm2835)",
				"dtparam=audio=on",
				"[pi4]",
				"# Enable DRM VC4 V3D driver on top of the dispmanx display stack",
				"dtoverlay=vc4-fkms-v3d",
				"max_framebuffers=2",
				"[all]",
				"#dtoverlay=vc4-fkms-v3d",
				"#disable_overscan=1",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var overscan rpi.Exec
			var err error
			a := actions.New()

			if tc.isSuccess {
				if tc.createFromAsset == false {
					// create and populate file
					if err := actions.OverwriteToFile(actions.WriteToFileArg{
						File:        dummyfilepath,
						Data:        append(tc.originalLines, tc.addLines...),
						Multiline:   true,
						Permissions: 0755,
					}); err != nil {
						log.Fatal(err)
					}
				}

				overscan, err = a.DisableOrEnableOverscan(tc.argument)
				if err != nil {
					log.Fatal(err)
				}

				// if err := actions.ApplyPermissionsToFile(dummyfilepath, 0755); err != nil {
				// 	log.Fatal(err)
				// }

				// read the new line and delete
				readLines, err := infos.New().ReadFile(dummyfilepath)
				if err != nil {
					log.Fatal(err)
				}

				if e := os.Remove(dummyfilepath); e != nil {
					fmt.Println(e)
				}

				assert.Equal(t, tc.wantedLines, readLines)
			} else {
				overscan, err = a.DisableOrEnableOverscan(tc.argument)
			}

			assert.Equal(t, tc.wantedExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedStderr, overscan.Stderr)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestCommentOrUncommentLineInFile(t *testing.T) {
	cases := []struct {
		name          string
		args          actions.CommentLineInFileArg
		isSuccess     bool
		originalLines []string
		addLines      []string
		wantedData    error
		wantedLines   []string
	}{
		{
			name:      "error: no file",
			isSuccess: false,
			args: actions.CommentLineInFileArg{
				File:        "",
				Permissions: 0755,
				Regex:       actions.CommentOverscanRegex,
				Action:      "dummyaction",
			},
			wantedData: fmt.Errorf("opening file failed"),
		},
		{
			name:      "error: action not right",
			isSuccess: false,
			args: actions.CommentLineInFileArg{
				File:        dummyfilepath,
				Permissions: 0755,
				Regex:       actions.CommentOverscanRegex,
				Action:      "dummyaction",
			},
			wantedData: fmt.Errorf("bad action: comment or uncomment"),
		},
		{
			name:      "success commenting",
			isSuccess: true,
			args: actions.CommentLineInFileArg{
				File:           dummyfilepath,
				Permissions:    0755,
				Regex:          actions.CommentOverscanRegex,
				Action:         actions.Comment,
				HasUniqueLines: false,
				ToAddIfNoMatch: []string{
					"#overscan_left=16",
					"#overscan_right=16",
					"#overscan_top=16",
					"#overscan_bottom=16",
				},
			},
			originalLines: []string{
				"this is line 1",
				"# this is line 2",
				"# and line 3",
			},
			addLines: []string{
				"         overscan_left=16",
				"overscan_right=16",
				"overscan_left=16",
			},
			wantedData: nil,
			wantedLines: []string{
				"this is line 1",
				"# this is line 2",
				"# and line 3",
				"#overscan_left=16",
				"#overscan_right=16",
				"#overscan_left=16",
				"#overscan_left=16",
				"#overscan_right=16",
				"#overscan_top=16",
				"#overscan_bottom=16",
			},
		},
		{
			name:      "success uncommenting",
			isSuccess: true,
			args: actions.CommentLineInFileArg{
				File:        dummyfilepath,
				Permissions: 0755,
				Regex:       actions.UncommentOverscanRegex,
				Action:      actions.Uncomment,
			},
			originalLines: []string{
				"this is line 1",
				"# this is line 2",
				"# and line 3",
			},
			addLines: []string{
				"   #  overscan_left=16 # another comment",
			},
			wantedData: nil,
			wantedLines: []string{
				"this is line 1",
				"# this is line 2",
				"# and line 3",
				"overscan_left=16 # another comment",
			},
		},
		{
			name:      "success but no replacement because no match",
			isSuccess: true,
			args: actions.CommentLineInFileArg{
				File:        dummyfilepath,
				Permissions: 0755,
				Regex:       actions.CommentOverscan,
			},
			originalLines: []string{
				"this is line 1",
				"# this is line 2",
				"# and line 3",
			},
			addLines: []string{
				" ## first comment   overscan_left=16 # another comment",
			},
			wantedData: nil,
			wantedLines: []string{
				"this is line 1",
				"# this is line 2",
				"# and line 3",
				" ## first comment   overscan_left=16 # another comment",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isSuccess {
				// create and populate file
				if err := actions.OverwriteToFile(actions.WriteToFileArg{
					File:        tc.args.File,
					Data:        append(tc.originalLines, tc.addLines...),
					Multiline:   true,
					Permissions: 0755,
				}); err != nil {
					log.Fatal(err)
				}

				// comment line in file
				commentLineInFile := actions.CommentOrUncommentLineInFile(tc.args)

				// read the new line
				readLines, err := infos.New().ReadFile(tc.args.File)
				if err != nil {
					log.Fatal(err)
				}

				if e := os.Remove(tc.args.File); e != nil {
					fmt.Println(e)
				}

				assert.Equal(t, tc.wantedLines, readLines)
				assert.Equal(t, tc.wantedData, commentLineInFile)
			}
		})
	}
}

func TestCommentOverscan(t *testing.T) {
	cases := []struct {
		name             string
		argument         interface{}
		isSuccess        bool
		createFromAsset  bool
		originalLines    []string
		addLines         []string
		wantedLines      []string
		wantedExitStatus uint8
		wantedStderr     string
		wantedErr        error
	}{
		{
			name: "error : no such file or directory",
			argument: actions.CommentOrUncommentConfig{
				DirOrFilePath: "",
				Action:        "comment",
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "creating and opening file failed",
			wantedErr:        nil,
		},
		{
			name: "error : too many arguments",
			argument: []actions.OtherParams{
				{Value: map[string]string{"path": dummyfilepath}},
				{Value: map[string]string{"action": "enable"}},
				{Value: map[string]string{"dummyarg": "dummyargvalue"}},
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "",
			wantedErr:        &actions.Error{Arguments: []string{"path", "action"}},
		},
		{
			name: "error : action not right",
			argument: actions.OtherParams{
				Value: map[string]string{
					"path":   dummyfilepath,
					"action": "comment-xxx",
				},
			},
			isSuccess:        false,
			wantedExitStatus: 1,
			wantedStderr:     "bad action type",
			wantedErr:        nil,
		},
		{
			name: "error: bad action type with uncomment",
			argument: actions.OtherParams{
				Value: map[string]string{
					"path":   dummyfilepath,
					"action": "uncomment",
				},
			},
			isSuccess: false,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
			},
			addLines: []string{
				"#overscan_left=1",
			},
			wantedExitStatus: 1,
			wantedStderr:     "bad action type",
			wantedErr:        nil,
		},
		{
			name: "success with regular params but not enough matches (comment)",
			argument: actions.CommentOrUncommentConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "comment",
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#disable_overscan=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
			},
			addLines: []string{
				"overscan_left=1",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#disable_overscan=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=1",
				"#overscan_left=16",
				"#overscan_right=16",
				"#overscan_top=16",
				"#overscan_bottom=16",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success with regular params (comment)",
			argument: actions.CommentOrUncommentConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "comment",
			},
			isSuccess: true,
			originalLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
			},
			addLines: []string{
				"overscan_left=1",
				"overscan_right=16",
				"overscan_top=16",
				"overscan_bottom=16",
			},
			wantedLines: []string{
				"# uncomment if you get no picture on HDMI for a default safe mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=1",
				"#overscan_right=16",
				"#overscan_top=16",
				"#overscan_bottom=16",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name: "success: file created from asset",
			argument: actions.CommentOrUncommentConfig{
				DirOrFilePath: dummyfilepath,
				Action:        "comment",
			},
			isSuccess:       true,
			createFromAsset: true,
			wantedLines: []string{
				"# For more options and information see",
				"# http://rpf.io/configtxt",
				"# Some settings may impact device functionality. See link above for details",
				"",
				"# uncomment if you get no picture on HDMI for a default \"safe\" mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"#disable_overscan=1",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"#overscan_right=16",
				"#overscan_top=16",
				"#overscan_bottom=16",
				"# uncomment to force a console size. By default it will be display's size minus",
				"# overscan.",
				"#framebuffer_width=1280",
				"#framebuffer_height=720",
				"# uncomment if hdmi display is not detected and composite is being output",
				"#hdmi_force_hotplug=1",
				"# uncomment to force a specific HDMI mode (this will force VGA)",
				"#hdmi_group=1",
				"#hdmi_mode=1",
				"# uncomment to force a HDMI mode rather than DVI. This can make audio work in",
				"# DMT (computer monitor) modes",
				"#hdmi_drive=2",
				"# uncomment to increase signal to HDMI, if you have interference, blanking, or",
				"# no display",
				"#config_hdmi_boost=4",
				"# uncomment for composite PAL",
				"#sdtv_mode=2",
				"#uncomment to overclock the arm. 700 MHz is the default.",
				"#arm_freq=800",
				"# Uncomment some or all of these to enable the optional hardware interfaces",
				"#dtparam=i2c_arm=on",
				"#dtparam=i2s=on",
				"#dtparam=spi=on",
				"# Uncomment this to enable infrared communication.",
				"#dtoverlay=gpio-ir,gpio_pin=17",
				"#dtoverlay=gpio-ir-tx,gpio_pin=18",
				"# Additional overlays and parameters are documented /boot/overlays/README",
				"# Enable audio (loads snd_bcm2835)",
				"dtparam=audio=on",
				"[pi4]",
				"# Enable DRM VC4 V3D driver on top of the dispmanx display stack",
				"dtoverlay=vc4-fkms-v3d",
				"max_framebuffers=2",
				"[all]",
				"#dtoverlay=vc4-fkms-v3d",
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var commentOverscan rpi.Exec
			var err error
			a := actions.New()

			if tc.isSuccess {
				if tc.createFromAsset == false {
					// create and populate file
					if err := actions.OverwriteToFile(actions.WriteToFileArg{
						File:        dummyfilepath,
						Data:        append(tc.originalLines, tc.addLines...),
						Multiline:   true,
						Permissions: 0755,
					}); err != nil {
						log.Fatal(err)
					}
				}

				commentOverscan, err = a.CommentOverscan(tc.argument)
				if err != nil {
					log.Fatal(err)
				}

				// read the new line and delete
				readLines, err := infos.New().ReadFile(dummyfilepath)
				if err != nil {
					log.Fatal(err)
				}

				if e := os.Remove(dummyfilepath); e != nil {
					fmt.Println(e)
				}

				assert.Equal(t, tc.wantedLines, readLines)
			} else {
				commentOverscan, err = a.CommentOverscan(tc.argument)
			}

			assert.Equal(t, tc.wantedExitStatus, commentOverscan.ExitStatus)
			assert.Equal(t, tc.wantedStderr, commentOverscan.Stderr)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestCreateAssetFile(t *testing.T) {
	cases := []struct {
		name             string
		argument         actions.CreateAssetFileArg
		isSuccess        bool
		wantedLines      []string
		wantedExitStatus int
		wantedStderr     string
	}{
		{
			name: "error: wrong key",
			argument: actions.CreateAssetFileArg{
				AssetFile:     "../assets/config.xxx",
				TargetFile:    dummyfilepath,
				HasUniqueLine: true,
			},
			wantedExitStatus: 1,
			wantedStderr:     "couldn't find asset file",
			wantedLines:      nil,
		},
		{
			name: "success: config.txt",
			argument: actions.CreateAssetFileArg{
				AssetFile:     "../assets/config.txt",
				TargetFile:    dummyfilepath,
				HasUniqueLine: true,
			},
			isSuccess:        true,
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedLines: []string{
				"# For more options and information see",
				"# http://rpf.io/configtxt",
				"# Some settings may impact device functionality. See link above for details",
				"",
				"# uncomment if you get no picture on HDMI for a default \"safe\" mode",
				"#hdmi_safe=1",
				"# uncomment this if your display has a black border of unused pixels visible",
				"# and your display can output without overscan",
				"#disable_overscan=1",
				"# uncomment the following to adjust overscan. Use positive numbers if console",
				"# goes off screen, and negative if there is too much border",
				"#overscan_left=16",
				"#overscan_right=16",
				"#overscan_top=16",
				"#overscan_bottom=16",
				"# uncomment to force a console size. By default it will be display's size minus",
				"# overscan.",
				"#framebuffer_width=1280",
				"#framebuffer_height=720",
				"# uncomment if hdmi display is not detected and composite is being output",
				"#hdmi_force_hotplug=1",
				"# uncomment to force a specific HDMI mode (this will force VGA)",
				"#hdmi_group=1",
				"#hdmi_mode=1",
				"# uncomment to force a HDMI mode rather than DVI. This can make audio work in",
				"# DMT (computer monitor) modes",
				"#hdmi_drive=2",
				"# uncomment to increase signal to HDMI, if you have interference, blanking, or",
				"# no display",
				"#config_hdmi_boost=4",
				"# uncomment for composite PAL",
				"#sdtv_mode=2",
				"#uncomment to overclock the arm. 700 MHz is the default.",
				"#arm_freq=800",
				"# Uncomment some or all of these to enable the optional hardware interfaces",
				"#dtparam=i2c_arm=on",
				"#dtparam=i2s=on",
				"#dtparam=spi=on",
				"# Uncomment this to enable infrared communication.",
				"#dtoverlay=gpio-ir,gpio_pin=17",
				"#dtoverlay=gpio-ir-tx,gpio_pin=18",
				"# Additional overlays and parameters are documented /boot/overlays/README",
				"# Enable audio (loads snd_bcm2835)",
				"dtparam=audio=on",
				"[pi4]",
				"# Enable DRM VC4 V3D driver on top of the dispmanx display stack",
				"dtoverlay=vc4-fkms-v3d",
				"max_framebuffers=2",
				"[all]",
				"#dtoverlay=vc4-fkms-v3d",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			exitStatus, stdErr := actions.CreateAssetFile(tc.argument)

			if tc.isSuccess {
				// read the new line and delete
				readLines, err := infos.New().ReadFile(dummyfilepath)
				if err != nil {
					log.Fatal(err)
				}

				if e := os.Remove(dummyfilepath); e != nil {
					fmt.Println(e)
				}

				assert.Equal(t, tc.wantedLines, readLines)
			}

			assert.Equal(t, tc.wantedExitStatus, exitStatus)
			assert.Equal(t, tc.wantedStderr, stdErr)
		})
	}
}
