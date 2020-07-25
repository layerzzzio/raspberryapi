package user

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of User model.
func (u *User) List() ([]rpi.User, error) {
	users, err := u.mt.Users()

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not list the user metrics")
	}

	return u.usys.List(users)
}

// View populates and returns a User model.
func (u *User) View(terminal string) (rpi.User, error) {
	users, err := u.mt.Users()

	if err != nil {
		return rpi.User{}, echo.NewHTTPError(http.StatusInternalServerError, "could not view the user metrics")
	}

	return u.usys.View(terminal, users)
}
