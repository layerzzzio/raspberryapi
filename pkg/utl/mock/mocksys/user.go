package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/host"
)

// User mock
type User struct {
	ListFn func([]host.UserStat) ([]rpi.User, error)
	ViewFn func(string, []host.UserStat) (rpi.User, error)
}

// List mock
func (u User) List(users []host.UserStat) ([]rpi.User, error) {
	return u.ListFn(users)
}

// View mock
func (u User) View(terminal string, users []host.UserStat) (rpi.User, error) {
	return u.ViewFn(terminal, users)
}
