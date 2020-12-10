package test_utl_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
	"github.com/stretchr/testify/assert"
)

func TestBuildFile(t *testing.T) {
	a := &rpi.File{
		Name:   "a",
		Parent: nil,
		Size:   100,
		IsDir:  false,
		Files:  []*rpi.File{},
	}
	build := test_utl.NewTestFile("a", 100)
	assert.Equal(t, a, build)
}

func TestBuildFolder(t *testing.T) {
	a := &rpi.File{
		Name:   "a",
		Parent: nil,
		Size:   0,
		IsDir:  true,
		Files:  []*rpi.File{},
	}
	build := test_utl.NewTestFolder("a")
	assert.Equal(t, a, build)
}

func TestBuildFolderWithFile(t *testing.T) {
	e := &rpi.File{
		Name:   "e",
		Parent: nil,
		Size:   100,
		IsDir:  false,
		Files:  []*rpi.File{},
	}
	d := &rpi.File{
		Name:   "d",
		Parent: nil,
		Size:   100,
		IsDir:  true,
		Files:  []*rpi.File{e},
	}
	e.Parent = d
	build := test_utl.NewTestFolder("d", test_utl.NewTestFile("e", 100))
	assert.Equal(t, d, build)
}

func TestBuildComplexFolder(t *testing.T) {
	e := &rpi.File{
		Name:   "e",
		Parent: nil,
		Size:   100,
		IsDir:  false,
		Files:  []*rpi.File{},
	}
	d := &rpi.File{
		Name:   "d",
		Parent: nil,
		Size:   100,
		IsDir:  true,
		Files:  []*rpi.File{e},
	}
	e.Parent = d
	b := &rpi.File{
		Name:   "b",
		Parent: nil,
		Size:   50,
		IsDir:  false,
		Files:  []*rpi.File{},
	}
	c := &rpi.File{
		Name:   "c",
		Parent: nil,
		Size:   100,
		IsDir:  false,
		Files:  []*rpi.File{},
	}
	a := &rpi.File{
		Name:   "a",
		Parent: nil,
		Size:   250,
		IsDir:  true,
		Files:  []*rpi.File{b, c, d},
	}
	b.Parent = a
	c.Parent = a
	d.Parent = a
	build := test_utl.NewTestFolder("a", test_utl.NewTestFile("b", 50), test_utl.NewTestFile("c", 100), test_utl.NewTestFolder("d", test_utl.NewTestFile("e", 100)))
	assert.Equal(t, a, build)
}

func TestFindTestFile(t *testing.T) {
	folder := test_utl.NewTestFolder("a",
		test_utl.NewTestFolder("b",
			test_utl.NewTestFile("c", 10),
			test_utl.NewTestFile("d", 100),
		),
	)
	expected := folder.Files[0].Files[1]
	foundFile := test_utl.FindTestFile(folder, "d")
	assert.Equal(t, expected, foundFile)
}

func TestIsError(t *testing.T) {
	cases := []struct {
		name       string
		params     error
		wantedData bool
	}{
		{
			name:       "error: wrong argument",
			params:     errors.Errorf("dummy"),
			wantedData: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := test_utl.IsError(tc.params)
			assert.Equal(t, tc.wantedData, err)
		})
	}
}

func TestCreateFile(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		wantedData bool
	}{
		{
			name:       "error: wrong argument",
			path:       "./dummy",
			wantedData: true,
		},
		{
			name:       "error: wrong argument",
			path:       "./dummypath/dummy",
			wantedData: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := actions.New()
			res := test_utl.CreateFile(tc.path)
			fmt.Println(res)
			if res {
				del, err := a.DeleteFile(actions.DF{Path: tc.path})
				if err != nil {
					fmt.Println(del)
					fmt.Println(err)
				}
			}
			assert.Equal(t, tc.wantedData, res)
		})
	}
}

func TestFuncA(t *testing.T) {
	cases := []struct {
		name       string
		params     interface{}
		wantedData rpi.Exec
		wantedErr  error
	}{
		{
			name:       "error: wrong argument arg0",
			params:     "dummy",
			wantedData: rpi.Exec{},
			wantedErr:  &actions.Error{Arguments: []string{"arg0", "arg1"}},
		},
		{
			name:   "success with params",
			params: test_utl.ArgFuncA{Arg0: "string0", Arg1: "string1"},
			wantedData: rpi.Exec{
				Name:       "FuncA",
				StartTime:  1,
				EndTime:    2,
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "string0-string1",
				Stderr:     "",
			},
			wantedErr: nil,
		},
		{
			name: "success with otherparams",
			params: actions.OtherParams{
				Value: map[string]string{
					"arg0": "string0",
					"arg1": "string1",
				},
			},
			wantedData: rpi.Exec{
				Name:       "FuncA",
				StartTime:  1,
				EndTime:    2,
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "string0-string1",
				Stderr:     "",
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := test_utl.FuncA(tc.params)
			assert.Equal(t, tc.wantedData, res)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestFuncB(t *testing.T) {
	cases := []struct {
		name       string
		params     interface{}
		wantedData rpi.Exec
		wantedErr  error
	}{
		{
			name:       "error: wrong argument",
			params:     "dummy",
			wantedData: rpi.Exec{},
			wantedErr:  &actions.Error{Arguments: []string{"arg2"}},
		},
		{
			name:   "success with params",
			params: test_utl.ArgFuncB{Arg2: "string2"},
			wantedData: rpi.Exec{
				Name:       "FuncB",
				StartTime:  1,
				EndTime:    2,
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "string2",
				Stderr:     "",
			},
			wantedErr: nil,
		},
		{
			name: "success with otherparams",
			params: actions.OtherParams{
				Value: map[string]string{
					"arg2": "string2",
				},
			},
			wantedData: rpi.Exec{
				Name:       "FuncB",
				StartTime:  1,
				EndTime:    2,
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "string2",
				Stderr:     "",
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := test_utl.FuncB(tc.params)
			assert.Equal(t, tc.wantedData, res)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestFuncC(t *testing.T) {
	cases := []struct {
		name       string
		params     interface{}
		wantedData rpi.Exec
		wantedErr  error
	}{
		{
			name:       "error: wrong argument",
			params:     "dummy",
			wantedData: rpi.Exec{},
			wantedErr:  &actions.Error{Arguments: []string{"arg3"}},
		},
		{
			name:   "success with params",
			params: test_utl.ArgFuncC{Arg3: "string3"},
			wantedData: rpi.Exec{
				Name:       "FuncC",
				StartTime:  1,
				EndTime:    2,
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "string3",
				Stderr:     "",
			},
			wantedErr: nil,
		},
		{
			name: "success with otherparams",
			params: actions.OtherParams{
				Value: map[string]string{
					"arg3": "string3",
				},
			},
			wantedData: rpi.Exec{
				Name:       "FuncC",
				StartTime:  1,
				EndTime:    2,
				ExitStatus: 0,
				Stdin:      "",
				Stdout:     "string3",
				Stderr:     "",
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := test_utl.FuncC(tc.params)
			assert.Equal(t, tc.wantedData, res)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
