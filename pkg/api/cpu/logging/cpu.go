package cpu

import (
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/cpu"
)

// New creates a new CPU logging service instance.
func New(svc cpu.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a CPU logging service.
type LogService struct {
	cpu.Service
	logger rpi.Logger
}

const name = "cpu"

// List is the logging function attached to the List CPU services and responsible for logging it out.
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
