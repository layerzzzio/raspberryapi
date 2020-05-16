package transport_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/process"
	"github.com/raspibuddy/rpi/pkg/api/process/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response []rpi.Process

	cases := []struct {
		name         string
		psys         *mocksys.Process
		wantedStatus int
		wantedResp   []rpi.Process
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			psys: &mocksys.Process{
				ListFn: func([]metrics.PInfo) ([]rpi.Process, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		// {
		// 	name: "success",
		// 	psys: &mocksys.Process{
		// 		ListFn: func([]metrics.PInfo) ([]rpi.Process, error) {
		// 			return []rpi.Process{
		// 				{
		// 					ID:         1,
		// 					Name:       "process_1",
		// 					CPUPercent: 1.1,
		// 					MemPercent: 2.2,
		// 				},
		// 				{
		// 					ID:         2,
		// 					Name:       "process_2",
		// 					CPUPercent: 3.3,
		// 					MemPercent: 4.4,
		// 				},
		// 			}, nil
		// 		},
		// 	},
		// 	wantedStatus: http.StatusOK,
		// 	wantedResp: []rpi.Process{
		// 		{
		// 			ID:         1,
		// 			Name:       "process_1",
		// 			CPUPercent: 1.1,
		// 			MemPercent: 2.2,
		// 		},
		// 		{
		// 			ID:         2,
		// 			Name:       "process_2",
		// 			CPUPercent: 3.3,
		// 			MemPercent: 4.4,
		// 		},
		// 	},
		// },
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			s := process.New(tc.psys, metrics.Service{})
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/processes"
			res, err := http.Get(path)

			if err != nil {
				t.Fatal(err)
			}

			fmt.Println(res)

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
	var response rpi.Process

	cases := []struct {
		name         string
		req          string
		psys         *mocksys.Process
		wantedStatus int
		wantedResp   rpi.Process
	}{
		{
			name:         "error: invalid request",
			wantedStatus: http.StatusBadRequest,
			req:          `a`,
		},
		// {
		// 	name: "error: View result is nil",
		// 	req:  "1",
		// 	psys: &mocksys.Process{
		// 		ViewFn: func(int32, []metrics.PInfo) (rpi.Process, error) {
		// 			return rpi.Process{}, errors.New("test error")
		// 		},
		// 	},
		// 	wantedStatus: http.StatusInternalServerError,
		// },
		{
			name: "success",
			req:  "99",
			psys: &mocksys.Process{
				ViewFn: func(int32, []metrics.PInfo) (rpi.Process, error) {
					return rpi.Process{
						ID:           int32(99),
						Name:         "process_99",
						CPUPercent:   1.1,
						MemPercent:   2.2,
						Username:     "pi",
						CommandLine:  "/cmd/text",
						Status:       "S",
						CreationTime: 1666666,
						Foreground:   true,
						Background:   false,
						IsRunning:    true,
						ParentP:      int32(1),
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Process{
				ID:           int32(99),
				Name:         "process_99",
				CPUPercent:   1.1,
				MemPercent:   2.2,
				Username:     "pi",
				CommandLine:  "/cmd/text",
				Status:       "S",
				CreationTime: 1666666,
				Foreground:   true,
				Background:   false,
				IsRunning:    true,
				ParentP:      int32(1),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			s := process.New(tc.psys, metrics.Service{})
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/processes/" + tc.req
			res, err := http.Get(path)

			if err != nil {
				t.Fatal(err)
			}

			fmt.Println(res)

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
