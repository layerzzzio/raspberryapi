package filestructure

import (
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/filestructure"
)

// New creates a new FileStructure logging service instance.
func New(svc filestructure.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a FileStructure logging service.
type LogService struct {
	filestructure.Service
	logger rpi.Logger
}

const name = "filestructure"

// View is the logging function attached to the a file structure services and responsible for logging it out.
func (ls *LogService) ViewLF(ctx echo.Context, path string, pathSize uint64, fileLimit float32) (resp rpi.FileStructure, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: viewing filestructure", err,
			map[string]interface{}{
				"resp":      resp,
				"took":      time.Since(begin),
				"path":      path,
				"pathSize":  pathSize,
				"fileLimit": fileLimit,
			},
		)
	}(time.Now())
	return ls.Service.ViewLF(path, pathSize, fileLimit)
}
