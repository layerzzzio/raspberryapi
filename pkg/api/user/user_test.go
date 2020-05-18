package user_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/user"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/shirou/gopsutil/host"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    *mock.Metrics
		usys       *mocksys.User
		wantedData []rpi.User
		wantedErr  error
	}{
		{
			name: "error: users array is nil",
			metrics: &mock.Metrics{
				UsersFn: func() ([]host.UserStat, error) {
					return nil, errors.New("test error info")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not list the user metrics"),
		},
		{
			name: "success",
			metrics: &mock.Metrics{
				UsersFn: func() ([]host.UserStat, error) {
					return []host.UserStat{
						{
							User:     "U1",
							Terminal: "T1",
							Started:  11111,
						},
					}, nil
				},
			},
			usys: &mocksys.User{
				ListFn: func([]host.UserStat) ([]rpi.User, error) {
					return []rpi.User{
						{
							User:     "U1",
							Terminal: "T1",
							Started:  11111,
						},
					}, nil
				},
			},
			wantedData: []rpi.User{
				{
					User:     "U1",
					Terminal: "T1",
					Started:  11111,
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := user.New(tc.usys, tc.metrics)
			users, err := s.List()
			assert.Equal(t, tc.wantedData, users)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		terminal   string
		metrics    *mock.Metrics
		usys       *mocksys.User
		wantedData rpi.User
		wantedErr  error
	}{
		{
			name: "error: users array is nil",
			metrics: &mock.Metrics{
				UsersFn: func() ([]host.UserStat, error) {
					return nil, errors.New("test error info")
				},
			},
			wantedData: rpi.User{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not view the user metrics"),
		},
		{
			name:     "success",
			terminal: "T1",
			metrics: &mock.Metrics{
				UsersFn: func() ([]host.UserStat, error) {
					return []host.UserStat{
						{
							User:     "U1",
							Terminal: "T1",
							Started:  11111,
						},
					}, nil
				},
			},
			usys: &mocksys.User{
				ViewFn: func(string, []host.UserStat) (rpi.User, error) {
					return rpi.User{
						User:     "U1",
						Terminal: "T1",
						Started:  11111,
					}, nil
				},
			},
			wantedData: rpi.User{
				User:     "U1",
				Terminal: "T1",
				Started:  11111,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := user.New(tc.usys, tc.metrics)
			users, err := s.View(tc.terminal)
			assert.Equal(t, tc.wantedData, users)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
