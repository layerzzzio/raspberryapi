package user

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/host"
)

// Service represents all User application services.
type Service interface {
	List() ([]rpi.User, error)
	View(string) (rpi.User, error)
}

// User represents a User application service.
type User struct {
	usys USYS
	mt   Metrics
}

// USYS represents a User repository service.
type USYS interface {
	List([]host.UserStat) ([]rpi.User, error)
	View(string, []host.UserStat) (rpi.User, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	Users() ([]host.UserStat, error)
}

// New creates a User application service instance.
func New(usys USYS, mt Metrics) *User {
	return &User{usys: usys, mt: mt}
}
