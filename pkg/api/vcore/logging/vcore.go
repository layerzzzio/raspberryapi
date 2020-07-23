package vcore

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/vcore"
)

// New creates a new vCore logging service instance.
func New(svc vcore.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a vCore logging service.
type LogService struct {
	vcore.Service
	logger rpi.Logger
}

const name = "vcore"

// List is the logging function attached to the List vCore services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.VCore, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing all the vcores", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}

// View is the logging function attached to the View vCore services and responsible for logging it out.
func (ls *LogService) View(ctx echo.Context, id int) (resp rpi.VCore, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: viewing vcore #%v", id), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(id)
}
