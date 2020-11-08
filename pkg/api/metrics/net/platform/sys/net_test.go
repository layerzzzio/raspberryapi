package sys

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/net"
	next "github.com/shirou/gopsutil/net"

	"github.com/stretchr/testify/assert"
)

func TestExtractIPv4(t *testing.T) {
	cases := []struct {
		name       string
		addrs      []next.InterfaceAddr
		wantedData string
	}{
		{
			name:       "addrs array is nil",
			addrs:      nil,
			wantedData: "",
		},
		{
			name:       "addrs array is empty",
			addrs:      []next.InterfaceAddr{},
			wantedData: "",
		},
		{
			name: "addrs array contains one valid ipv4 address",
			addrs: []next.InterfaceAddr{
				{
					Addr: "191.22.3.1/29",
				},
			},
			wantedData: "191.22.3.1",
		},
		{
			name: "addrs array contains two valid ipv4 addresses",
			addrs: []next.InterfaceAddr{
				{
					Addr: "191.22.3.1/29",
				},
				{
					Addr: "191.22.3.4/21",
				},
			},
			wantedData: "191.22.3.1",
		},
		{
			name: "addrs array contains one valid and one non-valid ipv4 addresses",
			addrs: []next.InterfaceAddr{
				{
					Addr: "191.22.3.1/29",
				},
				{
					Addr: "fe80::1cde:b7e4:d558:d0dd/64",
				},
			},
			wantedData: "191.22.3.1",
		},
		{
			name: "addrs array contains one non-valid and one valid ipv4 addresses",
			addrs: []next.InterfaceAddr{
				{
					Addr: "fe80::1cde:b7e4:d558:d0dd/64",
				},
				{
					Addr: "191.22.3.1/29",
				},
			},
			wantedData: "191.22.3.1",
		},
		{
			name: "addrs array contains a non-valid ipv4 address",
			addrs: []next.InterfaceAddr{
				{
					Addr: "257.22.3.1/24",
				},
			},
			wantedData: "",
		},
		{
			name: "addrs array contains an empty ipv4 address",
			addrs: []next.InterfaceAddr{
				{
					Addr: "...1/29",
				},
			},
			wantedData: "",
		},
		{
			name: "addrs array contains a non slash separator",
			addrs: []next.InterfaceAddr{
				{
					Addr: "192.168.1.1:90",
				},
			},
			wantedData: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ipv4 := ExtractIPv4(tc.addrs)
			assert.Equal(t, tc.wantedData, ipv4)
		})
	}
}

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		netInfo    []next.InterfaceStat
		wantedData []rpi.Net
		wantedErr  error
	}{
		{
			name:       "success: netInfo array is empty",
			netInfo:    []next.InterfaceStat{},
			wantedData: nil,
			wantedErr:  nil,
		},

		{
			name: "success: netInfo contains one ip address",
			netInfo: []next.InterfaceStat{
				{
					Index: 1,
					Name:  "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					Addrs: []next.InterfaceAddr{
						{
							Addr: "192.168.11.58",
						},
					},
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
					IPv4: "192.168.11.58",
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: netInfo contains two ip addresses",
			netInfo: []next.InterfaceStat{
				{
					Index: 1,
					Name:  "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					Addrs: []next.InterfaceAddr{
						{
							Addr: "192.168.11.58/289",
						},
						{
							Addr: "255.1.1.0/29",
						},
					},
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
					IPv4: "192.168.11.58",
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := net.NSYS(Net{})
			nets, err := s.List(tc.netInfo)
			assert.Equal(t, tc.wantedData, nets)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		id         int
		netInfo    []next.InterfaceStat
		netStats   []next.IOCountersStat
		wantedData rpi.Net
		wantedErr  error
	}{
		{
			name:       "success: netInfo and netStats arrays are empty",
			id:         1,
			netInfo:    []next.InterfaceStat{},
			netStats:   []next.IOCountersStat{},
			wantedData: rpi.Net{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "net interface 1 does not exist"),
		},
		{
			name: "success: netStats array is empty",
			id:   1,
			netInfo: []next.InterfaceStat{
				{
					Index: 1,
					Name:  "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					Addrs: []next.InterfaceAddr{
						{
							Addr: "192.168.11.58/29",
						},
						{
							Addr: "255.1.1.0",
						},
					},
				},
			},
			netStats: []next.IOCountersStat{},
			wantedData: rpi.Net{
				ID:   1,
				Name: "interface1",
				Flags: []string{
					"flag1",
					"flag2",
				},
				IPv4:        "192.168.11.58",
				BytesSent:   0,
				BytesRecv:   0,
				PacketsSent: 0,
				PacketsRecv: 0,
			},
			wantedErr: nil,
		},
		{
			name: "success: netStats array does not contains the input interface id",
			id:   1,
			netInfo: []next.InterfaceStat{
				{
					Index: 1,
					Name:  "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					Addrs: []next.InterfaceAddr{
						{
							Addr: "192.168.11.58/29",
						},
						{
							Addr: "255.1.1.0",
						},
					},
				},
			},
			netStats: []next.IOCountersStat{
				{
					Name:        "interface2",
					BytesSent:   1,
					BytesRecv:   2,
					PacketsSent: 3,
					PacketsRecv: 4,
				},
			},
			wantedData: rpi.Net{
				ID:   1,
				Name: "interface1",
				Flags: []string{
					"flag1",
					"flag2",
				},
				IPv4:        "192.168.11.58",
				BytesSent:   0,
				BytesRecv:   0,
				PacketsSent: 0,
				PacketsRecv: 0,
			},
			wantedErr: nil,
		},
		{
			name:    "success: netInfo array is empty",
			id:      1,
			netInfo: []next.InterfaceStat{},
			netStats: []next.IOCountersStat{
				{
					BytesSent:   1,
					BytesRecv:   2,
					PacketsSent: 3,
					PacketsRecv: 4,
				},
			},
			wantedData: rpi.Net{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "net interface 1 does not exist"),
		},
		{
			name: "success",
			id:   1,
			netInfo: []next.InterfaceStat{
				{
					Index: 1,
					Name:  "interface1",
					Flags: []string{
						"flag1",
						"flag2",
					},
					Addrs: []next.InterfaceAddr{
						{
							Addr: "192.168.11.58/29",
						},
						{
							Addr: "255.1.1.0",
						},
					},
				},
			},
			netStats: []next.IOCountersStat{
				{
					Name:        "interface1",
					BytesSent:   1,
					BytesRecv:   2,
					PacketsSent: 3,
					PacketsRecv: 4,
				},
			},
			wantedData: rpi.Net{
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
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := net.NSYS(Net{})
			nets, err := s.View(tc.id, tc.netInfo, tc.netStats)
			assert.Equal(t, tc.wantedData, nets)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

// func TestView(t *testing.T) {
// 	cases := []struct {
// 		name       string
// 		terminal   string
// 		users      []host.UserStat
// 		wantedData rpi.User
// 		wantedErr  error
// 	}{
// 		{
// 			name:     "error: terminal does not exist",
// 			terminal: "T2",
// 			users: []host.UserStat{
// 				{
// 					User:     "U1",
// 					Terminal: "T1",
// 					Started:  11111,
// 				},
// 			},
// 			wantedData: rpi.User{},
// 			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "T2 does not exist"),
// 		},
// 		{
// 			name:     "success",
// 			terminal: "T1",
// 			users: []host.UserStat{
// 				{
// 					User:     "U1",
// 					Terminal: "T1",
// 					Started:  11111,
// 				},
// 				{
// 					User:     "U2",
// 					Terminal: "T2",
// 					Started:  22222,
// 				},
// 			},
// 			wantedData: rpi.User{
// 				User:     "U1",
// 				Terminal: "T1",
// 				Started:  11111,
// 			},
// 			wantedErr: nil,
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			s := user.USYS(sys.User{})
// 			users, err := s.View(tc.terminal, tc.users)
// 			assert.Equal(t, tc.wantedData, users)
// 			assert.Equal(t, tc.wantedErr, err)
// 		})
// 	}
// }
