package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/host"
	"github.com/raspibuddy/rpi/pkg/api/host/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/shirou/gopsutil/cpu"
	hext "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response rpi.Host

	cases := []struct {
		name         string
		hsys         *mocksys.Host
		wantedStatus int
		wantedResp   rpi.Host
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			hsys: &mocksys.Host{
				ListFn: func(
					hext.InfoStat,
					[]hext.UserStat,
					[]cpu.InfoStat,
					[]float64,
					mem.VirtualMemoryStat,
					mem.SwapMemoryStat,
					string) (rpi.Host, error) {
					return rpi.Host{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			hsys: &mocksys.Host{
				ListFn: func(
					hext.InfoStat,
					[]hext.UserStat,
					[]cpu.InfoStat,
					[]float64,
					mem.VirtualMemoryStat,
					mem.SwapMemoryStat,
					string) (rpi.Host, error) {
					return rpi.Host{
						ID:                 "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
						Hostname:           "hostname_test",
						Uptime:             540165,
						BootTime:           1589223156,
						OS:                 "raspbian",
						Platform:           "plat_1",
						PlatformFamily:     "plat_1_1",
						PlatformVersion:    "1.1",
						KernelArch:         "arch_A",
						KernelVersion:      "A",
						CPU:                1,
						HyperThreading:     false,
						VCore:              3,
						CPUUsedPercent:     2.0,
						VUsedPercent:       99.9,
						SUsedPercent:       0.9,
						Processes:          400,
						ActiveVirtualUsers: 2,
						Temperature:        20.9,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Host{
				ID:                 "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:           "hostname_test",
				Uptime:             540165,
				BootTime:           1589223156,
				OS:                 "raspbian",
				Platform:           "plat_1",
				PlatformFamily:     "plat_1_1",
				PlatformVersion:    "1.1",
				KernelArch:         "arch_A",
				KernelVersion:      "A",
				CPU:                1,
				HyperThreading:     false,
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Processes:          400,
				ActiveVirtualUsers: 2,
				Temperature:        20.9,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := host.New(tc.hsys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/hosts"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if (tc.wantedResp != rpi.Host{}) {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
