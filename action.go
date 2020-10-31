package rpi

// Action represents the result of an action : Action = Σ Exec
type Action struct {
	Name          string         `json:"name"`
	Steps         map[int]string `json:"steps"`
	NumberOfSteps uint16         `json:"numberOfSteps"`
	Executions    []Exec         `json:"executions"`
	ExitStatus    uint8          `json:"exitStatus"`
	StartTime     uint64         `json:"startTime"`
	EndTime       uint64         `json:"endTime"`
}

// Exec represents the result of an execute.
type Exec struct {
	Name       string `json:"name"`
	StartTime  uint64 `json:"startTime"`
	EndTime    uint64 `json:"endTime"`
	ExitStatus uint8  `json:"exitStatus"`
	Stdin      string `json:"stdin,omitempty"`
	Stdout     string `json:"stdout,omitempty"`
	Stderr     string `json:"stderr,omitempty"`
}
