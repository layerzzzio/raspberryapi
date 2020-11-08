package sys

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/host"
)

// User represents an empty User entity on the current system.
type User struct{}

// List returns a list of User stats
func (u User) List(users []host.UserStat) ([]rpi.User, error) {
	var result []rpi.User

	for i := range users {
		result = append(result, rpi.User{
			User:     users[i].User,
			Terminal: users[i].Terminal,
			Started:  users[i].Started,
		})
	}

	return result, nil
}

// View returns a list of User stats
func (u User) View(terminal string, users []host.UserStat) (rpi.User, error) {
	var result rpi.User
	isTerminalFound := false

	for i := range users {
		if terminal == users[i].Terminal {
			result = rpi.User{
				User:     users[i].User,
				Terminal: users[i].Terminal,
				Started:  users[i].Started,
			}
			isTerminalFound = true
			break
		}
	}

	if !isTerminalFound {
		return rpi.User{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%v does not exist", terminal))
	}

	return result, nil
}
