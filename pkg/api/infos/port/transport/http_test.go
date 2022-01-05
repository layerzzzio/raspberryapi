package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/port"
	"github.com/raspibuddy/rpi/pkg/api/infos/port/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	var response rpi.Port

	cases := []struct {
		name         string
		req          string
		psys         *mocksys.Port
		wantedStatus int
		wantedResp   rpi.Port
	}{
		{
			name:         "error: invalid request",
			wantedStatus: http.StatusBadRequest,
			req:          `a`,
		},
		// WARNING !!!!!
		// using an impossible port that will return false
		{
			name: "error: View result is nil",
			req:  "1",
			psys: &mocksys.Port{
				ViewFn: func(bool) (rpi.Port, error) {
					return rpi.Port{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  "666666666",
			psys: &mocksys.Port{
				ViewFn: func(bool) (rpi.Port, error) {
					return rpi.Port{
						IsSpecificPortListen: false,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Port{
				IsSpecificPortListen: false,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := port.New(tc.psys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/ports/" + tc.req
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if !tc.wantedResp.IsSpecificPortListen || tc.wantedResp.IsSpecificPortListen {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
