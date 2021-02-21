package display_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/display"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		dissys     mocksys.Display
		wantedData rpi.Display
		wantedErr  error
	}{

		{
			name: "error: readLines nil",
			infos: mock.Infos{
				ReadFileFn: func(string) ([]string, error) {
					return nil, errors.New("test error readLines")
				},
			},
			dissys: mocksys.Display{
				ListFn: func([]string) (rpi.Display, error) {
					return rpi.Display{}, nil
				},
			},
			wantedData: rpi.Display{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the display details"),
		},
		{
			name: "success",
			infos: mock.Infos{
				ReadFileFn: func(string) ([]string, error) {
					return []string{
						"line1",
						"disable_overscan=0",
						"line3",
					}, nil
				},
			},
			dissys: mocksys.Display{
				ListFn: func([]string) (rpi.Display, error) {
					return rpi.Display{IsOverscan: true}, nil
				},
			},
			wantedData: rpi.Display{IsOverscan: true},
			wantedErr:  nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := display.New(&tc.dissys, tc.infos)
			display, err := s.List()
			assert.Equal(t, tc.wantedData, display)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
