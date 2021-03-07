package sys

import (
	"strconv"
	"strings"

	"github.com/raspibuddy/rpi"
)

// HumanUser represents a HumanUser entity on the current system.
type HumanUser struct{}

// List returns a list of HumanUser info
func (hu HumanUser) List(listUsers []string) ([]rpi.HumanUser, error) {
	var humanUsers []rpi.HumanUser

	for _, v := range listUsers {
		// skip all lines starting with #
		if equal := strings.Index(v, "#"); equal < 0 {
			// one array per line (separator = colon)
			// ex:
			// [pi x 1000 1000 ,,, /home/pi /bin/bash]
			// [_apt x 103 65534 /nonexistent /usr/sbin/nologin]
			// etc.
			lineSlice := strings.Split(v, ":")

			// business rule:
			// human users should have a uid >= 1000
			// and should ignore the "nobody" user
			if len(lineSlice) == 7 {
				uid, errU := strconv.Atoi(lineSlice[2])
				gid, errG := strconv.Atoi(lineSlice[3])

				if errU == nil && errG == nil && uid >= 1000 && lineSlice[0] != "nobody" {
					allAdditionalInfo := strings.FieldsFunc(lineSlice[4], func(divide rune) bool {
						return divide == ','
					})

					var additionalInfo []string
					if len(allAdditionalInfo) == 0 {
						additionalInfo = nil
					} else {
						for _, v := range allAdditionalInfo {
							if v != "" {
								additionalInfo = append(additionalInfo, v)
							}
						}
					}

					humanUsers = append(
						humanUsers,
						rpi.HumanUser{
							Username:       lineSlice[0],
							Password:       lineSlice[1],
							Uid:            uid,
							Gid:            gid,
							AdditionalInfo: additionalInfo,
							HomeDirectory:  lineSlice[5],
							DefaultShell:   lineSlice[6],
						},
					)
				}
			}
		}
	}

	return humanUsers, nil
}
