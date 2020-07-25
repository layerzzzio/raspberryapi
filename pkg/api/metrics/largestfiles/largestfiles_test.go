package largestfiles_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfiles"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    *mock.Metrics
		lfsys      mocksys.LargestFiles
		wantedData []rpi.LargestFiles
		wantedErr  error
	}{
		{
			name: "error: top100files array is nil",
			metrics: &mock.Metrics{
				Top100FilesFn: func() ([]metrics.PathSize, string, error) {
					return nil, "", errors.New("test error info")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the largest files"),
		},
		// {
		// 	name: "success",
		// 	metrics: &mock.Metrics{
		// 		Top100FilesFn: func() ([]metrics.PathSize, string, error) {
		// 			return []metrics.PathSize{
		// 				{
		// 					Path: "/bin/file1",
		// 					Size: 11,
		// 				},
		// 				{
		// 					Path: "/usr/include/file2",
		// 					Size: 22,
		// 				},
		// 				{
		// 					Path: "/usr/dummy/file3",
		// 					Size: 33,
		// 				},
		// 			}, "", nil
		// 		},
		// 	},
		// 	lfsys: mocksys.LargestFiles{
		// 		ListFn: func([]metrics.PathSize) ([]rpi.LargestFiles, error) {
		// 			return []rpi.LargestFiles{
		// 				{
		// 					Path:                "/bin/file1",
		// 					Size:                11,
		// 					Category:            "/bin",
		// 					CategoryDescription: "represents some essential user command binaries",
		// 				},
		// 				{
		// 					Path:                "/usr/include/file2",
		// 					Size:                22,
		// 					Category:            "/usr/include",
		// 					CategoryDescription: "contains system general-use include files for the C programming language",
		// 				},
		// 				{
		// 					Path:                "/usr/dummy/file3",
		// 					Size:                33,
		// 					Category:            "/usr",
		// 					CategoryDescription: "contains shareable and read-only data",
		// 				},
		// 			}, nil
		// 		},
		// 	},
		// 	wantedData: []rpi.LargestFiles{
		// 		{
		// 			Path:                "/bin/file1",
		// 			Size:                11,
		// 			Category:            "/bin",
		// 			CategoryDescription: "represents some essential user command binaries",
		// 		},
		// 		{
		// 			Path:                "/usr/include/file2",
		// 			Size:                22,
		// 			Category:            "/usr/include",
		// 			CategoryDescription: "contains system general-use include files for the C programming language",
		// 		},
		// 		{
		// 			Path:                "/usr/dummy/file3",
		// 			Size:                33,
		// 			Category:            "/usr",
		// 			CategoryDescription: "contains shareable and read-only data",
		// 		},
		// 	},
		// 	wantedErr: nil,
		// },
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := largestfiles.New(tc.lfsys, tc.metrics)
			users, err := s.List()
			assert.Equal(t, tc.wantedData, users)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
