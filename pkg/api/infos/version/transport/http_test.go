package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/version"
	"github.com/raspibuddy/rpi/pkg/api/infos/version/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestListAll(t *testing.T) {
	var response rpi.Version

	cases := []struct {
		name         string
		vsys         *mocksys.Version
		wantedStatus int
		wantedResp   rpi.Version
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			vsys: &mocksys.Version{
				ListAllFn: func(string, string) (rpi.Version, error) {
					return rpi.Version{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			vsys: &mocksys.Version{
				ListAllFn: func(string, string) (rpi.Version, error) {
					return rpi.Version{
						RaspiBuddyVersion:       "1.0.0",
						RaspiBuddyDeployVersion: "1.1.1",
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Version{
				RaspiBuddyVersion:       "1.0.0",
				RaspiBuddyDeployVersion: "1.1.1",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := version.New(tc.vsys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/versions"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.RaspiBuddyVersion != "" {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}

func TestListAllApis(t *testing.T) {
	var response rpi.Version

	cases := []struct {
		name         string
		vsys         *mocksys.Version
		wantedStatus int
		wantedResp   rpi.Version
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			vsys: &mocksys.Version{
				ListAllApisFn: func(string, string) (rpi.Version, error) {
					return rpi.Version{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			vsys: &mocksys.Version{
				ListAllApisFn: func(string, string) (rpi.Version, error) {
					return rpi.Version{
						RaspiBuddyVersion:       "1.0.0",
						RaspiBuddyDeployVersion: "1.1.1",
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Version{
				RaspiBuddyVersion:       "1.0.0",
				RaspiBuddyDeployVersion: "1.1.1",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := version.New(tc.vsys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/versions/apis"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.RaspiBuddyVersion != "" {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
