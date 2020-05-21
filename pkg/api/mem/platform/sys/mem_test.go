package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/mem"
	"github.com/raspibuddy/rpi/pkg/api/mem/platform/sys"
	mext "github.com/shirou/gopsutil/mem"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		swapMem    mext.SwapMemoryStat
		virtualMem mext.VirtualMemoryStat
		wantedData rpi.Mem
		wantedErr  error
	}{
		{
			name:       "success: virtualMem and swapMem are empty",
			swapMem:    mext.SwapMemoryStat{},
			virtualMem: mext.VirtualMemoryStat{},
			wantedData: rpi.Mem{
				STotal:       0,
				SUsed:        0,
				SFree:        0,
				SUsedPercent: 0,
				VTotal:       0,
				VUsed:        0,
				VAvailable:   0,
				VUsedPercent: 0,
			},
			wantedErr: nil,
		},
		{
			name:    "success: swapMem is empty",
			swapMem: mext.SwapMemoryStat{},
			virtualMem: mext.VirtualMemoryStat{
				Total:       333,
				Used:        222,
				Available:   111,
				UsedPercent: 66.6,
			},
			wantedData: rpi.Mem{
				STotal:       0,
				SUsed:        0,
				SFree:        0,
				SUsedPercent: 0,
				VTotal:       333,
				VUsed:        222,
				VAvailable:   111,
				VUsedPercent: 66.6,
			},
			wantedErr: nil,
		},
		{
			name: "success: virtualMem is empty",
			swapMem: mext.SwapMemoryStat{
				Total:       333,
				Used:        222,
				Free:        111,
				UsedPercent: 66.6,
			},
			virtualMem: mext.VirtualMemoryStat{
				Total:       0,
				Used:        0,
				Available:   0,
				UsedPercent: 0,
			},
			wantedData: rpi.Mem{
				STotal:       333,
				SUsed:        222,
				SFree:        111,
				SUsedPercent: 66.6,
				VTotal:       0,
				VUsed:        0,
				VAvailable:   0,
				VUsedPercent: 0,
			},
			wantedErr: nil,
		},
		{
			name: "success",
			swapMem: mext.SwapMemoryStat{
				Total:       333,
				Used:        222,
				Free:        111,
				UsedPercent: 66.6,
			},
			virtualMem: mext.VirtualMemoryStat{
				Total:       333,
				Used:        222,
				Available:   111,
				UsedPercent: 66.6,
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
			s := mem.MSYS(sys.Mem{})
			mems, err := s.List(tc.swapMem, tc.virtualMem)
			assert.Equal(t, tc.wantedData, mems)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
