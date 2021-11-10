package infos

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/karrick/godirwalk"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/constants"
)

// Service represents several system scripts.
type Service struct{}

// Infos represents multiple system functions to get infos about the current system.
type Infos interface{}

// New creates a service instance.
func New() *Service {
	return &Service{}
}

// ReadFile reads a file
func (s Service) ReadFile(filePath string) ([]string, error) {
	var result []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening file failed")
	}

	// replacing line logic
	// source: https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	reader := bufio.NewReader(file)
	var line string

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if line != "" {
			result = append(result, strings.TrimSuffix(line, "\n"))
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return nil, fmt.Errorf("reading file failed")
	}

	if err = file.Close(); err != nil {
		return nil, fmt.Errorf("closing file failed")
	}

	return result, nil
}

// IsFileExists checks if a file exists
func (s Service) IsFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else {
		return false
	}
}

// IsDirectory determines if a file represented
// by `path` is a directory or not
// func (s Service) IsDirectory(path string) (bool, error) {
// 	fileInfo, err := os.Stat(path)
// 	if err != nil {
// 		fmt.Println(err)
// 		return false, err
// 	}

// 	return fileInfo.IsDir(), err
// }

// HasDeepestDirectoryFiles check if a parent directory contains at least one file in its child directories
func (s Service) HasDirectoryAtLeastOneFile(directoryPath string, isIgnoreZip bool) bool {
	result := false

	if s.IsFileExists(directoryPath) {
		err := godirwalk.Walk(directoryPath, &godirwalk.Options{
			Callback: func(osPathname string, de *godirwalk.Dirent) error {
				if strings.Contains(osPathname, ".zip") == isIgnoreZip {
					return godirwalk.SkipThis
				}

				if de.IsRegular() {
					result = true
				}

				if result {
					return fmt.Errorf("found a file")
				}

				return nil
			},
			// (optional) set true for faster yet non-deterministic enumeration (see godoc)
			Unsorted: true,
		})

		if err != nil {
			if err.Error() != "found a file" {
				log.Fatal(err)
			}
		}

	}

	return result
}

// GetConfigFiles returns a map of the unix config file used in the Raspberry Pi
// source: https://qvault.io/2019/10/21/golang-constant-maps-slices/
func (s Service) GetConfigFiles() map[string]rpi.ConfigFileDetails {
	return map[string]rpi.ConfigFileDetails{
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
	}
}

// GetEnrichedConfigFiles returns the list of config file with some extra fields
func (s Service) GetEnrichedConfigFiles(configFiles map[string]rpi.ConfigFileDetails) map[string]rpi.ConfigFileDetails {
	for k, v := range configFiles {
		stat, err := os.Stat(v.Path)
		if err != nil {
			configFiles[k] = rpi.ConfigFileDetails{
				Path:        v.Path,
				Description: v.Description,
				Name:        filepath.Base(v.Path),
				IsExist:     false,
			}
		} else {
			configFiles[k] = rpi.ConfigFileDetails{
				Path:         v.Path,
				Description:  v.Description,
				Name:         filepath.Base(v.Path),
				IsExist:      true,
				LastModified: uint64(stat.ModTime().Unix()),
				Size:         stat.Size(),
			}
		}
	}

	return configFiles
}

// IsXscreenSaverInstalled checks if XscreenSaver is installed
func (s Service) IsXscreenSaverInstalled() (bool, error) {
	isInstalled := false

	res, err := exec.Command(
		"sh",
		"-c",
		"dpkg -l xscreensaver | tail -n 1 | cut -d ' ' -f 1",
	).Output()

	if err != nil {
		err = fmt.Errorf("checking xscreensaver installation failed")
	} else {
		if string(res) == "ii" {
			isInstalled = true
		}
	}

	return isInstalled, err
}

