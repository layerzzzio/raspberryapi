package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/admin/deployment"
	"github.com/raspibuddy/rpi/pkg/api/admin/deployment/transport"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestExecuteDPTOOL(t *testing.T) {
	var response rpi.Version

	cases := []struct {
		name         string
		req          string
		dsys         *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: no url, no version",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: url but no version",
			req:          "?url=https//url.com",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: url but badly formatted version",
			req:          "?url=https//url.com&version=X.X.X",
			wantedStatus: http.StatusBadRequest,
		},
		{
			name: "error: ExecuteDF result is nil",
			req:  "?url=https//url.com&version=1.1.1",
			dsys: &mocksys.Action{
				ExecuteDPTOOLFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  "?url=https//url.com&version=1.1.1",
			dsys: &mocksys.Action{
				ExecuteDPTOOLFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.DeployVersion,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
					}, nil
				}},
			wantedStatus: http.StatusOK,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			a := actions.New()
			s := deployment.New(tc.dsys, a)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/deploy/version" + tc.req
			res, err := http.Get(path)
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
