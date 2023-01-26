package rpi

// HumanUser represents an active linux human user
type HumanUser struct {
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	Uid            int      `json:"uid"`
	Gid            int      `json:"gid"`
	AdditionalInfo []string `json:"additionalInfo"`
	HomeDirectory  string   `json:"homeDirectory"`
	DefaultShell   string   `json:"defaultShell"`
}
