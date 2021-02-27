package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/display"
	"github.com/raspibuddy/rpi/pkg/api/infos/display/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response rpi.Display

	cases := []struct {
		name         string
		dissys       *mocksys.Display
		wantedStatus int
		wantedResp   rpi.Display
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			dissys: &mocksys.Display{
				ListFn: func([]string) (rpi.Display, error) {
					return rpi.Display{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		// {
		// 	name: "success",
		// 	dissys: &mocksys.Display{
		// 		ListFn: func([]string) (rpi.Display, error) {
		// 			return rpi.Display{
		// 				IsOverscan: true,
		// 			}, nil
		// 		},
		// 	},
		// 	wantedStatus: http.StatusOK,
		// 	wantedResp: rpi.Display{
		// 		IsOverscan: true,
		// 	},
		// },
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := display.New(tc.dissys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/displays"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if (tc.wantedResp != rpi.Display{}) {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
