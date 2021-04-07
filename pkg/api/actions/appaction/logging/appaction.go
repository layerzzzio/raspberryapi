package appaction

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/appaction"
)

// New creates a new AppAction logging service instance.
func New(svc appaction.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a AppAction logging service.
type LogService struct {
	appaction.Service
	logger rpi.Logger
}

const name = "appaction"

// ExecuteWOVA is the logging function attached to the execute app action vpn with ovpn service and responsible for logging it out.
func (ls *LogService) ExecuteWOVA(
	ctx echo.Context,
	action string,
	vpnName string,
	relativeConfigPath string,
	country string,
	username string,
	password string,
) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute %v for %v", action, vpnName), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteWOVA(action, vpnName, relativeConfigPath, country, username, password)
}
