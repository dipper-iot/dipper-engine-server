package models

type Node struct {
	ID       uint64                 `json:"id,omitempty"`
	NodeID   string                 `json:"node_id,omitempty"`
	Option   map[string]interface{} `json:"option,omitempty"`
	Infinite bool                   `json:"infinite,omitempty"`
	Debug    bool                   `json:"debug,omitempty"`
	End      bool                   `json:"end,omitempty"`
}
