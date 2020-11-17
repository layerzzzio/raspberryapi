package filestructure_test

// func TestList(t *testing.T) {
// 	cases := []struct {
// 		name       string
// 		path       string
// 		metrics    *mock.Metrics
// 		lfsys      mocksys.filestructure
// 		wantedData []rpi.filestructure
// 		wantedErr  error
// 	}{
// 		{
// 			name: "error: top100files array is nil",
// 			metrics: &mock.Metrics{
// 				Top100FilesFn: func(path string) ([]metrics.PathSize, string, error) {
// 					return nil, "", errors.New("test error info")
// 				},
// 			},
// 			wantedData: nil,
// 			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the largest files"),
// 		},
// 		{
// 			name: "success",
// 			path: "_System_Volumes_Data",
// 			metrics: &mock.Metrics{
// 				Top100FilesFn: func(path string) ([]metrics.PathSize, string, error) {
// 					return []metrics.PathSize{
// 						{
// 							Path: "/bin/file1",
// 							Size: 11,
// 						},
// 						{
// 							Path: "/usr/include/file2",
// 							Size: 22,
// 						},
// 						{
// 							Path: "/usr/dummy/file3",
// 							Size: 33,
// 						},
// 					}, "", nil
// 				},
// 			},
// 			lfsys: mocksys.filestructure{
// 				ViewFn: func([]metrics.PathSize) ([]rpi.filestructure, error) {
// 					return []rpi.filestructure{
// 						{
// 							Path:                "/bin/file1",
// 							Name:                "file1",
// 							Size:                11,
// 							Category:            "/bin",
// 							CategoryDescription: "represents some essential user command binaries",
// 						},
// 						{
// 							Path:                "/usr/include/file2",
// 							Name:                "file2",
// 							Size:                22,
// 							Category:            "/usr/include",
// 							CategoryDescription: "contains system general-use include files for the C programming language",
// 						},
// 						{
// 							Path:                "/usr/dummy/file3",
// 							Name:                "file3",
// 							Size:                33,
// 							Category:            "/usr",
// 							CategoryDescription: "contains shareable and read-only data",
// 						},
// 					}, nil
// 				},
// 			},
// 			wantedData: []rpi.filestructure{
// 				{
// 					Path:                "/bin/file1",
// 					Name:                "file1",
// 					Size:                11,
// 					Category:            "/bin",
// 					CategoryDescription: "represents some essential user command binaries",
// 				},
// 				{
// 					Path:                "/usr/include/file2",
// 					Name:                "file2",
// 					Size:                22,
// 					Category:            "/usr/include",
// 					CategoryDescription: "contains system general-use include files for the C programming language",
// 				},
// 				{
// 					Path:                "/usr/dummy/file3",
// 					Name:                "file3",
// 					Size:                33,
// 					Category:            "/usr",
// 					CategoryDescription: "contains shareable and read-only data",
// 				},
// 			},
// 			wantedErr: nil,
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			s := filestructure.New(tc.lfsys, tc.metrics)
// 			users, err := s.View(tc.path)
// 			assert.Equal(t, tc.wantedData, users)
// 			assert.Equal(t, tc.wantedErr, err)
// 		})
// 	}
// }
