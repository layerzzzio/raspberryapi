package humanuser

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/humanuser"
)

// New creates a new humanuser logging service instance.
func New(svc humanuser.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a humanuser logging service.
type LogService struct {
	humanuser.Service
	logger rpi.Logger
}

const name = "humanuser"

// List is the logging function attached to the List humanuser services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.HumanUser, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing human users", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
