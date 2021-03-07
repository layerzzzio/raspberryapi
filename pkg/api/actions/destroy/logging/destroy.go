package destroy

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
)

// New creates a new Destroy logging service instance.
func New(svc destroy.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a Destroy logging service.
type LogService struct {
	destroy.Service
	logger rpi.Logger
}

const name = "destroy"

// ExecuteDF is the logging function attached to the destroy services and responsible for logging it out.
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

// ExecuteSUS is the logging function attached to the destroy services and responsible for logging it out.
func (ls *LogService) ExecuteSUS(ctx echo.Context, processname string, processtype string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute stop user session %v of type %v", processname, processtype), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteSUS(processname, processtype)
}

// ExecuteKP is the logging function attached to the destroy services and responsible for logging it out.
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
