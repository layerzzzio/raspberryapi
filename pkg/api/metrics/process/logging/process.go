package process

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/process"
)

// New creates a new process logging service instance.
func New(svc process.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a process logging service.
type LogService struct {
	process.Service
	logger rpi.Logger
}

const name = "process"

// List is the logging function attached to the List process services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.Process, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing process", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}

// View is the logging function attached to the View process services and responsible for logging it out.
func (ls *LogService) View(ctx echo.Context, id int32) (resp rpi.Process, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: viewing process #%v", id), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(id)
}
