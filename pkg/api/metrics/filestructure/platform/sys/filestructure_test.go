package sys

// func TestView(t *testing.T) {
// 	cases := []struct {
// 		name          string
// 		path          string
// 		fileStructure *rpi.File
// 		largestFiles  map[int64]string
// 		wantedData    rpi.FileStructure
// 		wantedErr     error
// 	}{
// 		{
// 			name: "success",
// 			path: "/dummy/path",
// 			fileStructure: &rpi.File{
// 				Name:  "/dummy/path",
// 				Size:  10,
// 				IsDir: false,
// 			},
// 			largestFiles: map[int64]string{10: "/dummy/path"},
// 			wantedData: rpi.FileStructure{
// 				DirectoryPath: "/dummy/path",
// 				LargestFiles:  map[string]int64{"/dummy/path": 10},
// 			},
// 			wantedErr: nil,
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			s := filestructure.FSSYS(FileStructure{})
// 			fstruct, err := s.ViewLF(tc.fileStructure, tc.largestFiles)
// 			assert.Equal(t, tc.wantedData, fstruct)
// 			assert.Equal(t, tc.wantedErr, err)
// 		})
// 	}
// }
