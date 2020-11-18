package filestructure_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/filestructure"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"gopkg.in/go-playground/assert.v1"
)

func TestViewLF(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		pathSize   uint64
		fileLimit  int8
		metrics    *mock.Metrics
		fssys      mocksys.FileStructure
		wantedData rpi.FileStructure
		wantedErr  error
	}{
		{
			name:      "success",
			path:      "/dummy/path",
			pathSize:  1000,
			fileLimit: 1,
			metrics: &mock.Metrics{
				WalkFolderFn: func(
					string,
					metrics.ReadDir,
					uint64,
					int8,
					metrics.ShouldIgnoreFolder,
					chan int,
				) (*rpi.File, map[int64]string) {
					return &rpi.File{
							Name: "/dummy/path",
							Size: 1000,
							Files: []*rpi.File{
								{
									Name: "file1",
								},
								{
									Name: "file2",
								},
							},
						},
						map[int64]string{
							100: "/dummy/path/file1",
							200: "/dummy/path/file2",
						}
				},
			},
			fssys: mocksys.FileStructure{
				ViewFLFn: func(*rpi.File, map[int64]string) (rpi.FileStructure, error) {
					return rpi.FileStructure{
						DirectoryPath: "/dummy/path",
						LargestFiles: []*rpi.File{
							{
								Name: "/dummy/path/file1",
								Size: 10,
							},
							{
								Name: "/dummy/path/file2",
								Size: 20,
							},
						},
					}, nil
				},
			},
			wantedData: rpi.FileStructure{
				DirectoryPath: "/dummy/path",
				LargestFiles: []*rpi.File{
					{
						Name: "/dummy/path/file1",
						Size: 10,
					},
					{
						Name: "/dummy/path/file2",
						Size: 20,
					},
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := filestructure.New(tc.fssys, tc.metrics)
			largestFiles, err := s.ViewLF(tc.path, tc.pathSize, tc.fileLimit)
			assert.Equal(t, tc.wantedData, largestFiles)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
