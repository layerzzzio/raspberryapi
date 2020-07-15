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
	"github.com/shirou/gopsutil/load"
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
					load.AvgStat,
					string,
					string,
					map[string][]metrics.DStats) (rpi.Host, error) {
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
					load.AvgStat,
					string,
					string,
					map[string][]metrics.DStats) (rpi.Host, error) {
					return rpi.Host{
						ID:                 "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
						Hostname:           "hostname_test",
						UpTime:             540165,
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
						Load1:              1,
						Load5:              5,
						Load15:             15,
						Processes:          400,
						ActiveVirtualUsers: 2,
						Temperature:        20.9,
						RaspModel:          "pi zero",
						Disks: []rpi.Disk{
							{
								ID:         "dev1",
								Filesystem: "/dev1",
								Fstype:     "multi_fstype",
								Mountpoints: []rpi.MountPoint{
									{
										Mountpoint:        "/dev1/mp11",
										Fstype:            "fs11",
										Opts:              "rw11",
										Total:             1,
										Free:              2,
										Used:              3,
										UsedPercent:       4.4,
										InodesTotal:       5,
										InodesUsed:        6,
										InodesFree:        7,
										InodesUsedPercent: 8.8,
									},
									{
										Mountpoint:        "/dev1/mp12",
										Fstype:            "fs12",
										Opts:              "rw12",
										Total:             1,
										Free:              2,
										Used:              3,
										UsedPercent:       4.4,
										InodesTotal:       5,
										InodesUsed:        6,
										InodesFree:        7,
										InodesUsedPercent: 8.8,
									},
								},
							},
							{
								ID:         "dev2",
								Filesystem: "/dev2",
								Fstype:     "fs21",
								Mountpoints: []rpi.MountPoint{
									{
										Mountpoint:        "/dev2/mp21",
										Fstype:            "fs21",
										Opts:              "rw21",
										Total:             11,
										Free:              22,
										Used:              33,
										UsedPercent:       44.4,
										InodesTotal:       55,
										InodesUsed:        66,
										InodesFree:        77,
										InodesUsedPercent: 88.8,
									},
								},
							},
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Host{
				ID:                 "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:           "hostname_test",
				UpTime:             540165,
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
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          400,
				ActiveVirtualUsers: 2,
				Temperature:        20.9,
				RaspModel:          "pi zero",
				Disks: []rpi.Disk{
					{
						ID:         "dev1",
						Filesystem: "/dev1",
						Fstype:     "multi_fstype",
						Mountpoints: []rpi.MountPoint{
							{
								Mountpoint:        "/dev1/mp11",
								Fstype:            "fs11",
								Opts:              "rw11",
								Total:             1,
								Free:              2,
								Used:              3,
								UsedPercent:       4.4,
								InodesTotal:       5,
								InodesUsed:        6,
								InodesFree:        7,
								InodesUsedPercent: 8.8,
							},
							{
								Mountpoint:        "/dev1/mp12",
								Fstype:            "fs12",
								Opts:              "rw12",
								Total:             1,
								Free:              2,
								Used:              3,
								UsedPercent:       4.4,
								InodesTotal:       5,
								InodesUsed:        6,
								InodesFree:        7,
								InodesUsedPercent: 8.8,
							},
						},
					},
					{
						ID:         "dev2",
						Filesystem: "/dev2",
						Fstype:     "fs21",
						Mountpoints: []rpi.MountPoint{
							{
								Mountpoint:        "/dev2/mp21",
								Fstype:            "fs21",
								Opts:              "rw21",
								Total:             11,
								Free:              22,
								Used:              33,
								UsedPercent:       44.4,
								InodesTotal:       55,
								InodesUsed:        66,
								InodesFree:        77,
								InodesUsedPercent: 88.8,
							},
						},
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

			if tc.wantedResp.ID != "" {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
