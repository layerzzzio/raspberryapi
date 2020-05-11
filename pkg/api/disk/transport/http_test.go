package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/disk"
	"github.com/raspibuddy/rpi/pkg/api/disk/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var listResponse []rpi.Disk

	cases := []struct {
		name         string
		dsys         *mocksys.Disk
		wantedStatus int
		wantedResp   []rpi.Disk
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			dsys: &mocksys.Disk{
				ListFn: func(map[string][]metrics.DStats) ([]rpi.Disk, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			dsys: &mocksys.Disk{
				ListFn: func(map[string][]metrics.DStats) ([]rpi.Disk, error) {
					return []rpi.Disk{
						{
							ID:         "dev1",
							Filesystem: "/dev1",
							Fstype:     "fs11",
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
							},
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: []rpi.Disk{
				{
					ID:         "dev1",
					Filesystem: "/dev1",
					Fstype:     "fs11",
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
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			s := disk.New(tc.dsys, metrics.Service{})
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/disks"
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

func TestView(t *testing.T) {
	var listResponse rpi.Disk

	cases := []struct {
		name         string
		req          string
		dsys         *mocksys.Disk
		wantedStatus int
		wantedResp   rpi.Disk
	}{
		{
			name: "error: View result is nil",
			req:  `a`,
			dsys: &mocksys.Disk{
				ViewFn: func(string, map[string][]metrics.DStats) (rpi.Disk, error) {
					return rpi.Disk{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  "dev1",
			dsys: &mocksys.Disk{
				ViewFn: func(string, map[string][]metrics.DStats) (rpi.Disk, error) {
					return rpi.Disk{
						ID:         "dev1",
						Filesystem: "/dev1",
						Fstype:     "fs11",
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
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Disk{
				ID:         "dev1",
				Filesystem: "/dev1",
				Fstype:     "fs11",
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
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			s := disk.New(tc.dsys, metrics.Service{})
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/disks/" + tc.req
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
