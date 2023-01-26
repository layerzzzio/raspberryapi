package net

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/net"
)

// New creates a new net logging service instance.
func New(svc net.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a net logging service.
type LogService struct {
	net.Service
	logger rpi.Logger
}

const name = "net"

// List is the logging function attached to the List net services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.Net, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing net interfaces", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}

// View is the logging function attached to the View net services and responsible for logging it out.
func (ls *LogService) View(ctx echo.Context, id int) (resp rpi.Net, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, fmt.Sprintf("request: viewing net interfaces #%v", id), err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(id)
}