// IsCommandStatus checks if a command returns 0 (found) or 1 (not found)
func (s Service) IsQuietGrep(command string, grep string, grepType string) bool {
	var grepCommand string

	if grepType == "quiet" {
		grepCommand = fmt.Sprintf("%v | grep -q %v ; echo $?", command, grep)
	} else if grepType == "word-regexp" {
		grepCommand = fmt.Sprintf("%v | grep -q -w %v ; echo $?", command, grep)
	} else {
		log.Fatal("bad grep type")
	}

	res, err := exec.Command("sh", "-c", grepCommand).Output()
	if err != nil {
		log.Fatal(err)
	}

	resNum, err := strconv.Atoi(strings.TrimSpace(string(res)))
	if err != nil {
		log.Fatal(err)
	}

	if resNum > 0 {
		return false
	} else {
		return true
	}
}

// IsSSHKeyGenerating checks if SSH keys are getting generated
func (s Service) IsSSHKeyGenerating(path string) bool {
	isFinished, err := exec.Command(
		"sh",
		"-c",
		fmt.Sprintf("grep -q \"^finished\" %v ; echo $?", path),
	).Output()

	if err != nil {
		log.Fatal(err)
	}

	isFinishedNum, err := strconv.Atoi(strings.TrimSpace(string(isFinished)))
	if err != nil {
		log.Fatal(err)
	}

	isSSHLogExist := s.IsFileExists(path)

	if isFinishedNum == 1 && isSSHLogExist {
		return true
	} else {
		return false
	}
}

// IsDPKGInstalled checks if package is installed with dpkg
func (s Service) IsDPKGInstalled(packageName string) bool {
	command := fmt.Sprintf("dpkg -l %v 2> /dev/null | tail -n 1 | cut -d ' ' -f 1", packageName)
	res, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}

	if strings.TrimSpace(string(res)) == "ii" {
		return true
	} else {
		return false
	}
}

// IsSPI checks if SPI is enabled or disabled
func (s Service) IsSPI(path string) bool {
	command := fmt.Sprintf(
		"grep -q -E \"^(device_tree_param|dtparam)=([^,]*,)*spi(=(on|true|yes|1))?(,.*)?$\" %v ; echo $?",
		path,
	)
	res, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}

	resNum, err := strconv.Atoi(strings.TrimSpace(string(res)))
	if err != nil {
		log.Fatal(err)
	}

	if resNum == 1 {
		return false
	} else {
		return true
	}
}

// IsI2C checks if I2C is enabled or disabled
func (s Service) IsI2C(path string) bool {
	command := fmt.Sprintf(
		"grep -q -E \"^(device_tree_param|dtparam)=([^,]*,)*i2c(_arm)?(=(on|true|yes|1))?(,.*)?$\" %v ; echo $?",
		path,
	)
	res, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}

	resNum, err := strconv.Atoi(strings.TrimSpace(string(res)))
	if err != nil {
		log.Fatal(err)
	}

	if resNum == 1 {
		return false
	} else {
		return true
	}
}

// IsVariableSet checks if a variable equals a certain value in a file
func (s Service) IsVariableSet(rawLines []string, key string, value string) bool {
	reg := `^\s*` + key + `\s*=\s*` + value + `\s*#?.*`

	isMatch := false

	for _, line := range rawLines {
		if line != "" {
			re := regexp.MustCompile(reg)
			if re.MatchString(line) {
				isMatch = true
			}
		}
	}

	return isMatch
}

// ListWifiInterfaces returns a list of wifi interfaces
func (s Service) ListWifiInterfaces(directoryPath string) []string {
	var wifiInterfaces []string

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal()
	}

	for _, f := range files {
		wirelessPath := fmt.Sprintf("%v/%v/wireless", directoryPath, f.Name())
		if s.IsFileExists(wirelessPath) {
			wifiInterfaces = append(wifiInterfaces, f.Name())
		}
	}

	return wifiInterfaces
}

