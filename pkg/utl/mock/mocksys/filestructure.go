package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// FileStructure mock
type FileStructure struct {
	ViewFn func(*rpi.File, map[int64]string) (rpi.FileStructure, error)
}

// View mock
func (fs FileStructure) View(fileStructure *rpi.File, flattenFiles map[int64]string) (rpi.FileStructure, error) {
	return fs.ViewFn(fileStructure, flattenFiles)
}
