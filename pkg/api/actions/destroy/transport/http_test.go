package transport_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/raspibuddy/rpi"
// 	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
// 	"github.com/raspibuddy/rpi/pkg/api/actions/destroy/transport"
// 	"github.com/raspibuddy/rpi/pkg/utl/actions"
// 	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
// 	"github.com/raspibuddy/rpi/pkg/utl/server"
// 	"github.com/stretchr/testify/assert"
// )

// func TestExecuteDF(t *testing.T) {
// 	var response rpi.Action

// 	cases := []struct {
// 		name         string
// 		req          string
// 		dessys       *mocksys.Action
// 		wantedStatus int
// 		wantedResp   rpi.Action
// 	}{
// 		{
// 			name:         "error: invalid request response",
// 			req:          "",
// 			wantedStatus: http.StatusNotFound,
// 		},
// 		{
// 			name: "error: ExecuteDF result is nil",
// 			req:  "?filepath=/dummy",
// 			dessys: &mocksys.Action{
// 				ExecuteDFFn: func(map[int]rpi.Exec) (rpi.Action, error) {
// 					return rpi.Action{}, errors.New("test error")
// 				},
// 			},
// 			wantedStatus: http.StatusInternalServerError,
// 		},
// 		{
// 			name:         "success",
// 			wantedStatus: http.StatusOK,
// 			req:          "?filepath=/dummy",
// 			dessys: &mocksys.Action{
// 				ExecuteDFFn: func(map[int]rpi.Exec) (rpi.Action, error) {
// 					return rpi.Action{
// 						Name: actions.DeleteFile,
// 						Steps: map[int]string{
// 							1: actions.DeleteFile,
// 						},
// 						NumberOfSteps: 1,
// 						StartTime:     uint64(time.Now().Unix()),
// 						EndTime:       uint64(time.Now().Unix()),
// 						ExitStatus:    0,
// 						Executions: map[int]rpi.Exec{
// 							1: {
// 								Name:       actions.DeleteFile,
// 								StartTime:  uint64(time.Now().Unix()),
// 								EndTime:    uint64(time.Now().Unix()),
// 								ExitStatus: 0,
// 								Stdin:      "",
// 								Stdout:     "",
// 								Stderr:     "",
// 							},
// 						},
// 					}, nil
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			r := server.New()
// 			rg := r.Group("")
// 			a := actions.New()
// 			s := destroy.New(tc.dessys, a)
// 			transport.NewHTTP(s, rg)
// 			ts := httptest.NewServer(r)

// 			defer ts.Close()
// 			path := ts.URL + "/destroy/deletefile" + tc.req

// 			fmt.Println(path)

