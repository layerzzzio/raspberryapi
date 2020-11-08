package mem_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/mem"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	mext "github.com/shirou/gopsutil/mem"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    *mock.Metrics
		msys       *mocksys.Mem
		wantedData rpi.Mem
		wantedErr  error
	}{
		{
			name: "error: swapMem & virtualMem arrays are nil",
			metrics: &mock.Metrics{
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{}, errors.New("test error info")
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{}, errors.New("test error info")
				},
			},
			wantedData: rpi.Mem{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the mem metrics"),
		},
		{
			name: "error: swapMem array is nil",
			metrics: &mock.Metrics{
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{}, errors.New("test error info")
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{
						Total: 999,
					}, nil
				},
			},
			wantedData: rpi.Mem{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the mem metrics"),
		},
		{
			name: "error: virtualMem array is nil",
			metrics: &mock.Metrics{
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{
						Total: 999,
					}, errors.New("test error info")
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{}, errors.New("test error info")
				},
			},
			wantedData: rpi.Mem{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the mem metrics"),
		},
		{
			name: "success",
			metrics: &mock.Metrics{
				SwapMemFn: func() (mext.SwapMemoryStat, error) {
					return mext.SwapMemoryStat{
						Total:       333,
						Used:        222,
						Free:        111,
						UsedPercent: 66.6,
					}, nil
				},
				VirtualMemFn: func() (mext.VirtualMemoryStat, error) {
					return mext.VirtualMemoryStat{
						Total:       333,
						Used:        222,
						Available:   111,
						UsedPercent: 66.6,
					}, nil
				},
			},
			msys: &mocksys.Mem{
				ListFn: func(mext.SwapMemoryStat, mext.VirtualMemoryStat) (rpi.Mem, error) {
					return rpi.Mem{
						STotal:       333,
						SUsed:        222,
						SFree:        111,
						SUsedPercent: 66.6,
						VTotal:       333,
						VUsed:        222,
						VAvailable:   111,
						VUsedPercent: 66.6,
					}, nil
				},
			},
			wantedData: rpi.Mem{
				STotal:       333,
				SUsed:        222,
				SFree:        111,
				SUsedPercent: 66.6,
				VTotal:       333,
				VUsed:        222,
				VAvailable:   111,
				VUsedPercent: 66.6,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := mem.New(tc.msys, tc.metrics)
			mems, err := s.List()
			assert.Equal(t, tc.wantedData, mems)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
