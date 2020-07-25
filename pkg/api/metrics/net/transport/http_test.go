package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/net"
	"github.com/raspibuddy/rpi/pkg/api/metrics/net/transport"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	next "github.com/shirou/gopsutil/net"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response []rpi.Net

	cases := []struct {
		name         string
		nsys         *mocksys.Net
		wantedStatus int
		wantedResp   []rpi.Net
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			nsys: &mocksys.Net{
				ListFn: func([]next.InterfaceStat) ([]rpi.Net, error) {
					return nil, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			nsys: &mocksys.Net{
				ListFn: func([]next.InterfaceStat) ([]rpi.Net, error) {
					return []rpi.Net{
						{
							ID:   1,
							Name: "interface1",
							Flags: []string{
								"flag1",
								"flag2",
							},
							IPv4: "192.168.11.58",
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: []rpi.Net{
				{
					ID:   1,
					Name: "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					IPv4: "192.168.11.58",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := net.New(tc.nsys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/nets"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp != nil {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}

func TestView(t *testing.T) {
	var response rpi.Net

	cases := []struct {
		name         string
		req          string
		nsys         *mocksys.Net
		wantedStatus int
		wantedResp   rpi.Net
	}{
		{
			name:         "error: invalid request",
			wantedStatus: http.StatusBadRequest,
			req:          `a`,
		},
		{
			name: "error: View result is nil",
			req:  "1",
			nsys: &mocksys.Net{
				ViewFn: func(int, []next.InterfaceStat, []next.IOCountersStat) (rpi.Net, error) {
					return rpi.Net{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  "1",
			nsys: &mocksys.Net{
				ViewFn: func(int, []next.InterfaceStat, []next.IOCountersStat) (rpi.Net, error) {
					return rpi.Net{
						ID:   1,
						Name: "interface1",
						Flags: []string{
							"flag1",
							"flag2",
						},
						IPv4:        "192.168.11.58",
						BytesSent:   1,
						BytesRecv:   2,
						PacketsSent: 3,
						PacketsRecv: 4,
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.Net{
				ID:   1,
				Name: "interface1",
				Flags: []string{
					"flag1",
					"flag2",
				},
				IPv4:        "192.168.11.58",
				BytesSent:   1,
				BytesRecv:   2,
				PacketsSent: 3,
				PacketsRecv: 4,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			m := metrics.New(metrics.Service{})
			s := net.New(tc.nsys, m)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/nets/" + tc.req
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