// 			res, err := http.Post(path, "application/json", bytes.NewBufferString(tc.req))
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			if err != nil {
// 				panic(err)
// 			}

// 			if tc.wantedResp.Name != "" {
// 				if err := json.Unmarshal(body, &response); err != nil {
// 					t.Fatal(err)
// 				}
// 				assert.Equal(t, tc.wantedResp, response)
// 			}
// 			assert.Equal(t, tc.wantedStatus, res.StatusCode)
// 		})
// 	}
// }

// func TestExecuteSUS(t *testing.T) {
// 	var response rpi.Action

// 	cases := []struct {
// 		name         string
// 		req          string
// 		dessys       *mocksys.Action
// 		wantedStatus int
// 		wantedResp   rpi.Action
// 	}{
// 		{
// 			name:         "error: invalid request response",
// 			req:          "",
// 			wantedStatus: http.StatusNotFound,
// 		},
// 		{
// 			name: "error: ExecuteDF result is nil",
// 			req:  "?processname=dummyprocess",
// 			dessys: &mocksys.Action{
// 				ExecuteSUSFn: func(map[int]rpi.Exec) (rpi.Action, error) {
// 					return rpi.Action{}, errors.New("test error")
// 				},
// 			},
// 			wantedStatus: http.StatusInternalServerError,
// 		},
// 		{
// 			name:         "success",
// 			wantedStatus: http.StatusOK,
// 			req:          "?processname=dummyprocess",
// 			dessys: &mocksys.Action{
// 				ExecuteSUSFn: func(map[int]rpi.Exec) (rpi.Action, error) {
// 					return rpi.Action{
// 						Name: actions.StopUserSession,
// 						Steps: map[int]string{
// 							1: actions.KillProcessByName,
// 						},
// 						NumberOfSteps: 1,
// 						StartTime:     uint64(time.Now().Unix()),
// 						EndTime:       uint64(time.Now().Unix()),
// 						ExitStatus:    0,
// 						Executions: map[int]rpi.Exec{
// 							1: {
// 								Name:       actions.KillProcessByName,
// 								StartTime:  uint64(time.Now().Unix()),
// 								EndTime:    uint64(time.Now().Unix()),
// 								ExitStatus: 0,
// 								Stdin:      "",
// 								Stdout:     "",
// 								Stderr:     "",
// 							},
// 						},
// 					}, nil
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			r := server.New()
// 			rg := r.Group("")
// 			a := actions.New()
// 			s := destroy.New(tc.dessys, a)
// 			transport.NewHTTP(s, rg)
// 			ts := httptest.NewServer(r)

// 			defer ts.Close()
// 			path := ts.URL + "/destroy/stopusersession" + tc.req

// 			fmt.Println(path)

// 			res, err := http.Post(path, "application/json", bytes.NewBufferString(tc.req))
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			if err != nil {
// 				panic(err)
// 			}

// 			if tc.wantedResp.Name != "" {
// 				if err := json.Unmarshal(body, &response); err != nil {
// 					t.Fatal(err)
// 				}
// 				assert.Equal(t, tc.wantedResp, response)
// 			}
// 			assert.Equal(t, tc.wantedStatus, res.StatusCode)
// 		})
// 	}
// }

// func TestExecuteKP(t *testing.T) {
// 	var response rpi.Action

// 	cases := []struct {
// 		name         string
// 		req          string
// 		dessys       *mocksys.Action
// 		wantedStatus int
// 		wantedResp   rpi.Action
// 	}{
// 		{
// 			name:         "error: invalid request response",
// 			req:          "",
// 			wantedStatus: http.StatusNotFound,
// 		},
// 		{
// 			name:         "error: invalid pid",
// 			req:          "1A2B3C4",
// 			wantedStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name: "error: ExecuteKP result is nil",
// 			req:  "1234",
// 			dessys: &mocksys.Action{
// 				ExecuteKPFn: func(map[int]rpi.Exec) (rpi.Action, error) {
// 					return rpi.Action{}, errors.New("test error")
// 				},
// 			},
// 			wantedStatus: http.StatusInternalServerError,
// 		},
// 		{
// 			name: "success",
// 			req:  "1234",
// 			dessys: &mocksys.Action{
// 				ExecuteKPFn: func(map[int]rpi.Exec) (rpi.Action, error) {
// 					return rpi.Action{
// 						Name: actions.KillProcess,
// 						Steps: map[int]string{
// 							1: actions.KillProcess,
// 						},
// 						NumberOfSteps: 1,
// 						StartTime:     uint64(time.Now().Unix()),
// 						EndTime:       uint64(time.Now().Unix()),
// 						ExitStatus:    0,
// 						Executions: map[int]rpi.Exec{
// 							1: {
// 								Name:       actions.KillProcess,
// 								StartTime:  uint64(time.Now().Unix()),
// 								EndTime:    uint64(time.Now().Unix()),
// 								ExitStatus: 0,
// 								Stdin:      "",
// 								Stdout:     "",
// 								Stderr:     "",
// 							},
// 						},
// 					}, nil
// 				},
// 			},
// 			wantedResp: rpi.Action{
// 				Name: actions.KillProcess,
// 				Steps: map[int]string{
// 					1: actions.KillProcess,
// 				},
// 				NumberOfSteps: 1,
// 				StartTime:     uint64(time.Now().Unix()),
// 				EndTime:       uint64(time.Now().Unix()),
// 				ExitStatus:    0,
// 				Executions: map[int]rpi.Exec{
// 					1: {
// 						Name:       actions.KillProcess,
// 						StartTime:  uint64(time.Now().Unix()),
// 						EndTime:    uint64(time.Now().Unix()),
// 						ExitStatus: 0,
// 						Stdin:      "",
// 						Stdout:     "",
// 						Stderr:     "",
// 					},
// 				},
// 			},
// 			wantedStatus: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			r := server.New()
// 			rg := r.Group("")
// 			a := actions.New()
// 			s := destroy.New(tc.dessys, a)
// 			transport.NewHTTP(s, rg)
// 			ts := httptest.NewServer(r)

// 			defer ts.Close()
// 			path := ts.URL + "/destroy/killprocess/" + tc.req

// 			res, err := http.Post(path, "application/json", bytes.NewBufferString(tc.req))
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			if err != nil {
// 				panic(err)
// 			}

// 			if tc.wantedResp.Name != "" {
// 				if err := json.Unmarshal(body, &response); err != nil {
// 					t.Fatal(err)
// 				}
// 				assert.Equal(t, tc.wantedResp, response)
// 			}
// 			assert.Equal(t, tc.wantedStatus, res.StatusCode)
// 		})
// 	}
// }
