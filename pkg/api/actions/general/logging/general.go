package general

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/general"
)

// New creates a new General logging service instance.
func New(svc general.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a General logging service.
type LogService struct {
	general.Service
	logger rpi.Logger
}

const name = "reboot_or_shutdown"

// ExecuteRBS is the logging function attached to the general services and responsible for logging it out.
func (ls *LogService) ExecuteRBS(ctx echo.Context, actionType string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute reboot or shutdown #%v", actionType), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteRBS(actionType)
}

// ExecuteSASO is the logging function attached to the general services and responsible for logging it out.
func (ls *LogService) ExecuteSASO(ctx echo.Context, actionType string, service string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: %v %v", actionType, service), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteSASO(actionType, service)
}
