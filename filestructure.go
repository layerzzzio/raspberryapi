package rpi

// FileStructure represents a linux file (file, link, directory) with some metrics
type FileStructure struct {
	DirectoryPath string  `json:"directoryPath"`
	Structure     *File   `json:"structure,omitempty"`
	LargestFiles  []*File `json:"largestfiles,omitempty"`
}

// File structure representing files and folders with their accumulated sizes
type File struct {
	Name   string  `json:"name,omitempty"`
	Path   string  `json:"path,omitempty"`
	Parent *File   `json:"parent,omitempty"`
	Size   int64   `json:"size,omitempty"`
	IsDir  bool    `json:"isDir,omitempty"`
	Files  []*File `json:"files,omitempty"`
}
