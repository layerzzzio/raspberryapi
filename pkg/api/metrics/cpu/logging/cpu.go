package cpu

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/cpu"
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

// View is the logging function attached to the View cpu services and responsible for logging it out.
func (ls *LogService) View(ctx echo.Context, id int) (resp rpi.CPU, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: viewing cpu #%v", id), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(id)
}
