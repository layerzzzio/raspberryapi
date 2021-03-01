package rpi

type ConfigFile struct {
	IsFileMissing bool                `json:"isFileMissing"`
	ConfigFiles   []ConfigFileDetails `json:"configFiles"`
}

type ConfigFileDetails struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	Description  string `json:"description"`
	IsExist      bool   `json:"isExist"`
	LastModified uint64 `json:"lastModified,omitempty"`
	Size         int64  `json:"size,omitempty"`
}
