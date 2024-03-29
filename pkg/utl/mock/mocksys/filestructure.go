package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// FileStructure mock
type FileStructure struct {
	ViewLFFn func(*rpi.File, map[string]int64) (rpi.FileStructure, error)
}

// View mock
func (fs FileStructure) ViewLF(fileStructure *rpi.File, flattenFiles map[string]int64) (rpi.FileStructure, error) {
	return fs.ViewLFFn(fileStructure, flattenFiles)
}
