package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/boot"
	"github.com/raspibuddy/rpi/pkg/api/infos/boot/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response rpi.Boot

	cases := []struct {
		name         string
		boosys       *mocksys.Boot
		wantedStatus int
		wantedResp   rpi.Boot
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			boosys: &mocksys.Boot{
				ListFn: func(bool) (rpi.Boot, error) {
					return rpi.Boot{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			boosys: &mocksys.Boot{
				ListFn: func(bool) (rpi.Boot, error) {
					return rpi.Boot{
						IsWaitForNetwork: true,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Boot{
				IsWaitForNetwork: true,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := boot.New(tc.boosys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/boots"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if (tc.wantedResp != rpi.Boot{}) {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
