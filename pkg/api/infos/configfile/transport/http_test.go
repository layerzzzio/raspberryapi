package transport_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/configfile"
	"github.com/raspibuddy/rpi/pkg/api/infos/configfile/transport"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var response rpi.ConfigFile

	cases := []struct {
		name         string
		cofsys       *mocksys.ConfigFile
		wantedStatus int
		wantedResp   rpi.ConfigFile
	}{
		{
			name:         "error: invalid request response",
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "error: List result is nil",
			cofsys: &mocksys.ConfigFile{
				ListFn: func(map[string]rpi.ConfigFileDetails) (rpi.ConfigFile, error) {
					return rpi.ConfigFile{}, errors.New("test error")
				},
			},
			wantedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			cofsys: &mocksys.ConfigFile{
				ListFn: func(map[string]rpi.ConfigFileDetails) (rpi.ConfigFile, error) {
					return rpi.ConfigFile{
						IsFilesMissing: true,
						ConfigFiles: []rpi.ConfigFileDetails{
							{
								Path:         "/etc/passwd",
								Name:         "passwd",
								IsExist:      true,
								Size:         1,
								LastModified: 2,
								Description:  "dummy desc",
							},
						},
					}, nil
				},
			},
			wantedStatus: http.StatusOK,
			wantedResp: rpi.ConfigFile{
				IsFilesMissing: true,
				ConfigFiles: []rpi.ConfigFileDetails{
					{
						Path:         "/etc/passwd",
						Name:         "passwd",
						IsExist:      true,
						Size:         1,
						LastModified: 2,
						Description:  "dummy desc",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			i := infos.New()
			s := configfile.New(tc.cofsys, i)
			transport.NewHTTP(s, rg)
			ts := httptest.NewServer(r)

			defer ts.Close()
			path := ts.URL + "/configfiles"
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			if tc.wantedResp.ConfigFiles != nil {
				if err := json.Unmarshal(body, &response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.wantedResp, response)
			}
			assert.Equal(t, tc.wantedStatus, res.StatusCode)
		})
	}
}
