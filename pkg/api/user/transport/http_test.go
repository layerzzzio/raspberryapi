package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/user"
	"github.com/raspibuddy/rpi/pkg/api/user/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/shirou/gopsutil/host"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response []rpi.User

	cases := []struct {
		name         string
		usys         *mocksys.User
		wantedStatus int
		wantedResp   []rpi.User
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			usys: &mocksys.User{
				ListFn: func([]host.UserStat) ([]rpi.User, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			usys: &mocksys.User{
				ListFn: func([]host.UserStat) ([]rpi.User, error) {
					return []rpi.User{
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
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: []rpi.User{
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
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := user.New(tc.usys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/users"
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

func TestView(t *testing.T) {
	var response rpi.User

	cases := []struct {
		name         string
		req          string
		usys         *mocksys.User
		wantedStatus int
		wantedResp   rpi.User
	}{
		{
			name: "error: View result is nil",
			req:  `a`,
			usys: &mocksys.User{
				ViewFn: func(string, []host.UserStat) (rpi.User, error) {
					return rpi.User{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  "T1",
			usys: &mocksys.User{
				ViewFn: func(string, []host.UserStat) (rpi.User, error) {
					return rpi.User{
						User:     "U1",
						Terminal: "T1",
						Started:  11111,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.User{
				User:     "U1",
				Terminal: "T1",
				Started:  11111,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := user.New(tc.usys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/users/" + tc.req
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.User != "" {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
