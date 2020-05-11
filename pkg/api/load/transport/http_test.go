package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/load"
	"github.com/raspibuddy/rpi/pkg/api/load/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	lext "github.com/shirou/gopsutil/load"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var listResponse rpi.Load

	cases := []struct {
		name         string
		lsys         *mocksys.Load
		wantedStatus int
		wantedResp   rpi.Load
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			lsys: &mocksys.Load{
				ListFn: func(lext.AvgStat, lext.MiscStat) (rpi.Load, error) {
					return rpi.Load{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			lsys: &mocksys.Load{
				ListFn: func(lext.AvgStat, lext.MiscStat) (rpi.Load, error) {
					return rpi.Load{
						Load1:        1.1,
						Load5:        5.5,
						Load15:       15.15,
						ProcsTotal:   4,
						ProcsBlocked: 3,
						ProcsRunning: 1,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Load{
				Load1:        1.1,
				Load5:        5.5,
				Load15:       15.15,
				ProcsTotal:   4,
				ProcsBlocked: 3,
				ProcsRunning: 1,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			s := load.New(tc.lsys, metrics.Service{})
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/loads"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if (tc.wantedResp != rpi.Load{}) {
				response := listResponse
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
