package port_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/port"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		port       int32
		infos      mock.Infos
		psys       mocksys.Port
		wantedData rpi.Port
		wantedErr  error
	}{
		{
			name: "success: port listen",
			port: 6666,
			infos: mock.Infos{
				IsPortListeningFn: func(int32) bool {
					return true
				},
			},
			psys: mocksys.Port{
				ViewFn: func(
					bool,
				) (rpi.Port, error) {
					return rpi.Port{
						IsSpecificPortListen: true,
					}, nil
				},
			},
			wantedData: rpi.Port{
				IsSpecificPortListen: true,
			},
			wantedErr: nil,
		},
		{
			name: "success: port doesn't listen",
			port: 6666,
			infos: mock.Infos{
				IsPortListeningFn: func(int32) bool {
					return false
				},
			},
			psys: mocksys.Port{
				ViewFn: func(
					bool,
				) (rpi.Port, error) {
					return rpi.Port{
						IsSpecificPortListen: false,
					}, nil
				},
			},
			wantedData: rpi.Port{
				IsSpecificPortListen: false,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := port.New(&tc.psys, tc.infos)
			intf, err := s.View(tc.port)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
