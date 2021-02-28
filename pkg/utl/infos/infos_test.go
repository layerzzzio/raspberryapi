package infos_test

import (
	"fmt"
	"testing"

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
		wantData map[string]string
	}{
		{
			name: "success",
			wantData: map[string]string{
				"bootconfig":     "/boot/config.txt",
				"etcpasswd":      "/etc/passwd",
				"waitfornetwork": "/etc/systemd/system/dhcpcd.service.d/wait.conf",
				"hosts":          "/etc/hosts",
				"hostname":       "/etc/hostname",
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
