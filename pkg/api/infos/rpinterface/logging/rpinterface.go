package rpinterface

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/rpinterface"
)

// New creates a new rpinterface logging service instance.
func New(svc rpinterface.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a rpinterface logging service.
type LogService struct {
	rpinterface.Service
	logger rpi.Logger
}

const name = "rpinterface"

// List is the logging function attached to the List rpinterface services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.RpInterface, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: listing rpinterface configuration",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
