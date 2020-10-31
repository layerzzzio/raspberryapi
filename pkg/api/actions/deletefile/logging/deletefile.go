package deletefile

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/deletefile"
)

// New creates a new DeleteFile logging service instance.
func New(svc deletefile.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a DeleteFile logging service.
type LogService struct {
	deletefile.Service
	logger rpi.Logger
}

const name = "deletefile"

// Execute is the logging function attached to the Execute deletefile services and responsible for logging it out.
func (ls *LogService) Execute(ctx echo.Context, path string) (resp rpi.Action, err error) {
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
	return ls.Service.Execute(path)
}
