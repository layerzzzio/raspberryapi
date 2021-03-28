package infos_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
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
					IsCritical:  true,
					Description: "contains some system configuration parameters. It is read at boot time by the device.",
				},
				"etcpasswd": {
					Path:        "/etc/passwd",
					IsCritical:  true,
					Description: "is a text-based database of information about users that may log into the system or other operating system user identities that own running processes.",
				},
				"waitfornetwork": {
					Path:        "/etc/systemd/system/dhcpcd.service.d/wait.conf",
					IsCritical:  false,
					Description: "is a configuration file that forces the dhcp service to wait for the network to be configured before running.",
				},
				"hosts": {
					Path:        "/etc/hosts",
					IsCritical:  true,
					Description: "is a text file that associates IP addresses with hostnames, one line per IP address.",
				},
				"hostname": {
					Path:        "/etc/hostname",
					IsCritical:  true,
					Description: "configures the name of the local system. It contains a single newline-terminated hostname string.",
				},
				"blanking": {
					Path:        "/etc/X11/xorg.conf.d/10-blanking.conf",
					IsCritical:  false,
					Description: "configures the blanking behavior of the monitor.",
				},
				"start_x_elf": {
					Path:        "/boot/start_x.elf",
					IsCritical:  true,
					Description: "is a binary blob (firmware) that is loaded on to the VideoCore in the SoC and that includes camera drivers and codec.",
				},
				"rgpio_public_conf": {
					Path:        "/etc/systemd/system/pigpiod.service.d/public.conf",
					IsCritical:  false,
					Description: "is the daemon file for the remote GPIO service.",
				},
				"iso3166": {
					Path:        "/usr/share/zoneinfo/iso3166.tab",
					IsCritical:  false,
					Description: "is a file containing the standards published by the International Organization for Standardization (ISO) that defines codes for the names of countries, dependent territories, special areas of geographical interest, and their principal subdivisions (e.g., provinces or states).",
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

func TestIsXscreenSaverInstalled(t *testing.T) {
	cases := []struct {
		name     string
		wantData bool
		wantErr  error
	}{
		{
			name:     "success",
			wantData: false,
			wantErr:  nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			command, err := infos.New().IsXscreenSaverInstalled()
			assert.Equal(t, tc.wantData, command)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestIsQuietGrep(t *testing.T) {
	cases := []struct {
		name       string
		command    string
		quietGrep  string
		grepType   string
		wantedData bool
	}{
		{
			name:       "success: 0 (with quiet)",
			command:    "pwd",
			quietGrep:  ".",
			grepType:   "quiet",
			wantedData: true,
		},
		{
			name:       "success: 1 (with quiet)",
			command:    "pwd",
			quietGrep:  "grep -q ABCDEFGHIJK",
			grepType:   "quiet",
			wantedData: false,
		},
		{
			name:       "success: 1 (with quiet)",
			command:    "cat /etc/passwd",
			quietGrep:  "ssh",
			grepType:   "quiet",
			wantedData: true,
		},
		{
			name:       "success: 1 (with word-regexp)",
			command:    "cat /etc/passwd",
			quietGrep:  "ssh",
			grepType:   "word-regexp",
			wantedData: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			isFileExists := i.IsQuietGrep(tc.command, tc.quietGrep, tc.grepType)
			assert.Equal(t, tc.wantedData, isFileExists)
		})
	}
}

func TestIsSSHKeyGenerating(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		addLines   []string
		wantedData bool
	}{
		{
			name: "success: there is not finished",
			path: "./test",
			addLines: []string{
				"yes man",
				"no it is not",
			},
			wantedData: true,
		},
		{
			name: "success: cannot find a line that starts with 'finished'",
			path: "./test",
			addLines: []string{
				"yes man",
				"no it is not",
				"   ^finished", // ^ == regex (start of the line)
			},
			wantedData: true,
		},
		{
			name: "success: find 'finished'",
			path: "./test",
			addLines: []string{
				"yes man",
				"no it is not",
				"finished",
			},
			wantedData: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			if err := actions.OverwriteToFile(actions.WriteToFileArg{
				File:        tc.path,
				Data:        tc.addLines,
				Multiline:   true,
				Permissions: 0755,
			}); err != nil {
				log.Fatal(err)
			}

			isSSHGen := i.IsSSHKeyGenerating(tc.path)

			os.Remove(tc.path)

			assert.Equal(t, tc.wantedData, isSSHGen)
		})
	}
}

func TestIsDPKGInstalled(t *testing.T) {
	cases := []struct {
		name        string
		packageName string
		wantedData  bool
	}{
		{
			name:        "success: 0 (with quiet)",
			packageName: "pwd",
			wantedData:  false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			isFileExists := i.IsDPKGInstalled(tc.packageName)
			assert.Equal(t, tc.wantedData, isFileExists)
		})
	}
}

