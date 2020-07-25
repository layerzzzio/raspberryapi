package mem

import (
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/mem"
)

// New creates a new mem logging service instance.
func New(svc mem.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a mem logging service.
type LogService struct {
	mem.Service
	logger rpi.Logger
}

const name = "mem"

// List is the logging function attached to the List mem services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp rpi.Mem, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing mems", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
