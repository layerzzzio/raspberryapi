package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/mem"
	"github.com/raspibuddy/rpi/pkg/api/metrics/mem/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	mext "github.com/shirou/gopsutil/mem"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response rpi.Mem

	cases := []struct {
		name         string
		msys         *mocksys.Mem
		wantedStatus int
		wantedResp   rpi.Mem
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			msys: &mocksys.Mem{
				ListFn: func(mext.SwapMemoryStat, mext.VirtualMemoryStat) (rpi.Mem, error) {
					return rpi.Mem{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			msys: &mocksys.Mem{
				ListFn: func(mext.SwapMemoryStat, mext.VirtualMemoryStat) (rpi.Mem, error) {
					return rpi.Mem{
						STotal:       333,
						SUsed:        222,
						SFree:        111,
						SUsedPercent: 66.6,
						VTotal:       333,
						VUsed:        222,
						VAvailable:   111,
						VUsedPercent: 66.6,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Mem{
				STotal:       333,
				SUsed:        222,
				SFree:        111,
				SUsedPercent: 66.6,
				VTotal:       333,
				VUsed:        222,
				VAvailable:   111,
				VUsedPercent: 66.6,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := mem.New(tc.msys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/mems"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if (tc.wantedResp != rpi.Mem{}) {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
