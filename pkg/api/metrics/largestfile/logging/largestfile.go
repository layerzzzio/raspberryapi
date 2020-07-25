package largestfile

import (
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfile"
)

// New creates a new largest files logging service instance.
func New(svc largestfile.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a largest files logging service.
type LogService struct {
	largestfile.Service
	logger rpi.Logger
}

const name = "largestfile"

// List is the logging function attached to the List largest files services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.LargestFile, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing largestfile", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
