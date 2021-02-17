package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/humanuser"
	"github.com/raspibuddy/rpi/pkg/api/infos/humanuser/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response []rpi.HumanUser

	cases := []struct {
		name         string
		humsys       *mocksys.HumanUser
		wantedStatus int
		wantedResp   []rpi.HumanUser
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			humsys: &mocksys.HumanUser{
				ListFn: func([]string) ([]rpi.HumanUser, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			humsys: &mocksys.HumanUser{
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
			wantedStatus: http.StatusOK,
			wantedResp: []rpi.HumanUser{
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
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := humanuser.New(tc.humsys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/humanusers"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp != nil {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
