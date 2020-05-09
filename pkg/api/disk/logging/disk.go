package disk

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/disk"
)

// New creates a new disk logging service instance.
func New(svc disk.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a disk logging service.
type LogService struct {
	disk.Service
	logger rpi.Logger
}

const name = "disk"

// List is the logging function attached to the List disk services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.Disk, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing disk", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}

// View is the logging function attached to the View disk services and responsible for logging it out.
func (ls *LogService) View(ctx echo.Context, dev string) (resp rpi.Disk, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: viewing disk #%v", dev), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(dev)
}
