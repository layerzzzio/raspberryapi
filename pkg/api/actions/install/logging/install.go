package install

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/install"
)

// New creates a new Install logging service instance.
func New(svc install.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a Install logging service.
type LogService struct {
	install.Service
	logger rpi.Logger
}

const name = "install"

// ExecuteAG is the logging function attached to the execute install apt-get service and responsible for logging it out.
func (ls *LogService) ExecuteAG(ctx echo.Context, action string, pkg string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute %v #%v", action, pkg), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteAG(action, pkg)
}

// ExecuteNV is the logging function attached to the execute install nordvpn service and responsible for logging it out.
func (ls *LogService) ExecuteNV(ctx echo.Context, action string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute %v", action), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteNV(action)
}
