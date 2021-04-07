package appconfig

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/appconfig"
)

// New creates a new appconfig logging service instance.
func New(svc appconfig.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a appconfig logging service.
type LogService struct {
	appconfig.Service
	logger rpi.Logger
}

const name = "appconfig"

// List is the logging function attached to the List appconfig services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.AppConfig, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: listing appconfig configuration",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
