package vcore

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/vcore"
)

// New creates a new vcore logging service.
func New(svc vcore.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a cpu logging service.
type LogService struct {
	vcore.Service
	logger rpi.Logger
}

const name = "vcore"

// List logs the requests when listing the vcores.
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

// View logs the request when listing a specific vcore.
func (ls *LogService) View(ctx echo.Context, id int) (resp *rpi.VCore, err error) {
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
