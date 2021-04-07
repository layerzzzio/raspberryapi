package appstatus

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/appstatus"
)

// New creates a new appstatus logging service instance.
func New(svc appstatus.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a appstatus logging service.
type LogService struct {
	appstatus.Service
	logger rpi.Logger
}

const name = "appstatus"

// List is the logging function attached to the List appstatus services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.AppStatus, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: listing appstatus configuration",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
