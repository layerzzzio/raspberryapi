package host

import (
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/host"
)

// New creates a new host logging service instance.
func New(svc host.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a host logging service.
type LogService struct {
	host.Service
	logger rpi.Logger
}

const name = "host"

// List is the logging function attached to the List host services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.Host, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing hosts", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
