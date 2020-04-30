package cpu_test

import (
	"testing"

	"github.com/raspibuddy/rpi/pkg/api/cpu"
	"github.com/raspibuddy/rpi/utl/mock/mocksys"
	cext "github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name            string
		wantDataInfo    []cext.InfoStat
		wantDataPercent []float64
		wantDataTime    []cext.TimesStat
		wantErr         bool
		csys            *mocksys.CPU
	}{
		// {
		// 	name:    "Failure list service",
		// 	wantErr: true,
		// },
		{
			name: "Success list service",
			csys: &mocksys.CPU{
				ListFn: func() ([]cext.InfoStat, []float64, []cext.TimesStat, error) {
					return []cext.InfoStat{
							{
								CPU:       0,
								ModelName: "test model",
								Cores:     16,
								Mhz:       2300.99,
							},
						},
						[]float64{99.9},
						[]cext.TimesStat{
							{
								CPU:    "total-cpu",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						}, nil
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := cpu.New(tc.csys)
			cpus, err := s.List()
			assert.Equal(t, cpus[0].ID, 0)
			assert.Equal(t, cpus[0].ModelName, "test model")
			assert.Equal(t, cpus[0].Cores, int32(16))
			assert.Equal(t, cpus[0].Mhz, 2300.99)
			assert.Equal(t, cpus[0].Stats.Used, 99.9)
			assert.Equal(t, cpus[0].Stats.User, 111.1)
			assert.Equal(t, cpus[0].Stats.System, 222.2)
			assert.Equal(t, cpus[0].Stats.Idle, 333.3)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
