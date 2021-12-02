package transport_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/appaction"
	"github.com/raspibuddy/rpi/pkg/api/actions/appaction/transport"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestExecuteWOVA(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		aacsys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: no action type",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action",
			req:          "?action=connectxxx",
			wantedStatus: http.StatusNotFound,
		},
		// {
		// 	name:         "error: no vpn name",
		// 	req:          "?action=connect&vpnName=",
		// 	wantedStatus: http.StatusNotFound,
		// },
		// {
		// 	name:         "error connect: no country",
		// 	req:          "?action=connect&vpnName=surfshark&country=&username=loic&password=abcd",
		// 	wantedStatus: http.StatusNotFound,
		// },
		// {
		// 	name:         "error connect: no username",
		// 	req:          "?action=connect&vpnName=surfshark&country=france&username=&password=abcd",
		// 	wantedStatus: http.StatusNotFound,
		// },
		// {
		// 	name:         "error connect: no password",
		// 	req:          "?action=connect&vpnName=surfshark&country=france&username=loic&password=",
		// 	wantedStatus: http.StatusNotFound,
		// },
		// {
		// 	name: "error: ExecuteWOVA result is nil",
		// 	req:  "?action=disconnect&vpnName=surfshark&country=france&username=loic&password=abcd",
		// 	aacsys: &mocksys.Action{
		// 		ExecuteWOVAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
		// 			return rpi.Action{}, errors.New("test error")
		// 		},
		// 	},
		// 	wantedStatus: http.StatusInternalServerError,
		// },
		// {
		// 	name:         "success",
		// 	wantedStatus: http.StatusOK,
		// 	req:          "?action=disconnect&vpnName=surfshark&country=france&username=loic&password=abcd",
		// 	aacsys: &mocksys.Action{
		// 		ExecuteWOVAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
		// 			return rpi.Action{
		// 				Name:          actions.ActionVPNWithOVPN,
		// 				NumberOfSteps: 1,
		// 				StartTime:     uint64(time.Now().Unix()),
		// 				EndTime:       uint64(time.Now().Unix()),
		// 				ExitStatus:    0,
		// 			}, nil
		// 		},
		// 	},
		// },
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			a := actions.New()
			i := infos.New()
			s := appaction.New(tc.aacsys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/appaction/vpnwithovpn" + tc.req

			res, err := http.Post(path, "application/json", bytes.NewBufferString(tc.req))
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.Name != "" {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
