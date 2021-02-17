package humanuser_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/humanuser"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		humsys     mocksys.HumanUser
		wantedData []rpi.HumanUser
		wantedErr  error
	}{
		{
			name: "error: lines are nil",
			infos: mock.Infos{
				ReadFileFn: func(string) ([]string, error) {
					return nil, errors.New("test error info")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the human users"),
		},
		{
			name: "success",
			infos: mock.Infos{
				ReadFileFn: func(string) ([]string, error) {
					return []string{
						"systemd-network:x:101:103:systemd Network Management,,,:/run/systemd:/usr/sbin/nologin",
						"systemd-resolve:x:102:104:systemd Resolver,,,:/run/systemd:/usr/sbin/nologin",
						"_apt:x:103:65534::/nonexistent:/usr/sbin/nologin",
						"pi:x:1000:1000:,,,:/home/pi:/bin/bash",
						"#pi2:x:1001:1001:,,,:/home/pi2:/bin/bash",
						"messagebus:x:104:110::/nonexistent:/usr/sbin/nologin",
						"_rpc:x:105:65534::/run/rpcbind:/usr/sbin/nologin",
						"statd:x:106:65534::/var/lib/nfs:/usr/sbin/nologin",
					}, nil
				},
			},
			humsys: mocksys.HumanUser{
				ListFn: func([]string) ([]rpi.HumanUser, error) {
					return []rpi.HumanUser{
						{
							Username:       "pi",
							Password:       "x",
							Uid:            1000,
							Gid:            1000,
							AdditionalInfo: nil,
							HomeDirectory:  "/home/pi",
							DefaultShell:   "/bin/bash",
						},
					}, nil
				},
			},
			wantedData: []rpi.HumanUser{
				{
					Username:       "pi",
					Password:       "x",
					Uid:            1000,
					Gid:            1000,
					AdditionalInfo: nil,
					HomeDirectory:  "/home/pi",
					DefaultShell:   "/bin/bash",
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := humanuser.New(&tc.humsys, tc.infos)
			humanUsers, err := s.List()
			assert.Equal(t, tc.wantedData, humanUsers)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
