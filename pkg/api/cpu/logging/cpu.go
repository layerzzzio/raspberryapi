package cpu

import (
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/cpu"
)

// New creates a new cpu logging service.
func New(svc cpu.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a cpu logging service.
type LogService struct {
	cpu.Service
	logger rpi.Logger
}

const name = "cpu"

// List is a logging method specific to the cpu.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.CPU, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing cpus", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