func TestIsSPI(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		addLines   []string
		wantedData bool
	}{
		{
			name: "success: no match",
			path: "./test",
			addLines: []string{
				"yes man",
				"no it is not",
			},
			wantedData: false,
		},
		{
			name: "success: match",
			path: "./test",
			addLines: []string{
				"yes man",
				"no it is not",
				"dtparam=spi=on",
			},
			wantedData: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			if err := actions.OverwriteToFile(actions.WriteToFileArg{
				File:        tc.path,
				Data:        tc.addLines,
				Multiline:   true,
				Permissions: 0755,
			}); err != nil {
				log.Fatal(err)
			}

			isSPI := i.IsSPI(tc.path)

			os.Remove(tc.path)

			assert.Equal(t, tc.wantedData, isSPI)
		})
	}
}

func TestIsI2C(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		addLines   []string
		wantedData bool
	}{
		{
			name: "success: no match",
			path: "./test",
			addLines: []string{
				"yes man",
				"no it is not",
			},
			wantedData: false,
		},
		{
			name: "success: match",
			path: "./test",
			addLines: []string{
				"yes man",
				"no it is not",
				"device_tree_param=i2c=on",
			},
			wantedData: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			if err := actions.OverwriteToFile(actions.WriteToFileArg{
				File:        tc.path,
				Data:        tc.addLines,
				Multiline:   true,
				Permissions: 0755,
			}); err != nil {
				log.Fatal(err)
			}

			isI2C := i.IsI2C(tc.path)

			os.Remove(tc.path)

			assert.Equal(t, tc.wantedData, isI2C)
		})
	}
}

func TestIsVariableSet(t *testing.T) {
	cases := []struct {
		name       string
		rawLines   []string
		key        string
		value      string
		addLines   []string
		wantedData bool
	}{
		{
			name:  "success: no match",
			key:   "dtoverlay",
			value: "w1-gpio",
			rawLines: []string{
				"yes man",
				"no it is not",
			},
			wantedData: false,
		},
		{
			name:  "success: match",
			key:   "dtoverlay",
			value: "w1-gpio",
			rawLines: []string{
				"yes man",
				"no it is not",
				"dtoverlay=w1-gpio",
			},
			wantedData: true,
		},
		{
			name:  "success: match with spaces",
			key:   "dtoverlay",
			value: "w1-gpio",
			rawLines: []string{
				"yes man",
				"no it is not",
				"    dtoverlay =        w1-gpio #comment man",
			},
			wantedData: true,
		},
		{
			name:  "success: no match with comment",
			key:   "dtoverlay",
			value: "w1-gpio",
			rawLines: []string{
				"yes man",
				"no it is not",
				" #dtoverlay=w1-gpio",
			},
			wantedData: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			isMatch := i.IsVariableSet(tc.rawLines, tc.key, tc.value)
			assert.Equal(t, tc.wantedData, isMatch)
		})
	}
}

func TestListWifiInterfaces(t *testing.T) {
	currentDir, _ := os.Getwd()

	cases := []struct {
		name          string
		isCreateFile  bool
		directoryPath string
		wantedData    []string
	}{
		{
			name:          "success: cannot find the wireless file",
			isCreateFile:  false,
			directoryPath: currentDir,
			wantedData:    nil,
		},
		{
			name:          "success: found wireless file",
			isCreateFile:  true,
			directoryPath: currentDir,
			wantedData:    []string{"directory"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()

			if tc.isCreateFile {
				if err := os.MkdirAll(currentDir+"/directory", 0755); err != nil {
					log.Print(err)
				}

				if err := os.Mkdir(currentDir+"/directory/wireless", 0755); err != nil {
					log.Fatal(err)
				}
			}

			interfaces := i.ListWifiInterfaces(tc.directoryPath)

			if tc.isCreateFile {
				os.RemoveAll(currentDir + "/directory")
			}

			assert.Equal(t, tc.wantedData, interfaces)
		})
	}
}

