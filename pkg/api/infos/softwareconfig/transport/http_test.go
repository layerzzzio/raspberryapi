package transport_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"gotest.tools/assert"
)

func TestList(t *testing.T) {
	var response rpi.SoftwareConfig

	cases := []struct {
		name         string
		intsys       *mocksys.SoftwareConfig
		wantedStatus int
		wantedResp   rpi.SoftwareConfig
	}{
		// {
		// 	name:         "error: invalid request response",
		// 	wantedStatus: http.StatusInternalServerError,
		// },
		// {
		// 	name: "error: List result is nil",
		// 	intsys: &mocksys.SoftwareConfig{
		// 		ListFn: func(
		// 			softwareconfig.NordVPN,
		// 		) (rpi.SoftwareConfig, error) {
		// 			return rpi.SoftwareConfig{}, errors.New("test error")
		// 		},
		// 	},
		// 	wantedStatus: http.StatusInternalServerError,
		// },
		// {
		// 	name: "success",
		// 	intsys: &mocksys.SoftwareConfig{
		// 		ListFn: func(
		// 			softwareconfig.NordVPN,
		// 		) (rpi.SoftwareConfig, error) {
		// 			return rpi.SoftwareConfig{
		// 				NordVPN: rpi.NordVPN{
		// 					TCPCountries: []string{"file"},
		// 				},
		// 			}, nil
		// 		},
		// 	},
		// 	wantedStatus: http.StatusOK,
		// 	wantedResp: rpi.SoftwareConfig{
		// 		NordVPN: rpi.NordVPN{
		// 			TCPCountries: []string{"file"},
		// 		},
		// 	},
		// },
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := softwareconfig.New(tc.intsys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/softwareconfigs"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.NordVPN.TCPCountries != nil {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
