package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/deletefile"
	"github.com/raspibuddy/rpi/pkg/api/actions/deletefile/transport"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		delsys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: result nil",
			req:  "_dummy",
			delsys: &mocksys.Action{
				ExecuteFn: func(map[int]rpi.Exec) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "_dummy",
			delsys: &mocksys.Action{
				ExecuteFn: func(map[int]rpi.Exec) (rpi.Action, error) {
					return rpi.Action{
						Name: actions.DeleteFile,
						Steps: map[int]string{
							1: actions.DeleteFile,
						},
						NumberOfSteps: 1,
						// StartTime:     uint64(time.Now().Unix()),
						// EndTime:       uint64(time.Now().Unix()),
						ExitStatus: 0,
						Executions: map[int]rpi.Exec{
							1: {
								Name: actions.DeleteFile,
								// StartTime:  uint64(time.Now().Unix()),
								// EndTime:    uint64(time.Now().Unix()),
								ExitStatus: 0,
								Stdin:      "",
								Stdout:     "",
								Stderr:     "",
							},
						},
					}, nil
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			a := actions.New()
			s := deletefile.New(tc.delsys, a)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/deletefile/" + tc.req
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
