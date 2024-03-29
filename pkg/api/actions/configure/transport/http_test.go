package transport_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure/transport"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestExecuteCH(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: hostname badly formatted",
			req:          "?hostname=jkfd@jkfdkd.com",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: hostname nil",
			req:          "?hostname=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteCH result is nil",
			req:  "?hostname=new-hostname",
			consys: &mocksys.Action{
				ExecuteCHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?hostname=new-hostname",
			consys: &mocksys.Action{
				ExecuteCHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.ChangeHostname,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/changehostname" + tc.req

			fmt.Println(path)

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

func TestExecuteCP(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no password)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: no username",
			req:          "?password=new_password&username",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteCP result is nil",
			req:  "?password=new_password&username=new_username",
			consys: &mocksys.Action{
				ExecuteCPFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?password=new_password&username=new_username",
			consys: &mocksys.Action{
				ExecuteCPFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.ChangePassword,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/changepassword" + tc.req

			fmt.Println(path)

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

func TestActionCheck(t *testing.T) {
	cases := []struct {
		name       string
		action     string
		wantedResp error
	}{
		{
			name:       "error: bad action type",
			action:     "dummyaction",
			wantedResp: echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type or action type is null"),
		},
		{
			name:       "error: null action type",
			action:     "dummyaction",
			wantedResp: echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type or action type is null"),
		},
		{
			name:       "error: beginning ok, end nok",
			action:     "disable-xxx",
			wantedResp: echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type or action type is null"),
		},
		{
			name:       "error: beginning nok, end ok",
			action:     "xxx-disable",
			wantedResp: echo.NewHTTPError(http.StatusNotFound, "Not found - bad action type or action type is null"),
		},
		{
			name:       "success: enable",
			action:     "enable",
			wantedResp: nil,
		},
		{
			name:       "success: disable",
			action:     "disable",
			wantedResp: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := transport.ActionCheck(tc.action, `enable|disable`)
			assert.Equal(t, tc.wantedResp, res)
		})
	}
}

func TestExecuteWNB(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteWNB result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteWNBFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=enable",
			consys: &mocksys.Action{
				ExecuteWNBFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.WaitForNetworkAtBoot,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/waitfornetworkatboot" + tc.req

			fmt.Println(path)

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

func TestExecuteOV(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteOV result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteOVFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=enable",
			consys: &mocksys.Action{
				ExecuteOVFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.Overscan,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/overscan" + tc.req

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

func TestExecuteBL(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteBL result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteBLFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=enable",
			consys: &mocksys.Action{
				ExecuteBLFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.Blanking,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/blanking" + tc.req

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

func TestExecuteAUS(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no password)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: no username",
			req:          "?username=&password=new_password",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: no password",
			req:          "?username=username&password=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteAUS result is nil",
			req:  "?password=new_password&username=username",
			consys: &mocksys.Action{
				ExecuteAUSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?username=username&password=new_password",
			consys: &mocksys.Action{
				ExecuteAUSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.AddUser,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/adduser" + tc.req

			fmt.Println(path)

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

func TestExecuteDUS(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no password)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: no username",
			req:          "?username=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteDUS result is nil",
			req:  "?username=username",
			consys: &mocksys.Action{
				ExecuteDUSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?username=username",
			consys: &mocksys.Action{
				ExecuteDUSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.DeleteUser,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/deleteuser" + tc.req

			fmt.Println(path)

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

func TestExecuteCA(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteCA result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteCAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=enable",
			consys: &mocksys.Action{
				ExecuteCAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.Overscan,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/camera" + tc.req

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

func TestExecuteSSH(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteSSH result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteSSHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		// {
		// 	name:         "success",
		// 	wantedStatus: http.StatusOK,
		// 	req:          "?action=enable",
		// 	consys: &mocksys.Action{
		// 		ExecuteSSHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
		// 			return rpi.Action{
		// 				Name:          actions.SSH,
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
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/ssh" + tc.req

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

func TestExecuteVNC(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteVNC result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteVNCFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		// {
		// 	name:         "success",
		// 	wantedStatus: http.StatusOK,
		// 	req:          "?action=enable",
		// 	consys: &mocksys.Action{
		// 		ExecuteSSHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
		// 			return rpi.Action{
		// 				Name:          actions.SSH,
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
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/vnc" + tc.req

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

func TestExecuteSPI(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteSPI result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteSPIFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		// {
		// 	name:         "success",
		// 	wantedStatus: http.StatusOK,
		// 	req:          "?action=enable",
		// 	consys: &mocksys.Action{
		// 		ExecuteSSHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
		// 			return rpi.Action{
		// 				Name:          actions.SSH,
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
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/spi" + tc.req

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

func TestExecuteI2C(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteI2C result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteI2CFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		// {
		// 	name:         "success",
		// 	wantedStatus: http.StatusOK,
		// 	req:          "?action=enable",
		// 	consys: &mocksys.Action{
		// 		ExecuteSSHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
		// 			return rpi.Action{
		// 				Name:          actions.SSH,
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
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/i2c" + tc.req

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

func TestExecuteONW(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteONW result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteONWFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=enable",
			consys: &mocksys.Action{
				ExecuteONWFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.OneWire,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/onewire" + tc.req

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

func TestExecuteRG(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response (no action)",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: action enable is empty",
			req:          "?action=",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: bad action type",
			req:          "?action=dummyaction",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteRG result is nil",
			req:  "?action=enable",
			consys: &mocksys.Action{
				ExecuteRGFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?action=enable",
			consys: &mocksys.Action{
				ExecuteRGFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.RGPIO,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/rgpio" + tc.req

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

func TestExecuteUPD(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{

		{
			name: "error: ExecuteUPD result is nil",
			consys: &mocksys.Action{
				ExecuteUPDFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			consys: &mocksys.Action{
				ExecuteUPDFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.Update,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/update" + tc.req

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

func TestExecuteUPG(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{

		{
			name: "error: ExecuteUPG result is nil",
			consys: &mocksys.Action{
				ExecuteUPGFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			consys: &mocksys.Action{
				ExecuteUPGFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.Upgrade,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/upgrade" + tc.req

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

func TestExecuteUPDG(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{

		{
			name: "error: ExecuteUPDG result is nil",
			consys: &mocksys.Action{
				ExecuteUPDGFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			consys: &mocksys.Action{
				ExecuteUPDGFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.UpDateGrade,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/updateupgrade" + tc.req

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

func TestExecuteWC(t *testing.T) {
	var response rpi.Action

	cases := []struct {
		name         string
		req          string
		consys       *mocksys.Action
		wantedStatus int
		wantedResp   rpi.Action
	}{
		{
			name:         "error: invalid request response",
			req:          "",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: invalid request response (no iface)",
			req:          "?country=france&iface",
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "error: invalid request response (no country)",
			req:          "?country=&iface=wlan0",
			wantedStatus: http.StatusNotFound,
		},
		{
			name: "error: ExecuteWC result is nil",
			req:  "?country=france&iface=wlan0",
			consys: &mocksys.Action{
				ExecuteWCFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name:         "success",
			wantedStatus: http.StatusOK,
			req:          "?country=france&iface=wlan0",
			consys: &mocksys.Action{
				ExecuteWCFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.WifiCountry,
						NumberOfSteps: 1,
						StartTime:     uint64(time.Now().Unix()),
						EndTime:       uint64(time.Now().Unix()),
						ExitStatus:    0,
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
			i := infos.New()
			s := configure.New(tc.consys, a, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configure/wificountry" + tc.req

			fmt.Println(path)

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
