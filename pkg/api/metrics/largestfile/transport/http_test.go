package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfile"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfile/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	var response []rpi.LargestFile

	cases := []struct {
		name         string
		req          string
		lfsys        *mocksys.LargestFile
		wantedStatus int
		wantedResp   []rpi.LargestFile
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusNotFound,
			req:          "",
		},
		{
			name: "error: View result is nil",
			req:  "a",
			lfsys: &mocksys.LargestFile{
				ViewFn: func([]metrics.PathSize) ([]rpi.LargestFile, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		// {
		// 	name: "success",
		// 	req:  "_dummy_path",
		// 	lfsys: &mocksys.LargestFile{
		// 		ViewFn: func([]metrics.PathSize) ([]rpi.LargestFile, error) {
		// 			return []rpi.LargestFile{
		// 				{
		// 					Path:                "/bin/file1",
		// 					Name:                "file1",
		// 					Size:                11,
		// 					Category:            "/bin",
		// 					CategoryDescription: "represents some essential user command binaries",
		// 				},
		// 				{
		// 					Path:                "/usr/include/file2",
		// 					Name:                "file2",
		// 					Size:                22,
		// 					Category:            "/usr/include",
		// 					CategoryDescription: "contains system general-use include files for the C programming language",
		// 				},
		// 				{
		// 					Path:                "/usr/dummy/file3",
		// 					Name:                "file3",
		// 					Size:                33,
		// 					Category:            "/usr",
		// 					CategoryDescription: "contains shareable and read-only data",
		// 				},
		// 			}, nil
		// 		},
		// 	},
		// 	wantedStatus: http.StatusOK,
		// 	wantedResp: []rpi.LargestFile{
		// 		{
		// 			Size:                11,
		// 			Name:                "file1",
		// 			Path:                "/bin/file1",
		// 			Category:            "/bin",
		// 			CategoryDescription: "represents some essential user command binaries",
		// 		},
		// 		{
		// 			Size:                22,
		// 			Name:                "file2",
		// 			Path:                "/usr/include/file2",
		// 			Category:            "/usr/include",
		// 			CategoryDescription: "contains system general-use include files for the C programming language",
		// 		},
		// 		{
		// 			Size:                33,
		// 			Name:                "file3",
		// 			Path:                "/usr/dummy/file3",
		// 			Category:            "/usr",
		// 			CategoryDescription: "contains shareable and read-only data",
		// 		},
		// 	},
		// },
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := largestfile.New(tc.lfsys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/largestfiles/" + tc.req
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
