package filestructure

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Service represents all filestructure application services.
type Service interface {
	ViewLF(path string, pathSize uint64, fileLimit float32) (rpi.FileStructure, error)
}

// FileStructure represents a FileStructure application service.
type FileStructure struct {
	fssys FSSYS
	mt    Metrics
}

// LFSYS represents a filestructure repository service.
type FSSYS interface {
	ViewLF(*rpi.File, map[int64]string) (rpi.FileStructure, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	WalkFolder(
		string,
		metrics.ReadDir,
		uint64,
		float32,
		metrics.ShouldIgnoreFolder,
		chan int,
	) (*rpi.File, map[int64]string)
}

// New creates a FileStructure application service instance.
func New(lfsys FSSYS, mt Metrics) *FileStructure {
	return &FileStructure{fssys: lfsys, mt: mt}
}
