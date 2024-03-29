package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/vcore"
	"github.com/raspibuddy/rpi/pkg/api/metrics/vcore/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response []rpi.VCore

	cases := []struct {
		name         string
		vsys         *mocksys.VCore
		wantedStatus int
		wantedResp   []rpi.VCore
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			vsys: &mocksys.VCore{
				ListFn: func([]float64, []cpu.TimesStat) ([]rpi.VCore, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			vsys: &mocksys.VCore{
				ListFn: func([]float64, []cpu.TimesStat) ([]rpi.VCore, error) {
					return []rpi.VCore{
						{
							ID:     0,
							Used:   50.0,
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: []rpi.VCore{
				{
					ID:     0,
					Used:   50.0,
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := vcore.New(tc.vsys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/vcores"
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
	var response rpi.VCore

	cases := []struct {
		name         string
		req          string
		vsys         *mocksys.VCore
		wantedStatus int
		wantedResp   rpi.VCore
	}{
		{
			name:         "error: invalid request",
			wantedStatus: http.StatusBadRequest,
			req:          `a`,
		},
		{
			name: "error: View result is nil",
			req:  "1",
			vsys: &mocksys.VCore{
				ViewFn: func(int, []float64, []cpu.TimesStat) (rpi.VCore, error) {
					return rpi.VCore{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  "1",
			vsys: &mocksys.VCore{
				ViewFn: func(int, []float64, []cpu.TimesStat) (rpi.VCore, error) {
					return rpi.VCore{
						ID:     0,
						Used:   50.0,
						User:   111.1,
						System: 222.2,
						Idle:   333.3,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.VCore{
				ID:     0,
				Used:   50.0,
				User:   111.1,
				System: 222.2,
				Idle:   333.3,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := vcore.New(tc.vsys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/vcores/" + tc.req
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.ID > 0 {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
