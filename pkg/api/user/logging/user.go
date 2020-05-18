package user

import (
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/user"
)

// New creates a new user logging service instance.
func New(svc user.Service, logger rpi.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents a user logging service.
type LogService struct {
	user.Service
	logger rpi.Logger
}

const name = "user"

// List is the logging function attached to the List user services and responsible for logging it out.
func (ls *LogService) List(ctx echo.Context) (resp []rpi.User, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			name, "request: listing users", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List()
}
