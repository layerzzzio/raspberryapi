package net_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/net"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	next "github.com/shirou/gopsutil/net"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    *mock.Metrics
		nsys       *mocksys.Net
		wantedData []rpi.Net
		wantedErr  error
	}{
		{
			name: "error: netInfo array is nil",
			metrics: &mock.Metrics{
				NetInfoFn: func() ([]next.InterfaceStat, error) {
					return nil, errors.New("test error info")
				},
			},
			nsys: &mocksys.Net{
				ListFn: func([]next.InterfaceStat) ([]rpi.Net, error) {
					return nil, errors.New("test error info")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not list the net metrics"),
		},
		{
			name: "error: netInfo array is empty",
			metrics: &mock.Metrics{
				NetInfoFn: func() ([]next.InterfaceStat, error) {
					return []next.InterfaceStat{}, errors.New("test error info")
				},
			},
			nsys: &mocksys.Net{
				ListFn: func([]next.InterfaceStat) ([]rpi.Net, error) {
					return nil, errors.New("test error info")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not list the net metrics"),
		},
		{
			name: "success",
			metrics: &mock.Metrics{
				NetInfoFn: func() ([]next.InterfaceStat, error) {
					return []next.InterfaceStat{
						{
							Index: 1,
							Name:  "interface1",
							Flags: []string{
								"flag1",
								"flag2",
							},
							Addrs: []next.InterfaceAddr{
								{
									Addr: "192.2.1.2/28",
								},
							},
						},
					}, nil
				},
			},
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
							IPv4: "192.2.1.2",
						},
					}, nil
				},
			},
			wantedData: []rpi.Net{
				{
					ID:   1,
					Name: "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					IPv4: "192.2.1.2",
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := net.New(tc.nsys, tc.metrics)
			nets, err := s.List()
			assert.Equal(t, tc.wantedData, nets)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		id         int
		metrics    *mock.Metrics
		nsys       *mocksys.Net
		wantedData rpi.Net
		wantedErr  error
	}{
		{
			name: "error: netInfo and netStats array are nil",
			metrics: &mock.Metrics{
				NetInfoFn: func() ([]next.InterfaceStat, error) {
					return nil, errors.New("test error info")
				},
				NetStatsFn: func() ([]next.IOCountersStat, error) {
					return nil, errors.New("test error info")
				},
			},
			nsys: &mocksys.Net{
				ViewFn: func(int, []next.InterfaceStat, []next.IOCountersStat) (rpi.Net, error) {
					return rpi.Net{}, errors.New("test error info")
				},
			},
			wantedData: rpi.Net{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not view the net metrics"),
		},
		{
			name: "error: netInfo and netStats arrays are empty",
			metrics: &mock.Metrics{
				NetInfoFn: func() ([]next.InterfaceStat, error) {
					return []next.InterfaceStat{}, nil
				},
				NetStatsFn: func() ([]next.IOCountersStat, error) {
					return []next.IOCountersStat{}, nil
				},
			},
			nsys: &mocksys.Net{
				ViewFn: func(int, []next.InterfaceStat, []next.IOCountersStat) (rpi.Net, error) {
					return rpi.Net{}, nil
				},
			},
			wantedData: rpi.Net{},
			wantedErr:  nil,
		},
		{
			name: "error: netInfo array is empty",
			metrics: &mock.Metrics{
				NetInfoFn: func() ([]next.InterfaceStat, error) {
					return []next.InterfaceStat{}, nil
				},
				NetStatsFn: func() ([]next.IOCountersStat, error) {
					return []next.IOCountersStat{
						{
							Name:        "interface1",
							BytesSent:   1,
							BytesRecv:   2,
							PacketsSent: 3,
							PacketsRecv: 4,
						},
					}, nil
				},
			},
			nsys: &mocksys.Net{
				ViewFn: func(int, []next.InterfaceStat, []next.IOCountersStat) (rpi.Net, error) {
					return rpi.Net{}, errors.New("test error info")
				},
			},
			wantedData: rpi.Net{},
			wantedErr:  errors.New("test error info"),
		},
		{
			name: "error: netStats array is empty",
			metrics: &mock.Metrics{
				NetInfoFn: func() ([]next.InterfaceStat, error) {
					return []next.InterfaceStat{
						{
							Index: 1,
							Name:  "interface1",
							Flags: []string{
								"flag1",
								"flag2",
							},
							Addrs: []next.InterfaceAddr{
								{
									Addr: "192.2.1.2/28",
								},
							},
						},
					}, nil
				},
				NetStatsFn: func() ([]next.IOCountersStat, error) {
					return []next.IOCountersStat{}, nil
				},
			},
			nsys: &mocksys.Net{
				ViewFn: func(int, []next.InterfaceStat, []next.IOCountersStat) (rpi.Net, error) {
					return rpi.Net{
						ID:   1,
						Name: "interface1",
						Flags: []string{
							"flag1",
							"flag2",
						},
						IPv4: "192.2.1.2",
					}, nil
				},
			},
			wantedData: rpi.Net{
				ID:   1,
				Name: "interface1",
				Flags: []string{
					"flag1",
					"flag2",
				},
				IPv4: "192.2.1.2",
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := net.New(tc.nsys, tc.metrics)
			nets, err := s.View(tc.id)
			assert.Equal(t, tc.wantedData, nets)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
