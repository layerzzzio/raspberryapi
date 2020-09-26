package sys

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/host"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/shirou/gopsutil/cpu"
	dext "github.com/shirou/gopsutil/disk"
	hext "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"

	"github.com/stretchr/testify/assert"
)

func TestExtractTemp(t *testing.T) {
	cases := []struct {
		name       string
		input      string
		wantedData float32
	}{
		{
			name:       "error: no numerical characters",
			input:      "temp=AB.CD",
			wantedData: -1,
		},
		{
			name:       "success: dirty temp string",
			input:      "temp=20.CD",
			wantedData: 20.0,
		},
		{
			name:       "success: clean temp string",
			input:      "temp=20.9C",
			wantedData: 20.9,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			temp := extractTemp(tc.input)
			assert.Equal(t, tc.wantedData, temp)
		})
	}
}

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		info       hext.InfoStat
		users      []hext.UserStat
		cpus       []cpu.InfoStat
		vcores     []float64
		vMemPer    mem.VirtualMemoryStat
		sMemPer    mem.SwapMemoryStat
		load       load.AvgStat
		temp       string
		rpiv       string
		listDev    map[string][]metrics.DStats
		netInfo    []net.InterfaceStat
		wantedData rpi.Host
		wantedErr  error
	}{
		{
			name: "parsing disk id unsuccessful",
			netInfo: []net.InterfaceStat{
				{
					Index: 1,
					Name:  "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					Addrs: []net.InterfaceAddr{
						{
							Addr: "192.168.11.58",
						},
					},
				},
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			rpiv: "pi zero",
			listDev: map[string][]metrics.DStats{
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
			wantedData: rpi.Host{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful"),
		},
		{
			name: "multiple devices containing multiple mount points",
			info: hext.InfoStat{
				HostID:          "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:        "hostname_test",
				Uptime:          540165,
				BootTime:        1589223156,
				OS:              "raspbian",
				Procs:           400,
				Platform:        "plat_1",
				PlatformFamily:  "plat_1_1",
				PlatformVersion: "1.1",
				KernelArch:      "arch_A",
				KernelVersion:   "A",
			},
			cpus: []cpu.InfoStat{
				{
					CPU: 1,
				},
				{
					CPU: 2,
				},
			},
			vcores: []float64{1.0, 2.0, 3.0},
			vMemPer: mem.VirtualMemoryStat{
				UsedPercent: 99.9,
			},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			rpiv: "pi zero",
			listDev: map[string][]metrics.DStats{
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
			netInfo: []net.InterfaceStat{
				{
					Index: 1,
					Name:  "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					Addrs: []net.InterfaceAddr{
						{
							Addr: "192.168.11.58",
						},
					},
				},
			},
			wantedData: rpi.Host{
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
				CPU:                2,
				HyperThreading:     true,
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          400,
				ActiveVirtualUsers: 0,
				Temperature:        20.9,
				RaspModel:          "pi zero",
				Nets: []rpi.Net{
					{
						ID:   1,
						Name: "interface1",
						Flags: []string{
							"flag1",
							"flag2",
						},
						IPv4: "192.168.11.58",
					},
				},
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
			wantedErr: nil,
		},
		{
			name: "success: users array is nil",
			info: hext.InfoStat{
				HostID:          "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:        "hostname_test",
				Uptime:          540165,
				BootTime:        1589223156,
				OS:              "raspbian",
				Procs:           400,
				Platform:        "plat_1",
				PlatformFamily:  "plat_1_1",
				PlatformVersion: "1.1",
				KernelArch:      "arch_A",
				KernelVersion:   "A",
			},
			netInfo: []net.InterfaceStat{
				{
					Index: 1,
					Name:  "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					Addrs: []net.InterfaceAddr{
						{
							Addr: "192.168.11.58",
						},
					},
				},
			},
			cpus: []cpu.InfoStat{
				{
					CPU: 1,
				},
				{
					CPU: 2,
				},
			},
			vcores: []float64{1.0, 2.0, 3.0},
			vMemPer: mem.VirtualMemoryStat{
				UsedPercent: 99.9,
			},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			rpiv: "pi zero",
			wantedData: rpi.Host{
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
				CPU:                2,
				HyperThreading:     true,
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          400,
				ActiveVirtualUsers: 0,
				Temperature:        20.9,
				RaspModel:          "pi zero",
				Nets: []rpi.Net{
					{
						ID:   1,
						Name: "interface1",
						Flags: []string{
							"flag1",
							"flag2",
						},
						IPv4: "192.168.11.58",
					},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: cpus array is nil",
			info: hext.InfoStat{
				HostID:          "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:        "hostname_test",
				Uptime:          540165,
				BootTime:        1589223156,
				OS:              "raspbian",
				Procs:           400,
				Platform:        "plat_1",
				PlatformFamily:  "plat_1_1",
				PlatformVersion: "1.1",
				KernelArch:      "arch_A",
				KernelVersion:   "A",
			},
			users: []hext.UserStat{
				{
					User: "U1",
				},
				{
					User: "U2",
				},
			},
			vcores: []float64{1.0, 2.0, 3.0},
			vMemPer: mem.VirtualMemoryStat{
				UsedPercent: 99.9,
			},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			wantedData: rpi.Host{
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
				CPU:                0,
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
				Users: []rpi.User{
					{
						User: "U1",
					},
					{
						User: "U2",
					},
				},
				Temperature: 20.9,
			},
			wantedErr: nil,
		},
		{
			name: "success: vcores array is nil",
			users: []hext.UserStat{
				{
					User: "U1",
				},
				{
					User: "U2",
				},
			},
			info: hext.InfoStat{
				HostID:          "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:        "hostname_test",
				Uptime:          540165,
				BootTime:        1589223156,
				OS:              "raspbian",
				Procs:           400,
				Platform:        "plat_1",
				PlatformFamily:  "plat_1_1",
				PlatformVersion: "1.1",
				KernelArch:      "arch_A",
				KernelVersion:   "A",
			},
			cpus: []cpu.InfoStat{
				{
					CPU: 1,
				},
			},
			vMemPer: mem.VirtualMemoryStat{
				UsedPercent: 99.9,
			},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			wantedData: rpi.Host{
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
				VCore:              0,
				CPUUsedPercent:     0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          400,
				ActiveVirtualUsers: 2,
				Temperature:        20.9,
				Users: []rpi.User{
					{
						User: "U1",
					},
					{
						User: "U2",
					},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: vMemPer is nil",
			users: []hext.UserStat{
				{
					User: "U1",
				},
				{
					User: "U2",
				},
			},
			info: hext.InfoStat{
				HostID:          "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:        "hostname_test",
				Uptime:          540165,
				BootTime:        1589223156,
				OS:              "raspbian",
				Procs:           400,
				Platform:        "plat_1",
				PlatformFamily:  "plat_1_1",
				PlatformVersion: "1.1",
				KernelArch:      "arch_A",
				KernelVersion:   "A",
			},
			cpus: []cpu.InfoStat{
				{
					CPU: 1,
				},
			},
			vcores: []float64{1.0, 2.0, 3.0},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			wantedData: rpi.Host{
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
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       0,
				SUsedPercent:       0.9,
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          400,
				ActiveVirtualUsers: 2,
				Temperature:        20.9,
				Users: []rpi.User{
					{
						User: "U1",
					},
					{
						User: "U2",
					},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: sMemPer is nil",
			users: []hext.UserStat{
				{
					User: "U1",
				},
				{
					User: "U2",
				},
			},
			info: hext.InfoStat{
				HostID:          "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:        "hostname_test",
				Uptime:          540165,
				BootTime:        1589223156,
				OS:              "raspbian",
				Procs:           400,
				Platform:        "plat_1",
				PlatformFamily:  "plat_1_1",
				PlatformVersion: "1.1",
				KernelArch:      "arch_A",
				KernelVersion:   "A",
			},
			cpus: []cpu.InfoStat{
				{
					CPU: 1,
				},
			},
			vcores: []float64{1.0, 2.0, 3.0},
			vMemPer: mem.VirtualMemoryStat{
				UsedPercent: 99.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			wantedData: rpi.Host{
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
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0,
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          400,
				ActiveVirtualUsers: 2,
				Temperature:        20.9,
				Users: []rpi.User{
					{
						User: "U1",
					},
					{
						User: "U2",
					},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: infos is nil",
			users: []hext.UserStat{
				{
					User: "U1",
				},
				{
					User: "U2",
				},
			},
			cpus: []cpu.InfoStat{
				{
					CPU: 1,
				},
			},
			vcores: []float64{1.0, 2.0, 3.0},
			vMemPer: mem.VirtualMemoryStat{
				UsedPercent: 99.9,
			},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			wantedData: rpi.Host{
				ID:                 "",
				Hostname:           "",
				UpTime:             0,
				BootTime:           0,
				OS:                 "",
				Platform:           "",
				PlatformFamily:     "",
				PlatformVersion:    "",
				KernelArch:         "",
				KernelVersion:      "",
				CPU:                1,
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          0,
				ActiveVirtualUsers: 2,
				Temperature:        20.9,
				Users: []rpi.User{
					{
						User: "U1",
					},
					{
						User: "U2",
					},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: temperature is nil",
			users: []hext.UserStat{
				{
					User: "U1",
				},
				{
					User: "U2",
				},
			},
			info: hext.InfoStat{
				HostID:          "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:        "hostname_test",
				Uptime:          540165,
				BootTime:        1589223156,
				OS:              "raspbian",
				Procs:           400,
				Platform:        "plat_1",
				PlatformFamily:  "plat_1_1",
				PlatformVersion: "1.1",
				KernelArch:      "arch_A",
				KernelVersion:   "A",
			},
			cpus: []cpu.InfoStat{
				{
					CPU: 1,
				},
			},
			vcores: []float64{1.0, 2.0, 3.0},
			vMemPer: mem.VirtualMemoryStat{
				UsedPercent: 99.9,
				Total:       450,
			},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			wantedData: rpi.Host{
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
				VCore:              3,
				VTotal:             450,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          400,
				ActiveVirtualUsers: 2,
				Temperature:        -1,
				Users: []rpi.User{
					{
						User: "U1",
					},
					{
						User: "U2",
					},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success",
			users: []hext.UserStat{
				{
					User: "U1",
				},
				{
					User: "U2",
				},
			},
			info: hext.InfoStat{
				HostID:          "ab0aa7ee-3d03-3c21-91ad-5719d79d7af6",
				Hostname:        "hostname_test",
				Uptime:          540165,
				BootTime:        1589223156,
				OS:              "raspbian",
				Procs:           400,
				Platform:        "plat_1",
				PlatformFamily:  "plat_1_1",
				PlatformVersion: "1.1",
				KernelArch:      "arch_A",
				KernelVersion:   "A",
			},
			cpus: []cpu.InfoStat{
				{
					CPU: 1,
				},
			},
			vcores: []float64{1.0, 2.0, 3.0},
			vMemPer: mem.VirtualMemoryStat{
				UsedPercent: 99.9,
				Total:       450,
			},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			load: load.AvgStat{
				Load1:  1,
				Load5:  5,
				Load15: 15,
			},
			temp: "temp=20.9.C",
			wantedData: rpi.Host{
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
				VCore:              3,
				VTotal:             450,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Load1:              1,
				Load5:              5,
				Load15:             15,
				Processes:          400,
				ActiveVirtualUsers: 2,
				Temperature:        20.9,
				Users: []rpi.User{
					{
						User: "U1",
					},
					{
						User: "U2",
					},
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := host.HSYS(Host{})

			hosts, err := s.List(
				tc.info,
				tc.users,
				tc.cpus,
				tc.vcores,
				tc.vMemPer,
				tc.sMemPer,
				tc.load,
				tc.temp,
				tc.rpiv,
				tc.listDev,
				tc.netInfo)

			assert.Equal(t, tc.wantedData, hosts)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
