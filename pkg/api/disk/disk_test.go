package disk_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/disk"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	dext "github.com/shirou/gopsutil/disk"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    mock.Metrics
		dsys       mocksys.Disk
		wantedData []rpi.Disk
		wantedErr  error
	}{
		{
			name: "error: dstats is nil",
			metrics: mock.Metrics{
				DiskStatsFn: func(bool) (map[string][]metrics.DStats, error) {
					return nil, errors.New("test error dstats")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not list the disk metrics"),
		},
		{
			name: "success: dstats containing one device and one mountpoint",
			metrics: mock.Metrics{
				DiskStatsFn: func(bool) (map[string][]metrics.DStats, error) {
					return map[string][]metrics.DStats{
						"/dev1": {
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp11",
									Fstype:     "fs11",
									Opts:       "rw11",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp11",
									Fstype:            "fs11",
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
			dsys: mocksys.Disk{
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
			wantedData: []rpi.Disk{
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
				}},
			wantedErr: nil,
		},
		{
			name: "success: dstats containing one device and multiple mountpoints",
			metrics: mock.Metrics{
				DiskStatsFn: func(bool) (map[string][]metrics.DStats, error) {
					return map[string][]metrics.DStats{
						"/dev1": {
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp11",
									Fstype:     "fs11",
									Opts:       "rw11",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp11",
									Fstype:            "fs11",
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
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp12",
									Fstype:     "fs12",
									Opts:       "rw12",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp12",
									Fstype:            "fs12",
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
			dsys: mocksys.Disk{
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
					}, nil
				},
			},
			wantedData: []rpi.Disk{
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
				}},
			wantedErr: nil,
		},
		{
			name: "success: dstats containing multiple devices and multiple mountpoints",
			metrics: mock.Metrics{
				DiskStatsFn: func(bool) (map[string][]metrics.DStats, error) {
					return map[string][]metrics.DStats{
						"/dev1": {
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp11",
									Fstype:     "fs11",
									Opts:       "rw11",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp11",
									Fstype:            "fs11",
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
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp12",
									Fstype:     "fs12",
									Opts:       "rw12",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp12",
									Fstype:            "fs12",
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
						"/dev2": {
							{
								Partition: dext.PartitionStat{
									Device:     "/dev2",
									Mountpoint: "/dev2/mp21",
									Fstype:     "fs21",
									Opts:       "rw21",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev2/mp21",
									Fstype:            "fs21",
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
			dsys: mocksys.Disk{
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
			wantedData: []rpi.Disk{
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
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := disk.New(&tc.dsys, tc.metrics)
			disks, err := s.List()
			assert.Equal(t, tc.wantedData, disks)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		dev        string
		metrics    mock.Metrics
		dsys       mocksys.Disk
		wantedData rpi.Disk
		wantedErr  error
	}{
		{
			name: "error: dstats is nil",
			metrics: mock.Metrics{
				DiskStatsFn: func(bool) (map[string][]metrics.DStats, error) {
					return nil, errors.New("test error dstats")
				},
			},
			wantedData: rpi.Disk{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not view the disk metrics"),
		},
		{
			name: "success: dstats containing one device and one mountpoint",
			dev:  "dev1",
			metrics: mock.Metrics{
				DiskStatsFn: func(bool) (map[string][]metrics.DStats, error) {
					return map[string][]metrics.DStats{
						"/dev1": {
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp11",
									Fstype:     "fs11",
									Opts:       "rw11",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp11",
									Fstype:            "fs11",
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
			dsys: mocksys.Disk{
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
			wantedData: rpi.Disk{
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
			wantedErr: nil,
		},
		{
			name: "success: dstats containing one device and multiple mountpoints",
			dev:  "dev1",
			metrics: mock.Metrics{
				DiskStatsFn: func(bool) (map[string][]metrics.DStats, error) {
					return map[string][]metrics.DStats{
						"/dev1": {
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp11",
									Fstype:     "fs11",
									Opts:       "rw11",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp11",
									Fstype:            "fs11",
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
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp12",
									Fstype:     "fs12",
									Opts:       "rw12",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp12",
									Fstype:            "fs12",
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
			dsys: mocksys.Disk{
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
					}, nil
				},
			},
			wantedData: rpi.Disk{
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
				}},
			wantedErr: nil,
		},
		{
			name: "success: dstats containing multiple devices and multiple mountpoints",
			dev:  "dev2",
			metrics: mock.Metrics{
				DiskStatsFn: func(bool) (map[string][]metrics.DStats, error) {
					return map[string][]metrics.DStats{
						"/dev1": {
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp11",
									Fstype:     "fs11",
									Opts:       "rw11",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp11",
									Fstype:            "fs11",
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
							{
								Partition: dext.PartitionStat{
									Device:     "/dev1",
									Mountpoint: "/dev1/mp12",
									Fstype:     "fs12",
									Opts:       "rw12",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev1/mp12",
									Fstype:            "fs12",
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
						"/dev2": {
							{
								Partition: dext.PartitionStat{
									Device:     "/dev2",
									Mountpoint: "/dev2/mp21",
									Fstype:     "fs21",
									Opts:       "rw21",
								},
								Mountpoint: &dext.UsageStat{
									Path:              "/dev2/mp21",
									Fstype:            "fs21",
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
			dsys: mocksys.Disk{
				ViewFn: func(string, map[string][]metrics.DStats) (rpi.Disk, error) {
					return rpi.Disk{
						ID:         "dev2",
						Filesystem: "/dev2",
						Fstype:     "fs21",
						Mountpoints: []rpi.MountPoint{
							{
								Mountpoint:        "/dev2/mp21",
								Fstype:            "fs21",
								Opts:              "rw21",
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
			wantedData: rpi.Disk{
				ID:         "dev2",
				Filesystem: "/dev2",
				Fstype:     "fs21",
				Mountpoints: []rpi.MountPoint{
					{
						Mountpoint:        "/dev2/mp21",
						Fstype:            "fs21",
						Opts:              "rw21",
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
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := disk.New(&tc.dsys, tc.metrics)
			disks, err := s.View(tc.dev)
			assert.Equal(t, tc.wantedData, disks)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
