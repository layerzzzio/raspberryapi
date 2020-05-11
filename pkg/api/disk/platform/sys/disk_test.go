package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/disk"
	"github.com/raspibuddy/rpi/pkg/api/disk/platform/sys"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	dext "github.com/shirou/gopsutil/disk"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		dstats     map[string][]metrics.DStats
		wantedData []rpi.Disk
		wantedErr  error
	}{
		{
			name: "success: multiple devices containing multiple mount points",
			dstats: map[string][]metrics.DStats{
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
			s := disk.DSYS(sys.Disk{})
			disk, err := s.List(tc.dstats)
			assert.Equal(t, tc.wantedData, disk)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
