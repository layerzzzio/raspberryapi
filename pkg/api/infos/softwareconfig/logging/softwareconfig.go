package softwareconfig

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig"
)

// New creates a new softwareconfig logging service instance.
func New(svc softwareconfig.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a softwareconfig logging service.
type LogService struct {
	softwareconfig.Service
	logger rpi.Logger
}

const name = "softwareconfig"

// List is the logging function attached to the List softwareconfig services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.SoftwareConfig, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: listing softwareconfig configuration",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
