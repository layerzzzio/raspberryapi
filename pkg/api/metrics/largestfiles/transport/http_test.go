package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfiles"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfiles/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response []rpi.LargestFiles
	cases := []struct {
		name         string
		lfsys        *mocksys.LargestFiles
		wantedStatus int
		wantedResp   []rpi.LargestFiles
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			lfsys: &mocksys.LargestFiles{
				ListFn: func([]metrics.PathSize) ([]rpi.LargestFiles, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			lfsys: &mocksys.LargestFiles{
				ListFn: func([]metrics.PathSize) ([]rpi.LargestFiles, error) {
					return []rpi.LargestFiles{
						{
							Path:                "/bin/file1",
							Size:                11,
							Category:            "/bin",
							CategoryDescription: "represents some essential user command binaries",
						},
						{
							Path:                "/usr/include/file2",
							Size:                22,
							Category:            "/usr/include",
							CategoryDescription: "contains system general-use include files for the C programming language",
						},
						{
							Path:                "/usr/dummy/file3",
							Size:                33,
							Category:            "/usr",
							CategoryDescription: "contains shareable and read-only data",
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: []rpi.LargestFiles{
				{
					Size:                11,
					Path:                "/bin/file1",
					Category:            "/bin",
					CategoryDescription: "represents some essential user command binaries",
				},
				{
					Size:                22,
					Path:                "/usr/include/file2",
					Category:            "/usr/include",
					CategoryDescription: "contains system general-use include files for the C programming language",
				},
				{
					Size:                33,
					Path:                "/usr/dummy/file3",
					Category:            "/usr",
					CategoryDescription: "contains shareable and read-only data",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := largestfiles.New(tc.lfsys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/largestfiles"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp != nil {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
