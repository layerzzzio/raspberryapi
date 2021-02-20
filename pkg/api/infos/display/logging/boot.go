package boot

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/boot"
)

// New creates a new boot logging service instance.
func New(svc boot.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a boot logging service.
type LogService struct {
	boot.Service
	logger rpi.Logger
}

const name = "boot"

// List is the logging function attached to the List boot services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.Boot, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing boot config", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