func TestZoneInfo(t *testing.T) {
	cases := []struct {
		name       string
		filePath   string
		wantedData map[string]string
	}{
		{
			name:     "success: found wireless file",
			filePath: "./testdata/iso3166.tab",
			wantedData: map[string]string{
				"AD": "Andorra",
				"AE": "United Arab Emirates",
				"AF": "Afghanistan",
				"AG": "Antigua & Barbuda",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			interfaces := i.ZoneInfo(tc.filePath)
			assert.Equal(t, tc.wantedData, interfaces)
		})
	}
}

func TestListNameFilesInDirectory(t *testing.T) {
	cases := []struct {
		name          string
		directoryPath string
		wantedData    []string
	}{
		{
			name:          "success: found wireless file",
			directoryPath: "./testdata",
			wantedData: []string{
				"iso3166.tab", "passwd",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			interfaces := i.ListNameFilesInDirectory(tc.directoryPath)
			assert.Equal(t, tc.wantedData, interfaces)
		})
	}
}

func TestStringItemExists(t *testing.T) {
	cases := []struct {
		name       string
		item       string
		array      []string
		wantedData bool
	}{
		{
			name:       "success: array of strings",
			item:       "France",
			array:      []string{"India", "Canada", "Japan", "Germany", "France"},
			wantedData: true,
		},
		{
			name:       "failure: array of strings",
			item:       "Francexxx",
			array:      []string{"India", "Canada", "Japan", "Germany", "Italy"},
			wantedData: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			interfaces := infos.StringItemExists(tc.array, tc.item)
			assert.Equal(t, tc.wantedData, interfaces)
		})
	}
}

func TestVPNCountries(t *testing.T) {
	cases := []struct {
		name          string
		isCreateFile  bool
		directoryPath string
		wantedData    map[string][]string
	}{
		{
			name:          "success: the current file are not found in COUNTRYCODENAME",
			isCreateFile:  false,
			directoryPath: "./testdata",
			wantedData:    map[string][]string{},
		},
		{
			name:          "success: the current file are found in COUNTRYCODENAME",
			isCreateFile:  true,
			directoryPath: "./testdata",
			wantedData: map[string][]string{
				"nordvpn":   {"Germany"},
				"vyprvpn":   {"Canada"},
				"ipvanish":  {"France"},
				"surfshark": {"Singapore"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i := infos.New()
			if tc.isCreateFile {
				if err := os.MkdirAll("./testdata/wov_nordvpn/vpnconfigs", 0777); err != nil {
					log.Fatal(err)
				}

				if err := actions.OverwriteToFile(
					actions.WriteToFileArg{
						File: "./testdata/wov_nordvpn/vpnconfigs/de844.nordvpn.com.tcp.ovpn",
					},
				); err != nil {
					log.Fatal(err)
				}

				if err := os.MkdirAll("./testdata/wov_vyprvpn/vpnconfigs", 0777); err != nil {
					log.Fatal(err)
				}

				if err := actions.OverwriteToFile(
					actions.WriteToFileArg{
						File: "./testdata/wov_vyprvpn/vpnconfigs/Canada.ovpn",
					},
				); err != nil {
					log.Fatal(err)
				}

				if err := os.MkdirAll("./testdata/wov_ipvanish/vpnconfigs", 0777); err != nil {
					log.Fatal(err)
				}

				if err := actions.OverwriteToFile(
					actions.WriteToFileArg{
						File: "./testdata/wov_ipvanish/vpnconfigs/ipvanish-FR-Paris-par-a06.ovpn",
					},
				); err != nil {
					log.Fatal(err)
				}

				if err := os.MkdirAll("./testdata/wov_surfshark/vpnconfigs", 0777); err != nil {
					log.Fatal(err)
				}

				if err := actions.OverwriteToFile(
					actions.WriteToFileArg{
						File: "./testdata/wov_surfshark/vpnconfigs/sg-in.prod.surfshark.com_udp.ovpn",
					},
				); err != nil {
					log.Fatal(err)
				}

				if err := actions.OverwriteToFile(
					actions.WriteToFileArg{
						File: "./testdata/alcul",
					},
				); err != nil {
					log.Fatal(err)
				}
			}

			interfaces := i.VPNCountries(tc.directoryPath)

			if tc.isCreateFile {
				os.RemoveAll("./testdata/wov_nordvpn")
				os.RemoveAll("./testdata/wov_ipvanish")
				os.RemoveAll("./testdata/wov_vyprvpn")
				os.RemoveAll("./testdata/wov_surfshark")
				os.RemoveAll("./testdata/alcul")
			}

			assert.Equal(t, tc.wantedData, interfaces)
		})
	}
}
