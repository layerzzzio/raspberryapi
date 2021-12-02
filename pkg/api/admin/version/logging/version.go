package version

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/admin/version"
)

// New creates a new version logging service instance.
func New(svc version.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a version logging service.
type LogService struct {
	version.Service
	logger rpi.Logger
}

const name = "version"

// List is the logging function attached to the List service and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.Version, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: listing api version",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