func (s Service) IsWpaSupCom() map[string]bool {
	ifaces := s.ListWifiInterfaces(constants.NETWORKINTERFACES)
	result := map[string]bool{}

	for _, i := range ifaces {
		command := fmt.Sprintf(
			"wpa_cli -i %v status > /dev/null 2>&1 ; echo $?",
			i,
		)
		res, err := exec.Command("sh", "-c", command).Output()
		if err != nil {
			log.Fatal(err)
		}

		resNum, err := strconv.Atoi(strings.TrimSpace(string(res)))
		if err != nil {
			log.Fatal(err)
		}

		if resNum == 1 {
			result[i] = false
		} else {
			result[i] = true
		}
	}
	return result
}

func (s Service) ZoneInfo(filePath string) map[string]string {
	result := make(map[string]string)
	zi, err := s.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range zi {
		if !strings.HasPrefix(v, "#") {
			split := strings.SplitAfter(v, "\t")
			countryCode := strings.ReplaceAll(split[0], "\t", "")
			countryName := split[1]
			result[countryCode] = countryName
		}
	}

	return result
}

// ListNameFilesInDirectory lists all files in directory
func (s Service) ListNameFilesInDirectory(directoryPath string) []string {
	var result []string

	err := godirwalk.Walk(directoryPath, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if !de.IsDir() {
				result = append(result, de.Name())
			}
			return nil
		},
		// (optional) set true for faster yet non-deterministic enumeration (see godoc)
		Unsorted: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(result)

	return result
}

// VPNCountries lists countries available for vpn
func (s Service) VPNCountries(directoryPath string) map[string](map[string]string) {
	var result = make(map[string](map[string]string))
	// var countries []string
	var fileName []string
	var countryFiles = make(map[string](string))
	var regexCountry string

	if s.IsFileExists(directoryPath) {
		wovDir, err := ioutil.ReadDir(directoryPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, dir := range wovDir {
			if dir.IsDir() && dir.Name()[:4] == "wov_" {
				isValidDirectory := s.HasDirectoryAtLeastOneFile(directoryPath+"/"+dir.Name(), true)
				if isValidDirectory {
					err = godirwalk.Walk(directoryPath+"/"+dir.Name()+"/vpnconfigs", &godirwalk.Options{
						Callback: func(osPathname string, de *godirwalk.Dirent) error {
							re := regexp.MustCompile(`^[a-zA-Z]*`)

							if dir.Name() == "wov_ipvanish" {
								rawFileName := strings.ReplaceAll(de.Name(), "ipvanish-", "")
								regexCountry = strings.ReplaceAll(
									string(re.Find([]byte(rawFileName))),
									" ",
									"",
								)
							} else {
								regexCountry = strings.ReplaceAll(
									string(re.Find([]byte(de.Name()))),
									" ",
									"",
								)
							}

							if dir.Name() == "wov_vyprvpn" {
								// countries are clearly spelled
								if !de.IsDir() &&
									strings.HasSuffix(de.Name(), ".ovpn") &&
									!StringItemExists(fileName, regexCountry) {
									fileName = append(fileName, regexCountry)
									// osPathname is picked up randomly by the GoWalk
									countryFiles[regexCountry] = osPathname
								}
							} else if !de.IsDir() &&
								// countries are not spelled
								strings.HasSuffix(de.Name(), ".ovpn") &&
								!StringItemExists(fileName, regexCountry) {
								if country := constants.COUNTRYCODENAME[strings.ToUpper(regexCountry)]; country != "" {
									// osPathname is picked up randomly by the GoWalk
									countryFiles[country] = osPathname
								}
								fileName = append(fileName, regexCountry)
							}
							return nil
						},
						// (optional) set true for faster yet non-deterministic enumeration (see godoc)
						Unsorted: true,
					})
				}

				if err != nil {
					log.Fatal(err)
				}

				// sort.Strings(countries)
				result[strings.TrimPrefix(dir.Name(), "wov_")] = countryFiles
				countryFiles = make(map[string](string))
				// countries = nil
				fileName = nil
				regexCountry = ""
			}
		}
	}

	return result
}

func StringItemExists(array []string, item string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return true
		}
	}

	return false
}

