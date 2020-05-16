package sys

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/disk"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	dext "github.com/shirou/gopsutil/disk"
	"github.com/stretchr/testify/assert"
)

func TestExtractDeviceID(t *testing.T) {
	cases := []struct {
		name       string
		dev        string
		wantedData []string
		wantedErr  error
	}{
		{
			name:       "success: empty input string",
			dev:        "",
			wantedData: nil,
		},
		{
			name:       "success: input string is a slash",
			dev:        "/",
			wantedData: nil,
		},
		{
			name:       "success: alphanumeric input without slashes",
			dev:        "mp1",
			wantedData: []string{"mp1"},
		},
		{
			name:       "success: alphanumeric input including slashes",
			dev:        "/dev1/mp1",
			wantedData: []string{"mp1"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			id := extractDeviceID(tc.dev)
			assert.Equal(t, tc.wantedData, id)
		})
	}
}

func TestFsType(t *testing.T) {
	cases := []struct {
		name       string
		mp         []rpi.MountPoint
		wantedData string
	}{
		{
			name: "success: one mountpoint including one fstype",
			mp: []rpi.MountPoint{
				{
					Mountpoint: "/dev1/mp1",
					Fstype:     "fs1",
				},
			},
			wantedData: "fs1",
		},
		{
			name: "success: two mountpoints including two fstypes",
			mp: []rpi.MountPoint{
				{
					Mountpoint: "/dev1/mp1",
					Fstype:     "fs1",
				},
				{
					Mountpoint: "/dev1/mp2",
					Fstype:     "fs2",
				},
			},
			wantedData: "multi_fstype",
		},
		{
			name: "success: multiple mountpoints including multiples fstypes",
			mp: []rpi.MountPoint{
				{
					Mountpoint: "/dev1/mp1",
					Fstype:     "fs1",
				},
				{
					Mountpoint: "/dev1/mp2",
					Fstype:     "fs2",
				},
				{
					Mountpoint: "/dev1/mp3",
					Fstype:     "fs3",
				},
				{
					Mountpoint: "/dev1/mp3",
					Fstype:     "fs3",
				},
			},
			wantedData: "multi_fstype",
		},
		{
			name: "success: multiple mountpoints input including the same fstype",
			mp: []rpi.MountPoint{
				{
					Mountpoint: "/dev1/mp1",
					Fstype:     "fs1",
				},
				{
					Mountpoint: "/dev1/mp2",
					Fstype:     "fs1",
				},
				{
					Mountpoint: "/dev1/mp3",
					Fstype:     "fs1",
				},
			},
			wantedData: "fs1",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			id := fsType(tc.mp)
			assert.Equal(t, tc.wantedData, id)
		})
	}
}

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		dstats     map[string][]metrics.DStats
		wantedData []rpi.Disk
		wantedErr  error
	}{
		{
			name: "error: parsing id was unsuccessful",
			dstats: map[string][]metrics.DStats{
				"/": {
					{
						Partition: &dext.PartitionStat{
							Device:     "/",
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
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful"),
		},
		{
			name: "success: one device containing one mount point",
			dstats: map[string][]metrics.DStats{
				"/dev1": {
					{
						Partition: &dext.PartitionStat{
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
			name: "success: two devices, each containing one mount point",
			dstats: map[string][]metrics.DStats{
				"/dev1": {
					{
						Partition: &dext.PartitionStat{
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
				"/dev2": {
					{
						Partition: &dext.PartitionStat{
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
		{
			name: "success: multiple devices containing multiple mount points",
			dstats: map[string][]metrics.DStats{
				"/dev1": {
					{
						Partition: &dext.PartitionStat{
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
						Partition: &dext.PartitionStat{
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
						Partition: &dext.PartitionStat{
							Device:     "/dev2",
							Mountpoint: "/dev2/mp21",
							Fstype:     "fs21",
							Opts:       "rw21",
						},
						Mountpoint: &dext.UsageStat{
							Path:              "/dev2/mp21",
							Fstype:            "fs21",
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
			wantedData: []rpi.Disk{
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
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := disk.DSYS(Disk{})
			disk, err := s.List(tc.dstats)
			assert.Equal(t, tc.wantedData, disk)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		dev        string
		dstats     map[string][]metrics.DStats
		wantedData rpi.Disk
		wantedErr  error
	}{
		{
			name: "error: parsing id was unsuccessful",
			dev:  "/dev1",
			dstats: map[string][]metrics.DStats{
				"/": {
					{
						Partition: &dext.PartitionStat{
							Device:     "/",
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
			},
			wantedData: rpi.Disk{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful"),
		},
		{
			name: "error: input device does not exist",
			dev:  "dev666",
			dstats: map[string][]metrics.DStats{
				"/dev1": {
					{
						Partition: &dext.PartitionStat{
							Device:     "/",
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
			},
			wantedData: rpi.Disk{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "dev666 does not exist"),
		},
		{
			name: "success: one device containing one mount point",
			dev:  "dev1",
			dstats: map[string][]metrics.DStats{
				"/dev1": {
					{
						Partition: &dext.PartitionStat{
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
			name: "success: two devices, each containing one mount point",
			dev:  "dev2",
			dstats: map[string][]metrics.DStats{
				"/dev1": {
					{
						Partition: &dext.PartitionStat{
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
				"/dev2": {
					{
						Partition: &dext.PartitionStat{
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
		{
			name: "success: multiple devices containing multiple mount points",
			dev:  "dev1",
			dstats: map[string][]metrics.DStats{
				"/dev1": {
					{
						Partition: &dext.PartitionStat{
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
						Partition: &dext.PartitionStat{
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
						Partition: &dext.PartitionStat{
							Device:     "/dev2",
							Mountpoint: "/dev2/mp21",
							Fstype:     "fs21",
							Opts:       "rw21",
						},
						Mountpoint: &dext.UsageStat{
							Path:              "/dev2/mp21",
							Fstype:            "fs21",
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
			wantedData: rpi.Disk{
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
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := disk.DSYS(Disk{})
			disk, err := s.View(tc.dev, tc.dstats)
			assert.Equal(t, tc.wantedData, disk)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
