package version

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/version"
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

// ListAll is the logging function attached to the ListAll service and responsible for logging it out.
func (ls *LogService) ListAll(ctx echo.Context) (resp rpi.Version, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: listing all versions on the system",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ListAll()
}

// ListAllApis is the logging function attached to the ListAllApis service and responsible for logging it out.
func (ls *LogService) ListAllApis(ctx echo.Context) (resp rpi.Version, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: listing all apis versions on the system",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ListAllApis()
}
