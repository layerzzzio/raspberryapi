package test_utl

import (
	"fmt"
	"os"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// NewTestFolder is providing easy interface to create folders for automated tests
// Never use in production code!
func NewTestFolder(name string, files ...*rpi.File) *rpi.File {
	folder := &rpi.File{
		Name:   name,
		Parent: nil,
		Size:   0,
		IsDir:  true,
		Files:  []*rpi.File{},
	}

	if files == nil {
		return folder
	}
	for _, file := range files {
		file.Parent = folder
	}
	folder.Files = files
	metrics.UpdateSize(folder)
	return folder
}

// NewTestFile provides easy interface to create files for automated tests
// Never use in production code!
func NewTestFile(name string, size int64) *rpi.File {
	return &rpi.File{
		Name:   name,
		Parent: nil,
		Size:   size,
		IsDir:  false,
		Files:  []*rpi.File{},
	}
}

// FindTestFile helps testing by returning first occurrence of file with given name.
// Never use in production code!
func FindTestFile(folder *rpi.File, name string) *rpi.File {
	if folder.Name == name {
		return folder
	}
	for _, file := range folder.Files {
		result := FindTestFile(file, name)
		if result != nil {
			return result
		}
	}
	return nil
}

func CreateFile(path string) bool {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if IsError(err) {
			return false
		}
		defer file.Close()
	}
	return true
}

func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

type ArgFuncA struct {
	Arg0 string
	Arg1 string
}

func FuncA(arg interface{}) (rpi.Exec, error) {
	var arg0 string
	var arg1 string

	switch v := arg.(type) {
	case ArgFuncA:
		arg0 = v.Arg0
		arg1 = v.Arg1
	case actions.OtherParams:
		arg0 = arg.(actions.OtherParams).Value["arg0"]
		arg1 = arg.(actions.OtherParams).Value["arg1"]
	default:
		return rpi.Exec{}, &actions.Error{Arguments: []string{"arg0", "arg1"}}
	}

	stdOut := fmt.Sprintf("%v-%v", arg0, arg1)

	time.Sleep(5 * time.Second)

	res := rpi.Exec{
		Name:       "FuncA",
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

func FuncB(arg interface{}) (rpi.Exec, error) {
	var arg2 string

	switch v := arg.(type) {
	case ArgFuncB:
		arg2 = v.Arg2
	case actions.OtherParams:
		arg2 = arg.(actions.OtherParams).Value["arg2"]
	default:
		return rpi.Exec{}, &actions.Error{Arguments: []string{"arg2"}}
	}

	stdOut := fmt.Sprint(arg2)

	time.Sleep(2 * time.Second)

	res := rpi.Exec{
		Name:       "FuncB",
		StartTime:  1,
		EndTime:    2,
		ExitStatus: 0,
		Stdin:      "",
		Stdout:     stdOut,
		Stderr:     "",
	}

	return res, nil
}

type ArgFuncC struct {
	Arg3 string
}

func FuncC(arg interface{}) (rpi.Exec, error) {
	var arg3 string

	switch v := arg.(type) {
	case ArgFuncC:
		arg3 = v.Arg3
	case actions.OtherParams:
		arg3 = arg.(actions.OtherParams).Value["arg3"]
	default:
		return rpi.Exec{}, &actions.Error{Arguments: []string{"arg3"}}
	}

	stdOut := fmt.Sprint(arg3)

	time.Sleep(2 * time.Second)

	res := rpi.Exec{
		Name:       "FuncC",
		StartTime:  1,
		EndTime:    2,
		ExitStatus: 0,
		Stdin:      "",
		Stdout:     stdOut,
		Stderr:     "",
	}

	return res, nil
}
