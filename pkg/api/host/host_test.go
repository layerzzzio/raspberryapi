package host_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/host"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/shirou/gopsutil/cpu"
	hext "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	mext "github.com/shirou/gopsutil/mem"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    *mock.Metrics
		hsys       *mocksys.Host
		wantedData rpi.Host
		wantedErr  error
	}{
		{
			name: "error: all arrays are nil",
			metrics: &mock.Metrics{
				HostInfoFn: func() (hext.InfoStat, error) {
					return hext.InfoStat{}, errors.New("test error info")
				},
				UsersFn: func() ([]hext.UserStat, error) {
					return nil, errors.New("test error info")
				},
				CPUInfoFn: func() ([]cpu.InfoStat, error) {
					return nil, errors.New("test error info")
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error info")
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{}, errors.New("test error info")
				},
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{}, errors.New("test error info")
				},
				TemperatureFn: func() (string, string, error) {
					return "", "", errors.New("test error info")
				},
			},
			wantedData: rpi.Host{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the host metrics"),
		},
		{
			name: "error: temperature stdErr and err are not nil",
			metrics: &mock.Metrics{
				HostInfoFn: func() (hext.InfoStat, error) {
					return hext.InfoStat{}, nil
				},
				UsersFn: func() ([]hext.UserStat, error) {
					return nil, nil
				},
				CPUInfoFn: func() ([]cpu.InfoStat, error) {
					return nil, nil
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, nil
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{}, nil
				},
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{}, nil
				},
				TemperatureFn: func() (string, string, error) {
					return "", "error", errors.New("test error info")
				},
			},
			wantedData: rpi.Host{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the host metrics"),
		},
		{
			name: "error: infos array is nil",
			metrics: &mock.Metrics{
				HostInfoFn: func() (hext.InfoStat, error) {
					return hext.InfoStat{}, errors.New("test error info")
				},
				UsersFn: func() ([]hext.UserStat, error) {
					return []hext.UserStat{
						{
							User: "test_user",
						},
					}, nil
				},
				CPUInfoFn: func() ([]cpu.InfoStat, error) {
					return []cpu.InfoStat{
						{
							CPU: 1,
						},
					}, nil
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{
						UsedPercent: 99.9,
					}, nil
				},
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{
						UsedPercent: 1.1,
					}, nil
				},
				TemperatureFn: func() (string, string, error) {
					return "", "", errors.New("test error info")
				},
			},
			wantedData: rpi.Host{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the host metrics"),
		},
		{
			name: "error: users array is nil",
			metrics: &mock.Metrics{
				HostInfoFn: func() (hext.InfoStat, error) {
					return hext.InfoStat{
						Hostname: "hostname_test",
					}, nil
				},
				UsersFn: func() ([]hext.UserStat, error) {
					return []hext.UserStat{}, errors.New("test error info")
				},
				CPUInfoFn: func() ([]cpu.InfoStat, error) {
					return []cpu.InfoStat{
						{
							CPU: 1,
						},
					}, nil
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{
						UsedPercent: 99.9,
					}, nil
				},
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{
						UsedPercent: 1.1,
					}, nil
				},
				TemperatureFn: func() (string, string, error) {
					return "", "", errors.New("test error info")
				},
			},
			wantedData: rpi.Host{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the host metrics"),
		},
		{
			name: "error: cpus is nil",
			metrics: &mock.Metrics{
				HostInfoFn: func() (hext.InfoStat, error) {
					return hext.InfoStat{
						Hostname: "hostname_test",
					}, nil
				},
				UsersFn: func() ([]hext.UserStat, error) {
					return []hext.UserStat{}, nil
				},
				CPUInfoFn: func() ([]cpu.InfoStat, error) {
					return nil, errors.New("test error info")
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{
						UsedPercent: 99.9,
					}, nil
				},
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{
						UsedPercent: 1.1,
					}, nil
				},
				TemperatureFn: func() (string, string, error) {
					return "", "", errors.New("test error info")
				},
			},
			wantedData: rpi.Host{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the host metrics"),
		},
		{
			name: "success",
			metrics: &mock.Metrics{
				HostInfoFn: func() (hext.InfoStat, error) {
					return hext.InfoStat{
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
					}, nil
				},
				UsersFn: func() ([]hext.UserStat, error) {
					return []hext.UserStat{
						{
							User: "U1",
						},
						{
							User: "U2",
						},
					}, nil
				},
				CPUInfoFn: func() ([]cpu.InfoStat, error) {
					return []cpu.InfoStat{
						{
							CPU: 1,
						},
					}, nil
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{1.0, 2.0, 3.0}, nil
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mem.VirtualMemoryStat{
						UsedPercent: 99.9,
					}, nil
				},
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mem.SwapMemoryStat{
						UsedPercent: 0.9,
					}, nil
				},
				TemperatureFn: func() (string, string, error) {
					return "", "", errors.New("test error info")
				},
			},
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
						VCore:              3,
						CPUUsedPercent:     2.0,
						VUsedPercent:       99.9,
						SUsedPercent:       0.9,
						Processes:          400,
						ActiveVirtualUsers: 2,
					}, nil
				},
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
			s := host.New(tc.hsys, tc.metrics)
			hosts, err := s.List()
			assert.Equal(t, tc.wantedData, hosts)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
