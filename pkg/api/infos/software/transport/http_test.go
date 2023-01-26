package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
	"github.com/raspibuddy/rpi/pkg/api/infos/software/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response rpi.Software

	cases := []struct {
		name         string
		intsys       *mocksys.Software
		wantedStatus int
		wantedResp   rpi.Software
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			intsys: &mocksys.Software{
				ListFn: func(
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
				) (rpi.Software, error) {
					return rpi.Software{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			intsys: &mocksys.Software{
				ListFn: func(
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
				) (rpi.Software, error) {
					return rpi.Software{
						IsVNCInstalled:     true,
						IsOpenVPNInstalled: true,
						IsUnzipInstalled:   true,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Software{
				IsVNCInstalled:     true,
				IsOpenVPNInstalled: true,
				IsUnzipInstalled:   true,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := software.New(tc.intsys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/softwares"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.IsVNCInstalled == true || tc.wantedResp.IsVNCInstalled == false {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}

func TestView(t *testing.T) {
	var response rpi.Software

	cases := []struct {
		name         string
		req          string
		intsys       *mocksys.Software
		wantedStatus int
		wantedResp   rpi.Software
	}{
		{
			name:         "error: no package",
			req:          "?pkg=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: View result is nil",
			req:  "?pkg=XXXXXX",
			intsys: &mocksys.Software{
				ViewFn: func(bool) (rpi.Software, error) {
					return rpi.Software{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  "?pkg=good_pkg",
			intsys: &mocksys.Software{
				ViewFn: func(
					bool,
				) (rpi.Software, error) {
					return rpi.Software{
						IsSpecificSoftwareInstalled: true,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Software{
				IsSpecificSoftwareInstalled: true,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := software.New(tc.intsys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/softwares/specifics" + tc.req
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.IsSpecificSoftwareInstalled == true || tc.wantedResp.IsSpecificSoftwareInstalled == false {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
