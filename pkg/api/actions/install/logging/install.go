package install

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

// LogService represents a Install logging service.
type LogService struct {
	configure.Service
	logger rpi.Logger
}

const name = "install"

// ExecuteCH is the logging function attached to the execute change hostname service and responsible for logging it out.
func (ls *LogService) ExecuteCH(ctx echo.Context, hostname string) (resp rpi.Action, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: execute install openvpn #%v", hostname), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ExecuteCH(hostname)
}
