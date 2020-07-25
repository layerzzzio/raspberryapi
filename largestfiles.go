package rpi

// LargestFiles represents a top 100 largest file on the current host.
type LargestFiles struct {
	Size                int    `json:"size"`
	Path                string `json:"path"`
	Category            string `json:"category,omitempty"`
	CategoryDescription string `json:"categoryDescription,omitempty"`
}
