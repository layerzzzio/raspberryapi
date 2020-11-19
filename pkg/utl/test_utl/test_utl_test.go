package test_utl_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
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
