package infos_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	cases := []struct {
		name       string
		filepath   string
		wantedData []string
		wantedErr  error
	}{
		{
			name:     "success",
			filepath: "./testdata/passwd",
			wantedData: []string{
				"root:x:0:0:root:/root:/bin/bash",
				"daemon:x:1:1:daemon:/usr/sbin:/usr/sbin/nologin",
				"bin:x:2:2:bin:/bin:/usr/sbin/nologin",
				"sys:x:3:3:sys:/dev:/usr/sbin/nologin",
				"sync:x:4:65534:sync:/bin:/bin/sync",
				"games:x:5:60:games:/usr/games:/usr/sbin/nologin",
				"man:x:6:12:man:/var/cache/man:/usr/sbin/nologin",
				"lp:x:7:7:lp:/var/spool/lpd:/usr/sbin/nologin",
				"mail:x:8:8:mail:/var/mail:/usr/sbin/nologin",
				"news:x:9:9:news:/var/spool/news:/usr/sbin/nologin",
				"uucp:x:10:10:uucp:/var/spool/uucp:/usr/sbin/nologin",
				"proxy:x:13:13:proxy:/bin:/usr/sbin/nologin",
				"www-data:x:33:33:www-data:/var/www:/usr/sbin/nologin",
				"backup:x:34:34:backup:/var/backups:/usr/sbin/nologin",
				"list:x:38:38:Mailing List Manager:/var/list:/usr/sbin/nologin",
				"irc:x:39:39:ircd:/var/run/ircd:/usr/sbin/nologin",
				"gnats:x:41:41:Gnats Bug-Reporting System (admin):/var/lib/gnats:/usr/sbin/nologin",
				"nobody:x:65534:65534:nobody:/nonexistent:/usr/sbin/nologin",
				"systemd-timesync:x:100:102:systemd Time Synchronization,,,:/run/systemd:/usr/sbin/nologin",
				"systemd-network:x:101:103:systemd Network Management,,,:/run/systemd:/usr/sbin/nologin",
				"systemd-resolve:x:102:104:systemd Resolver,,,:/run/systemd:/usr/sbin/nologin",
				"_apt:x:103:65534::/nonexistent:/usr/sbin/nologin",
				"pi:x:1000:1000:,,,:/home/pi:/bin/bash",
				"#pi2:x:1001:1001:,,,:/home/pi2:/bin/bash",
				"messagebus:x:104:110::/nonexistent:/usr/sbin/nologin",
				"_rpc:x:105:65534::/run/rpcbind:/usr/sbin/nologin",
				"statd:x:106:65534::/var/lib/nfs:/usr/sbin/nologin",
				"sshd:x:107:65534::/run/sshd:/usr/sbin/nologin",
				"avahi:x:108:113:Avahi mDNS daemon,,,:/var/run/avahi-daemon:/usr/sbin/nologin",
				"lightdm:x:109:114:Light Display Manager:/var/lib/lightdm:/bin/false",
				"systemd-coredump:x:996:996:systemd Core Dumper:/:/usr/sbin/nologin",
			},
			wantedErr: nil,
		},
		{
			name:       "error: failure opening file",
			filepath:   "",
			wantedData: nil,
			wantedErr:  fmt.Errorf("opening file failed"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			humanUser, err := i.ReadFile(tc.filepath)
			assert.Equal(t, tc.wantedData, humanUser)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestIsFileExists(t *testing.T) {
	cases := []struct {
		name       string
		filepath   string
		wantedData bool
		wantedErr  error
	}{
		{
			name:       "success: file does exist",
			filepath:   "./testdata/passwd",
			wantedData: true,
		},
		{
			name:       "success: file does not exist",
			filepath:   "./testdata/passwd-xxxxxxxx",
			wantedData: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			isFileExists := i.IsFileExists(tc.filepath)
			assert.Equal(t, tc.wantedData, isFileExists)
		})
	}
}

func TestGetConfigFiles(t *testing.T) {
	cases := []struct {
		name     string
		wantData map[string]rpi.ConfigFileDetails
	}{
		{
			name: "success",
			wantData: map[string]rpi.ConfigFileDetails{
				"bootconfig": {
					Path:        "/boot/config.txt",
					Description: "contains some system configuration parameters. It is read at boot time by the device.",
				},
				"etcpasswd": {
					Path:        "/etc/passwd",
					Description: "is a text-based database of information about users that may log into the system or other operating system user identities that own running processes.",
				},
				"waitfornetwork": {
					Path:        "/etc/systemd/system/dhcpcd.service.d/wait.conf",
					Description: "is a configuration file that forces the dhcp service to wait for the network to be configured before running.",
				},
				"hosts": {
					Path:        "/etc/hosts",
					Description: "is a text file that associates IP addresses with hostnames, one line per IP address.",
				},
				"hostname": {
					Path:        "/etc/hostname",
					Description: "configures the name of the local system. It contains a single newline-terminated hostname string.",
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			configFiles := infos.New().GetConfigFiles()
			assert.Equal(t, tt.wantData, configFiles)
		})
	}
}

func TestGetEnrichedConfigFiles(t *testing.T) {
	fileTestExisting := "/etc/passwd"
	statExisting, _ := os.Stat(fileTestExisting)

	fileTestNotExisting := "/boot/config.txt"

	cases := []struct {
		name     string
		args     map[string]rpi.ConfigFileDetails
		wantData map[string]rpi.ConfigFileDetails
	}{
		{
			name: "success",
			args: map[string]rpi.ConfigFileDetails{
				"etcpasswd": {
					Path:        fileTestExisting,
					Description: "dummy existing file",
				},
				"bootconfig": {
					Path:        fileTestNotExisting,
					Description: "dummy not existing file",
				},
			},
			wantData: map[string]rpi.ConfigFileDetails{
				"etcpasswd": {
					Path:         fileTestExisting,
					Name:         "passwd",
					Description:  "dummy existing file",
					IsExist:      true,
					Size:         statExisting.Size(),
					LastModified: uint64(statExisting.ModTime().Unix()),
				},
				"bootconfig": {
					Path:         fileTestNotExisting,
					Name:         "config.txt",
					Description:  "dummy not existing file",
					IsExist:      false,
					LastModified: 0,
					Size:         0,
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			configFiles := infos.New().GetEnrichedConfigFiles(tc.args)
			assert.Equal(t, tc.wantData, configFiles)
		})
	}
}
