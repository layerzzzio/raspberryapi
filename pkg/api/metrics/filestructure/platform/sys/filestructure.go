package sys

import (
	"sort"

	"github.com/raspibuddy/rpi"
)

// FileStructure represents an empty FileStructure entity on the current system.
type FileStructure struct{}

// View returns a filestructure given a file path
func (fs FileStructure) ViewLF(fileStructure *rpi.File, flattenFiles map[string]int64) (rpi.FileStructure, error) {
	// fmt.Printf("====> filestructure: %v", fileStructure)
	// fmt.Printf("====> flattenFiles: %v", flattenFiles)
	// get 100 biggest files in the given path
	if len(fileStructure.Files) == 0 {
		return rpi.FileStructure{
			DirectoryPath: fileStructure.Name,
		}, nil
	}

	values := []int{}
	for _, v := range flattenFiles {
		if v != 0 {
			values = append(values, int(v))
		}
	}
	// sort keys array in increasing order
	// 1, 50, 1000, 200000 etc.
	sort.Ints(values)

	var truncatedValues = []int{}
	var largestFiles []*rpi.File

	if len(values) > 0 {
		max := len(values)
		min := len(values) - 100

		if len(values) == 1 {
			truncatedValues = values
		} else if len(values) > 100 {
			truncatedValues = values[min:max]
		} else if len(values) <= 100 {
			truncatedValues = values[0:max]
		}

		for _, v := range truncatedValues {
			key, _ := mapkey(flattenFiles, int64(v))
			delete(flattenFiles, key)
			largestFiles = append(
				largestFiles,
				&rpi.File{
					Path: key,
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

func mapkey(m map[string]int64, value int64) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
