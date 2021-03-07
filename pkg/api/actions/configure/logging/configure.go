package configure

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure"
)

// New creates a new Configure logging service instance.
func New(svc configure.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a Configure logging service.
type LogService struct {
	configure.Service
	logger rpi.Logger
}

const name = "configure"

// ExecuteCH is the logging function attached to the execute change hostname service and responsible for logging it out.
func (ls *LogService) ExecuteCH(ctx echo.Context, hostname string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute change hostname #%v", hostname), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteCH(hostname)
}

// ExecuteCP is the logging function attached to the execute change password service and responsible for logging it out.
func (ls *LogService) ExecuteCP(ctx echo.Context, password string, username string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			fmt.Sprintf("request: execute change hostname #%v for user %v", password, username),
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteCP(password, username)
}

// ExecuteWNB is the logging function attached to the execute wait for network at bool service and responsible for logging it out.
func (ls *LogService) ExecuteWNB(ctx echo.Context, action string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			fmt.Sprintf("request: execute %v wait for network at boot", action),
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteWNB(action)
}

// ExecuteOV is the logging function attached to the execute overscan service and responsible for logging it out.
func (ls *LogService) ExecuteOV(ctx echo.Context, action string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			fmt.Sprintf("request: execute %v overscan", action),
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteOV(action)
}

// ExecuteBL is the logging function attached to the execute blanking service and responsible for logging it out.
func (ls *LogService) ExecuteBL(ctx echo.Context, action string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			fmt.Sprintf("request: execute %v blanking", action),
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteBL(action)
}

// ExecuteAUS is the logging function attached to the execute add user service and responsible for logging it out.
func (ls *LogService) ExecuteAUS(
	ctx echo.Context,
	username string,
	password string,
) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: execute add user",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteAUS(username, password)
}

// ExecuteDUS is the logging function attached to the execute delete user service and responsible for logging it out.
func (ls *LogService) ExecuteDUS(
	ctx echo.Context,
	username string,
	password string,
) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			"request: execute delete user",
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteDUS(username)
}

// ExecuteCA is the logging function attached to the execute camera service and responsible for logging it out.
func (ls *LogService) ExecuteCA(ctx echo.Context, action string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name,
			fmt.Sprintf("request: execute %v camera", action),
			err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteCA(action)
}
