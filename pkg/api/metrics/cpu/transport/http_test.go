package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/cpu"
	"github.com/raspibuddy/rpi/pkg/api/metrics/cpu/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	cext "github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response []rpi.CPU

	cases := []struct {
		name         string
		csys         *mocksys.CPU
		wantedStatus int
		wantedResp   []rpi.CPU
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			csys: &mocksys.CPU{
				ListFn: func([]cext.InfoStat, []float64, []cext.TimesStat) ([]rpi.CPU, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			csys: &mocksys.CPU{
				ListFn: func([]cext.InfoStat, []float64, []cext.TimesStat) ([]rpi.CPU, error) {
					return []rpi.CPU{
						{
							ID:        1,
							Cores:     int32(8),
							ModelName: "intel processor",
							Mhz:       2300.99,
							Stats: rpi.CPUStats{
								Used:   99.9,
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: []rpi.CPU{
				{
					ID:        1,
					Cores:     int32(8),
					ModelName: "intel processor",
					Mhz:       2300.99,
					Stats: rpi.CPUStats{
						Used:   99.9,
						User:   111.1,
						System: 222.2,
						Idle:   333.3,
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := cpu.New(tc.csys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/cpus"
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
	var response rpi.CPU

	cases := []struct {
		name         string
		req          string
		csys         *mocksys.CPU
		wantedStatus int
		wantedResp   rpi.CPU
	}{
		{
			name:         "error: invalid request",
			wantedStatus: http.StatusBadRequest,
			req:          `a`,
		},
		{
			name: "error: View result is nil",
			req:  "1",
			csys: &mocksys.CPU{
				ViewFn: func(int, []cext.InfoStat, []float64, []cext.TimesStat) (rpi.CPU, error) {
					return rpi.CPU{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  "1",
			csys: &mocksys.CPU{
				ViewFn: func(int, []cext.InfoStat, []float64, []cext.TimesStat) (rpi.CPU, error) {
					return rpi.CPU{
						ID:        1,
						Cores:     int32(8),
						ModelName: "intel processor",
						Mhz:       2300.99,
						Stats: rpi.CPUStats{
							Used:   99.9,
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.CPU{
				ID:        1,
				Cores:     int32(8),
				ModelName: "intel processor",
				Mhz:       2300.99,
				Stats: rpi.CPUStats{
					Used:   99.9,
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
			s := cpu.New(tc.csys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/cpus/" + tc.req
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
