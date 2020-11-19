package test_utl

import (
	"github.com/raspibuddy/rpi"
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
