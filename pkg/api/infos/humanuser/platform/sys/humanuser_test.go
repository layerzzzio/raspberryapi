package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/humanuser"
	"github.com/raspibuddy/rpi/pkg/api/infos/humanuser/platform/sys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		lines      []string
		wantedData []rpi.HumanUser
		wantedErr  error
	}{
		{
			name: "success",
			lines: []string{
				"systemd-network:x:101:103:systemd Network Management,,,:/run/systemd:/usr/sbin/nologin",
				"systemd-resolve:x:102:104:systemd Resolver,,,:/run/systemd:/usr/sbin/nologin",
				"_apt:x:103:65534::/nonexistent:/usr/sbin/nologin",
				"pi:x:1000:1000:,,,:/home/pi:/bin/bash",
				"#pi2:x:1001:1001:,,,:/home/pi2:/bin/bash",
				"pi3:x:1002:1002:A,B,C,:/home/pi3:/bin/bash",
				"pi4:x:1003:1003::/home/pi4:/bin/bash",
				"messagebus:x:104:110::/nonexistent:/usr/sbin/nologin",
				"_rpc:x:105:65534::/run/rpcbind:/usr/sbin/nologin",
				"statd:x:106:65534::/var/lib/nfs:/usr/sbin/nologin",
			},
			wantedData: []rpi.HumanUser{
				{
					Username:       "pi",
					Password:       "x",
					Uid:            1000,
					Gid:            1000,
					AdditionalInfo: nil,
					HomeDirectory:  "/home/pi",
					DefaultShell:   "/bin/bash",
				},
				{
					Username:       "pi3",
					Password:       "x",
					Uid:            1002,
					Gid:            1002,
					AdditionalInfo: []string{"A", "B", "C"},
					HomeDirectory:  "/home/pi3",
					DefaultShell:   "/bin/bash",
				},
				{
					Username:       "pi4",
					Password:       "x",
					Uid:            1003,
					Gid:            1003,
					AdditionalInfo: nil,
					HomeDirectory:  "/home/pi4",
					DefaultShell:   "/bin/bash",
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := humanuser.HUMSYS(sys.HumanUser{})
			humanUsers, err := s.List(tc.lines)
			assert.Equal(t, tc.wantedData, humanUsers)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
