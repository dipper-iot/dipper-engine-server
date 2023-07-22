package models

type Node struct {
	ID     uint64                 `json:"id,omitempty"`
	NodeID string                 `json:"node_id,omitempty"`
	RuleID string                 `json:"rule_id"`
	Option map[string]interface{} `json:"option,omitempty"`
	Debug  bool                   `json:"debug,omitempty"`
	End    bool                   `json:"end,omitempty"`
}
