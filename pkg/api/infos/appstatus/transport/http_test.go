package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/appstatus"
	"github.com/raspibuddy/rpi/pkg/api/infos/appstatus/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response rpi.AppStatus

	cases := []struct {
		name         string
		apssys       *mocksys.AppStatus
		wantedStatus int
		wantedResp   rpi.AppStatus
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			apssys: &mocksys.AppStatus{
				ListFn: func(map[string]bool) (rpi.AppStatus, error) {
					return rpi.AppStatus{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			apssys: &mocksys.AppStatus{
				ListFn: func(map[string]bool) (rpi.AppStatus, error) {
					return rpi.AppStatus{
						VPNwithOpenVPN: map[string]bool{
							"nordvpn": true,
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.AppStatus{
				VPNwithOpenVPN: map[string]bool{
					"nordvpn": true,
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := appstatus.New(tc.apssys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/appstatuses/vpnwithovpn"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.VPNwithOpenVPN != nil {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
