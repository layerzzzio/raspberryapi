package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/host"
	"github.com/raspibuddy/rpi/pkg/api/host/platform/sys"
	"github.com/shirou/gopsutil/cpu"
	hext "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		info       hext.InfoStat
		users      []hext.UserStat
		cpus       []cpu.InfoStat
		vcores     []float64
		vMemPer    mem.VirtualMemoryStat
		sMemPer    mem.SwapMemoryStat
		wantedData rpi.Host
		wantedErr  error
	}{
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
			wantedData: rpi.Host{
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
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Processes:          400,
				ActiveVirtualUsers: 0,
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
			wantedData: rpi.Host{
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
				CPU:                0,
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Processes:          400,
				ActiveVirtualUsers: 2,
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
			wantedData: rpi.Host{
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
				VCore:              0,
				CPUUsedPercent:     0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Processes:          400,
				ActiveVirtualUsers: 2,
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
			wantedData: rpi.Host{
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
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       0,
				SUsedPercent:       0.9,
				Processes:          400,
				ActiveVirtualUsers: 2,
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
			wantedData: rpi.Host{
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
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0,
				Processes:          400,
				ActiveVirtualUsers: 2,
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
			wantedData: rpi.Host{
				ID:                 "",
				Hostname:           "",
				Uptime:             0,
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
				Processes:          0,
				ActiveVirtualUsers: 2,
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
			},
			sMemPer: mem.SwapMemoryStat{
				UsedPercent: 0.9,
			},
			wantedData: rpi.Host{
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
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Processes:          400,
				ActiveVirtualUsers: 2,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := host.HSYS(sys.Host{})

			hosts, err := s.List(
				tc.info,
				tc.users,
				tc.cpus,
				tc.vcores,
				tc.vMemPer,
				tc.sMemPer)

			assert.Equal(t, tc.wantedData, hosts)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
