package sys

import (
	"sort"

	"github.com/raspibuddy/rpi"
)

// FileStructure represents an empty FileStructure entity on the current system.
type FileStructure struct{}

// View returns a filestructure given a file path
func (fs FileStructure) ViewLF(fileStructure *rpi.File, flattenFiles map[int64]string) (rpi.FileStructure, error) {
	// get 100 biggest files in the given path
	if len(fileStructure.Files) == 0 {
		return rpi.FileStructure{
			DirectoryPath: fileStructure.Name,
		}, nil
	}

	keys := []int{}
	for key := range flattenFiles {
		if key != 0 {
			keys = append(keys, int(key))
		}
	}
	// sort keys array in increasing order
	// 1, 50, 1000, 200000 etc.
	sort.Ints(keys)

	var truncatedKeys = []int{}
	var largestFiles []*rpi.File

	if len(keys) > 0 {
		max := len(keys)
		min := len(keys) - 100

		if len(keys) == 1 {
			truncatedKeys = keys
		} else if len(keys) > 100 {
			truncatedKeys = keys[min:max]
		} else if len(keys) <= 100 {
			truncatedKeys = keys[0:max]
		}

		for _, v := range truncatedKeys {
			largestFiles = append(
				largestFiles,
				&rpi.File{
					Path: flattenFiles[int64(v)],
					Size: int64(v),
				})
		}
	} else {
		largestFiles = nil
	}

	return rpi.FileStructure{
		DirectoryPath: fileStructure.Name,
		LargestFiles:  largestFiles,
	}, nil
}
