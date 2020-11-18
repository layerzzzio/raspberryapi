package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/filestructure"
	"github.com/raspibuddy/rpi/pkg/api/metrics/filestructure/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"gopkg.in/go-playground/assert.v1"
)

func TestViewLF(t *testing.T) {
	var response rpi.FileStructure

	cases := []struct {
		name         string
		req          string
		fssys        *mocksys.FileStructure
		wantedStatus int
		wantedResp   rpi.FileStructure
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusNotFound,
			req:          "",
		},
		{
			name: "error: ViewLF directorypath not found",
			req:  "?directorypath=/dummy",
			fssys: &mocksys.FileStructure{
				ViewLFFn: func(*rpi.File, map[int64]string) (rpi.FileStructure, error) {
					return rpi.FileStructure{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ViewLF directorypath found",
			req:  "?directorypath=../transport",
			fssys: &mocksys.FileStructure{
				ViewLFFn: func(*rpi.File, map[int64]string) (rpi.FileStructure, error) {
					return rpi.FileStructure{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ViewLF directorypath found & filelimit not float",
			req:  "?directorypath=../transport&filelimit=0.1X",
			fssys: &mocksys.FileStructure{
				ViewLFFn: func(*rpi.File, map[int64]string) (rpi.FileStructure, error) {
					return rpi.FileStructure{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ViewLF directorypath, filelimit found & pathsize not int ",
			req:  "?directorypath=../transport&filelimit=1&pathsize=ABC",
			fssys: &mocksys.FileStructure{
				ViewLFFn: func(*rpi.File, map[int64]string) (rpi.FileStructure, error) {
					return rpi.FileStructure{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ViewLF return nil",
			req:  "?directorypath=../transport&filelimit=1&pathsize=1000",
			fssys: &mocksys.FileStructure{
				ViewLFFn: func(*rpi.File, map[int64]string) (rpi.FileStructure, error) {
					return rpi.FileStructure{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success: ViewLF directorypath, filelimit & pathsize found",
			req:  "?directorypath=../transport&filelimit=1&pathsize=1000",
			fssys: &mocksys.FileStructure{
				ViewLFFn: func(*rpi.File, map[int64]string) (rpi.FileStructure, error) {
					return rpi.FileStructure{
						DirectoryPath: "../transport",
						LargestFiles: []*rpi.File{
							{
								Path: "/boot/start.elf",
								Size: 3002560,
							},
							{
								Path: "/boot/start4x.elf",
								Size: 3038152,
							},
						},
					}, nil
				},
			},
			wantedResp: rpi.FileStructure{
				DirectoryPath: "../transport",
				LargestFiles: []*rpi.File{
					{
						Path: "/boot/start.elf",
						Size: 3002560,
					},
					{
						Path: "/boot/start4x.elf",
						Size: 3038152,
					},
				},
			},
			wantedStatus: http.StatusOK,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := filestructure.New(tc.fssys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/filestructure/largestfiles" + tc.req
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.DirectoryPath != "" {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
