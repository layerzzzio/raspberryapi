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
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
	"github.com/stretchr/testify/assert"
)

var (
	dummypath = "./dummyfile"
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
			argument:         actions.DF{Path: ""},
			wantedExitStatus: 1,
			wantedStderr:     "remove : no such file or directory",
			wantedErr:        nil,
		},
		{
			name: "error : too many arguments",
			argument: []actions.OtherParams{
				{Value: map[string]string{"path": dummypath}},
				{Value: map[string]string{"dummy": dummypath}},
			},
			wantedExitStatus: 1,
			wantedStderr:     "",
			wantedErr:        &actions.Error{Arguments: []string{"path"}},
		},
		{
			name: "success",
			argument: actions.OtherParams{
				Value: map[string]string{
					"path": dummypath,
				},
			},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
		{
			name:             "success",
			argument:         actions.DF{Path: dummypath},
			wantedExitStatus: 0,
			wantedStderr:     "",
			wantedErr:        nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := actions.New()
			test_utl.CreateFile(dummypath)
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

func TestBackupFile(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		perm       int
		wantedData error
	}{
		{
			name:       "success file does not exist",
			path:       "./TestBackupFile",
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
				f, err := os.Create(tc.path)
				if err != nil {
					fmt.Println(err)
					log.Fatal()
				}

				_, err = f.WriteString("hey man")
				if err != nil {
					fmt.Println(err)
					log.Fatal()
				}

				err = f.Close()
				if err != nil {
					fmt.Println(err)
					log.Fatal()
				}
			}
			backupFile := actions.BackupFile(tc.path, tc.perm)
			assert.Equal(t, tc.wantedData, backupFile)
		})
	}
}

func TestOverwriteToFile(t *testing.T) {
	cases := []struct {
		name       string
		args       actions.OverwriteToFileArg
		wantedData error
	}{
		{
			name: "success with multiline",
			args: actions.OverwriteToFileArg{
				File:      "./test_write_to_file",
				Data:      []string{"text_1", "text_2", "text_3"},
				Multiline: true,
			},
			wantedData: nil,
		},
		{
			name: "success not multiline",
			args: actions.OverwriteToFileArg{
				File:      "./test_write_to_file",
				Data:      []string{"text_1", "text_2", "text_3"},
				Multiline: false,
			},
			wantedData: nil,
		},
		{
			name: "success permissions not nill",
			args: actions.OverwriteToFileArg{
				File:        "./test_write_to_file",
				Data:        []string{"text_1", "text_2", "text_3"},
				Multiline:   false,
				Permissions: 0755,
			},
			wantedData: nil,
		},
		{
			name: "failure creating file",
			args: actions.OverwriteToFileArg{
				File:      "",
				Data:      []string{"text_1", "text_2", "text_3"},
				Multiline: false,
			},
			wantedData: fmt.Errorf("creating file failed"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			overwriteToFile := actions.OverwriteToFile(tc.args)

			if tc.name == "success with multiline" ||
				tc.name == "success not multiline" ||
				tc.name == "success permissions not nill" {
				file, err := os.Open(tc.args.File)
				if err != nil {
					log.Fatal(err)
				}

				var readLines = []string{}
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					readLines = append(readLines, scanner.Text())
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}

				file.Close()

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
