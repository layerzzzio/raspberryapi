package sys_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/user"
	"github.com/raspibuddy/rpi/pkg/api/user/platform/sys"
	"github.com/shirou/gopsutil/host"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		users      []host.UserStat
		wantedData []rpi.User
		wantedErr  error
	}{
		{
			name:       "success: users array is empty",
			users:      []host.UserStat{},
			wantedData: nil,
			wantedErr:  nil,
		},
		{
			name: "success: users array not empty",
			users: []host.UserStat{
				{
					User:     "U1",
					Terminal: "T1",
					Started:  11111,
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
			s := user.USYS(sys.User{})
			users, err := s.List(tc.users)
			assert.Equal(t, tc.wantedData, users)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		terminal   string
		users      []host.UserStat
		wantedData rpi.User
		wantedErr  error
	}{
		{
			name:     "error: terminal does not exist",
			terminal: "T2",
			users: []host.UserStat{
				{
					User:     "U1",
					Terminal: "T1",
					Started:  11111,
				},
			},
			wantedData: rpi.User{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "T2 does not exist"),
		},
		{
			name:     "success",
			terminal: "T1",
			users: []host.UserStat{
				{
					User:     "U1",
					Terminal: "T1",
					Started:  11111,
				},
				{
					User:     "U2",
					Terminal: "T2",
					Started:  22222,
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
			s := user.USYS(sys.User{})
			users, err := s.View(tc.terminal, tc.users)
			assert.Equal(t, tc.wantedData, users)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
