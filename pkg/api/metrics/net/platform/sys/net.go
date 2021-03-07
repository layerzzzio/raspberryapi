package sys

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/net"
)

// Net represents an empty Net entity on the current system.
type Net struct{}

// List returns a list of Net info
func (n Net) List(netInfo []net.InterfaceStat) ([]rpi.Net, error) {
	var result []rpi.Net

	for i := range netInfo {
		data := rpi.Net{
			ID:    netInfo[i].Index,
			Name:  netInfo[i].Name,
			Flags: netInfo[i].Flags,
			IPv4:  ExtractIPv4(netInfo[i].Addrs),
		}
		result = append(result, data)
	}

	return result, nil
}

// View returns a list of Net info and stats
func (n Net) View(id int, netInfo []net.InterfaceStat, netStats []net.IOCountersStat) (rpi.Net, error) {
	var result rpi.Net
	isIDFound := false

	if len(netInfo) > 0 && len(netStats) > 0 {
		for i := range netInfo {
			if id == netInfo[i].Index {
				for j := range netStats {
					if netInfo[i].Name == netStats[j].Name {
						result = rpi.Net{
							ID:          netInfo[i].Index,
							Name:        netInfo[i].Name,
							Flags:       netInfo[i].Flags,
							IPv4:        ExtractIPv4(netInfo[i].Addrs),
							BytesSent:   netStats[i].BytesSent,
							BytesRecv:   netStats[i].BytesRecv,
							PacketsSent: netStats[i].PacketsSent,
							PacketsRecv: netStats[i].PacketsRecv,
						}
						isIDFound = true
						break
					} else {
						result = rpi.Net{
							ID:    netInfo[i].Index,
							Name:  netInfo[i].Name,
							Flags: netInfo[i].Flags,
							IPv4:  ExtractIPv4(netInfo[i].Addrs),
						}
						isIDFound = true
						break
					}
				}
			}
		}
	} else if len(netInfo) > 0 && len(netStats) == 0 {
		for i := range netInfo {
			if id == netInfo[i].Index {
				result = rpi.Net{
					ID:    netInfo[i].Index,
					Name:  netInfo[i].Name,
					Flags: netInfo[i].Flags,
					IPv4:  ExtractIPv4(netInfo[i].Addrs),
				}
				isIDFound = true
				break
			}
		}
	}

	if !isIDFound {
		return rpi.Net{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("net interface %v does not exist", id))
	}

	return result, nil
}

// ExtractIPv4 extracts IP from string
func ExtractIPv4(addrs []net.InterfaceAddr) string {
	var ip string
	reNum := regexp.MustCompile("[^\"/,]{10,}")
	rev4 := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	if len(addrs) == 0 {
		return ""
	}

	var ipRaw string
	for i := range addrs {
		ipRaw += addrs[i].Addr + ", "
	}

	if len(rev4.FindAllString(ipRaw, -1)) == 0 {
		return ""
	}

	fullIP := reNum.FindAllString(ipRaw, -1)
	cleanIP := strings.TrimSpace(rev4.FindAllString(ipRaw, -1)[0])

	for i := range fullIP {
		if strings.TrimSpace(fullIP[i]) == cleanIP {
			ip = cleanIP
			break
		}
	}

	return ip
}
