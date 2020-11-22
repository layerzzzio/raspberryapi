package actions_test

import (
	"fmt"
	"os"
	"os/exec"
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
		name        string
		processname string
		processtype string
		wantedData  rpi.Exec
	}{
		{
			name:        "error killing process by its name",
			processname: "impossible_process_name",
			processtype: "dummy",
			wantedData: rpi.Exec{
				Name:       "kill_process_by_name",
				StartTime:  uint64(time.Now().Unix()),
				EndTime:    uint64(time.Now().Unix()),
				ExitStatus: 1,
				Stdin:      "",
				Stdout:     "",
				Stderr:     "exit status 1",
			},
		},
		{
			name:        "error killing process by its name (terminal)",
			processname: "impossible_process_name",
			processtype: "terminal",
			wantedData: rpi.Exec{
				Name:       "kill_process_by_name",
				StartTime:  uint64(time.Now().Unix()),
				EndTime:    uint64(time.Now().Unix()),
				ExitStatus: 1,
				Stdin:      "",
				Stdout:     "",
				Stderr:     "exit status 2",
			},
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

			largestfiles := a.KillProcessByName(tc.processname, tc.processtype)

			err = cmd.Wait()
			if err == nil {
				t.Errorf("Test process succeeded, but expected to fail")
			}

			assert.Equal(t, tc.wantedData, largestfiles)
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
