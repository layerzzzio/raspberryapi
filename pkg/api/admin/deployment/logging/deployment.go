package deployment

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/admin/deployment"
)

// New creates a new deployment logging service instance.
func New(svc deployment.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a deployment logging service.
type LogService struct {
	deployment.Service
	logger rpi.Logger
}

const name = "deployment"

// List is the logging function attached to the List service and responsible for logging it out.
func (ls *LogService) ExecuteDP(ctx echo.Context, version string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: deploying api version "+version,
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteDP(version)
}
