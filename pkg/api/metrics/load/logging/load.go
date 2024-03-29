package load

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/load"
)

// New creates a new load logging service instance.
func New(svc load.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a load logging service.
type LogService struct {
	load.Service
	logger rpi.Logger
}

const name = "load"

// List is the logging function attached to the List load services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.Load, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing loads", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
