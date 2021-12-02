package transport_test

// NOT TESTABLE AS-IS : INDEED IT CANNOT FIND A FILE IN LIST METHOD
// func TestList(t *testing.T) {
// 	var response rpi.RpInterface

// 	cases := []struct {
// 		name         string
// 		intsys       *mocksys.RpInterface
// 		wantedStatus int
// 		wantedResp   rpi.RpInterface
// 	}{
// 		// {
// 		// 	name:         "error: invalid request response",
// 		// 	wantedStatus: http.StatusInternalServerError,
// 		// },
// 		// {
// 		// 	name: "error: List result is nil",
// 		// 	intsys: &mocksys.RpInterface{
// 		// 		ListFn: func(
// 		// 			[]string, bool, bool, bool, bool, bool, bool, bool, bool, bool, []string,
// 		// 		) (rpi.RpInterface, error) {
// 		// 			return rpi.RpInterface{}, errors.New("test error")
// 		// 		},
// 		// 	},
// 		// 	wantedStatus: http.StatusInternalServerError,
// 		// },
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			r := server.New()
// 			rg := r.Group("")
// 			i := infos.New()
// 			s := rpinterface.New(tc.intsys, i)
// 			transport.NewHTTP(s, rg)
// 			ts := httptest.NewServer(r)

// 			defer ts.Close()
// 			path := ts.URL + "/rpinterfaces"
// 			res, err := http.Get(path)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			if err != nil {
// 				panic(err)
// 			}

// 			if tc.wantedResp.IsCamera != rpi.RpInterface.IsCamera {
// 				if err := json.Unmarshal(body, &response); err != nil {
// 					t.Fatal(err)
// 				}
// 				assert.Equal(t, tc.wantedResp, response)
// 			}
// 			assert.Equal(t, tc.wantedStatus, res.StatusCode)
// 		})
// 	}
// }
