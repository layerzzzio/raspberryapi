package actions_test

import (
	"fmt"
	"os"
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
