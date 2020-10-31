package rpi

// LargestFile represents a top 100 largest file on the current host.
type LargestFile struct {
	Size                int    `json:"size"`
	Name                string `json:"name"`
	Path                string `json:"path"`
	Category            string `json:"category,omitempty"`
	CategoryDescription string `json:"categoryDescription,omitempty"`
}