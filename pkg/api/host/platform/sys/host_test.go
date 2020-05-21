package sys

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/host"
	"github.com/shirou/gopsutil/cpu"
	hext "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

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
		temp       string
		rpiv       string
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
			temp: "temp=20.9.C",
			rpiv: "pi zero",
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
				CPU:                2,
				HyperThreading:     true,
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Processes:          400,
				ActiveVirtualUsers: 0,
				Temperature:        20.9,
				RaspModel:          "pi zero",
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
			temp: "temp=20.9.C",
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
				HyperThreading:     false,
				VCore:              3,
				CPUUsedPercent:     2.0,
				VUsedPercent:       99.9,
				SUsedPercent:       0.9,
				Processes:          400,
				ActiveVirtualUsers: 2,
				Temperature:        20.9,
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
			temp: "temp=20.9.C",
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
				Temperature:        20.9,
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
			temp: "temp=20.9.C",
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
				Temperature:        20.9,
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
			temp: "temp=20.9.C",
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
				Temperature:        20.9,
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
			temp: "temp=20.9.C",
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
				Temperature:        20.9,
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
				Temperature:        -1,
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
			temp: "temp=20.9.C",
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
				Temperature:        20.9,
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
				tc.temp,
				tc.rpiv)

			assert.Equal(t, tc.wantedData, hosts)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