// VPNConfigFiles returns a list of vpn files
func (s Service) VPNConfigFiles(
	vpnName string,
	vpnPath string,
	country string,
) []string {
	var countrycode string
	var result []string
	// vpnPath = /etc/openvpn/wov_ipvanish/vpnconfigs
	dir, err := ioutil.ReadDir(vpnPath)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range constants.COUNTRYCODENAME {
		if strings.EqualFold(v, country) {
			countrycode = k
		}
	}

	for _, dir := range dir {
		if vpnName == "vyprvpn" {
			if strings.Contains(dir.Name(), country) {
				return []string{vpnPath + "/" + dir.Name()}
			}
		} else {
			fileName := ""
			if vpnName == "ipvanish" {
				fileName = strings.ReplaceAll(dir.Name(), "ipvanish-", "")
			}

			if strings.HasPrefix(strings.ToUpper(fileName), countrycode) &&
				!strings.HasSuffix(fileName, "udp.ovpn") {
				result = append(result, vpnPath+"/"+dir.Name())
			}
		}

	}

	if result == nil {
		randomIndex := rand.Intn(len(dir))
		randomFileInfo := dir[randomIndex]
		result = []string{vpnPath + "/" + randomFileInfo.Name()}
	}

	return result
}

// ProcessesPids returns a list of pids
func (s Service) ProcessesPids(
	regex string,
) []string {
	psGrep := "ps -ef | grep"
	awk := "awk '{pid = $2 ; s = \"\"; for (i = 8; i <= NF; i++) s = s $i \" \"; print s \"<sep>\" pid}'"
	pidSearch := fmt.Sprintf(
		"%v \"%v\" | %v",
		psGrep,
		regex,
		awk,
	)

	out, err := exec.Command("sh", "-c", pidSearch).Output()

	if err != nil {
		log.Fatal(err)
	}

	stdOut := strings.Split(string(out), "\n")
	var pids []string
	// matched, err := regexp.MatchString(regex, "aaxbb")
	re, _ := regexp.Compile(regex)

	for _, ps := range stdOut {
		if strings.ReplaceAll(ps, " ", "") != "" &&
			!strings.Contains(ps, awk) &&
			!strings.HasPrefix(ps, "grep ") {
			split := strings.Split(ps, "<sep>")
			matched := re.MatchString(split[0])
			if matched {
				pids = append(pids, split[1])
			}
		}
	}

	return pids
}

// StatusVPNWithOpenVPN returns the status of the VPN with OpenVPN apps
func (s Service) StatusVPNWithOpenVPN(
	regexVPNPs string,
	regexVPNName string,
) map[string]bool {
	psGrep := "ps -ef | grep"
	awk := "awk '{pid = $2 ; s = \"\"; for (i = 8; i <= NF; i++) s = s $i \" \"; print s \"<sep>\" pid}'"
	pidSearch := fmt.Sprintf(
		"%v \"%v\" | %v",
		psGrep,
		regexVPNPs,
		awk,
	)

	out, err := exec.Command("sh", "-c", pidSearch).Output()

	if err != nil {
		log.Fatal(err)
	}

	stdOut := strings.Split(string(out), "\n")
	var result = make(map[string]bool)

	rePs, _ := regexp.Compile(regexVPNPs)
	reName, _ := regexp.Compile(regexVPNName)

	for _, ps := range stdOut {
		if strings.ReplaceAll(ps, " ", "") != "" &&
			!strings.Contains(ps, awk) &&
			!strings.HasPrefix(ps, "grep ") {
			split := strings.Split(ps, "<sep>")
			matched := rePs.MatchString(split[0])
			if matched {
				vpnNameRaw := string(reName.FindAllString(split[0], 1)[0])
				vpnNameClean := strings.ReplaceAll(vpnNameRaw, "wov_", "")
				result[vpnNameClean] = true
			}
		}
	}

	if len(result) == 0 {
		result = nil
	}

	return result
}
