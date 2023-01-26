package rpi

// Version represents the current API version.
type Version struct {
	RaspiBuddyVersion       string `json:"raspibuddyVersion"`
	RaspiBuddyDeployVersion string `json:"raspibuddyDeployVersion"`
}
