package filestructure

import (
	"io/ioutil"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// View returns a FileStructure model.
func (fs *FileStructure) ViewLF(path string, pathSize uint64, fileLimit int8) (rpi.FileStructure, error) {
	progress := make(chan int)

	fileStructure, flattenFiles := fs.mt.WalkFolder(
		path,
		ioutil.ReadDir,
		pathSize,
		fileLimit,
		metrics.IgnoreBasedOnIgnoreFile([]string{
			"opt",
			"usr",
			"proc",
			"sys",
			"var",
			"dev",
			"etc",
			"sbin",
			"bin",
			"lib",
			".",
		}),
		progress)

	return fs.fssys.ViewLF(fileStructure, flattenFiles)
}
