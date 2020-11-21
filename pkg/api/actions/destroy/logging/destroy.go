package destroy

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
)

// New creates a new DeleteFile logging service instance.
func New(svc destroy.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a DeleteFile logging service.
type LogService struct {
	destroy.Service
	logger rpi.Logger
}

const name = "destroy"

// ExecuteDF is the logging function attached to the Execute delete file destroy services and responsible for logging it out.
func (ls *LogService) ExecuteDF(ctx echo.Context, path string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute delete file #%v", path), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteDF(path)
}

// ExecuteDU is the logging function attached to the Execute delete file destroy services and responsible for logging it out.
func (ls *LogService) ExecuteDU(ctx echo.Context, terminalname string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute kill terminal %v", terminalname), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteDU(terminalname)
}

// ExecuteKP is the logging function attached to the Execute kill process destroy services and responsible for logging it out.
func (ls *LogService) ExecuteKP(ctx echo.Context, pid int) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: kill process %v", pid), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteKP(pid)
}
