package configfile

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/display"
)

// New creates a new configfile logging service instance.
func New(svc display.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a configfile logging service.
type LogService struct {
	display.Service
	logger rpi.Logger
}

const name = "configfile"

// List is the logging function attached to the List configfile services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.Display, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: list config files metadata", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
