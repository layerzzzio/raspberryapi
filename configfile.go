package rpi

type ConfigFile struct {
	IsFilesMissing         bool                `json:"isFilesMissing"`
	IsCriticalFilesMissing bool                `json:"isCriticalFilesMissing"`
	FilesMissing           []string            `json:"filesMissing"`
	CriticalFilesMissing   []string            `json:"criticalFilesMissing"`
	ConfigFiles            []ConfigFileDetails `json:"configFiles"`
}

type ConfigFileDetails struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	IsCritical   bool   `json:"isCritical"`
	Description  string `json:"description"`
	IsExist      bool   `json:"isExist"`
	LastModified uint64 `json:"lastModified,omitempty"`
	Size         int64  `json:"size,omitempty"`
}
