package software

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
)

// New creates a new software logging service instance.
func New(svc software.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a software logging service.
type LogService struct {
	software.Service
	logger rpi.Logger
}

const name = "software"

// List is the logging function attached to the List software services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.Software, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: listing software configuration",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
