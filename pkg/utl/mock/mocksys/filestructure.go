package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// FileStructure mock
type FileStructure struct {
	ViewFLFn func(*rpi.File, map[int64]string) (rpi.FileStructure, error)
}

// View mock
func (fs FileStructure) ViewLF(fileStructure *rpi.File, flattenFiles map[int64]string) (rpi.FileStructure, error) {
	return fs.ViewFLFn(fileStructure, flattenFiles)
}
