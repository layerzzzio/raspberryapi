package largestfiles

import (
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfiles"
)

// New creates a new largest files logging service instance.
func New(svc largestfiles.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a largest files logging service.
type LogService struct {
	largestfiles.Service
	logger rpi.Logger
}

const name = "largestfiles"

// List is the logging function attached to the List largest files services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.LargestFiles, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing largestfiles", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
