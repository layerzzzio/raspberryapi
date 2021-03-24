package transport_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/install"
	"github.com/raspibuddy/rpi/pkg/api/actions/install/transport"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestExecuteAG(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		inssys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: no package",
			req:          "?action=install&pkg=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: no action type",
			req:          "?action=&pkg=openvpn",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action",
			req:          "?action=installxxx&pkg=openvpn",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteAG result is nil",
			req:  "?action=install&pkg=openvpn",
			inssys: &mocksys.Action{
				ExecuteAGFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=install&pkg=openvpn",
			inssys: &mocksys.Action{
				ExecuteAGFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.InstallAptGet,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
					}, nil
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			a := actions.New()
			s := install.New(tc.inssys, a)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/install/aptget" + tc.req

			res, err := http.Post(path, "application/json", bytes.NewBufferString(tc.req))
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.Name != "" {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}

func TestExecuteNV(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		inssys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: no action type",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action",
			req:          "?action=installxxx",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteNV result is nil",
			req:  "?action=install",
			inssys: &mocksys.Action{
				ExecuteNVFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=install",
			inssys: &mocksys.Action{
				ExecuteNVFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.InstallNordVPN,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
					}, nil
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			a := actions.New()
			s := install.New(tc.inssys, a)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/install/nordvpn" + tc.req

			res, err := http.Post(path, "application/json", bytes.NewBufferString(tc.req))
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.Name != "" {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
