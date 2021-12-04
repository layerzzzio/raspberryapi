package transport_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/general"
	"github.com/raspibuddy/rpi/pkg/api/actions/general/transport"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestExecuteRBS(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		gensys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteRBS result is nil",
			req:  "?option=reboot",
			gensys: &mocksys.Action{
				ExecuteRBSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?option=reboot",
			gensys: &mocksys.Action{
				ExecuteRBSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.RebootShutdown,
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
			s := general.New(tc.gensys, a)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/general/boot" + tc.req

			fmt.Println(path)

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

func TestBootOptionCheck(t *testing.T) {
	cases := []struct {
		name       string
		action     string
		wantedResp error
	}{
		{
			name:       "error: bad action type",
			action:     "rebootxxx",
			wantedResp: echo.NewHTTPError(http.StatusNotFound, "Not found - bad option type"),
		},
		{
			name:       "error: null action type",
			action:     "xxxshutdown",
			wantedResp: echo.NewHTTPError(http.StatusNotFound, "Not found - bad option type"),
		},
		{
			name:       "success: reboot",
			action:     "reboot",
			wantedResp: nil,
		},
		{
			name:       "success: shutdown",
			action:     "shutdown",
			wantedResp: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := transport.BootOptionCheck(tc.action, `reboot|shutdown`)
			assert.Equal(t, tc.wantedResp, res)
		})
	}
}

func TestExecuteSASO(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		gensys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response",
			req:          "",
			wantedStatus: http.StatusBadRequest,
		},
		{
			name:         "error: bad action",
			req:          "?action=stoooop&service=raspibuddy_deploy",
			wantedStatus: http.StatusBadRequest,
		},
		{
			name:         "error: no service",
			req:          "?action=stop&service=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteSASO result is nil",
			req:  "?action=start&service=raspibuddy_deploy",
			gensys: &mocksys.Action{
				ExecuteSASOFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=start&service=raspibuddy_deploy",
			gensys: &mocksys.Action{
				ExecuteSASOFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.RebootShutdown,
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
			s := general.New(tc.gensys, a)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/general/systemctl" + tc.req

			fmt.Println(path)

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
