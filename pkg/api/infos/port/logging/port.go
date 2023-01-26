package port

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/port"
)

// New creates a new port logging service instance.
func New(svc port.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a port logging service.
type LogService struct {
	port.Service
	logger rpi.Logger
}

const name = "port"

// View is the logging function attached to the View service and responsible for logging it out.
func (ls *LogService) View(ctx echo.Context, port int32) (resp rpi.Port, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: view specific port status",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(port)
}
