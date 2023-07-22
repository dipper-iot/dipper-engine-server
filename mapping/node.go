package mapping

import (
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine/data"
)

func MapNode(node *ent.RuleNode) *data.NodeRule {
	return &data.NodeRule{
		NodeId: node.NodeID,
		RuleId: node.RuleID,
		Debug:  node.Debug,
		End:    node.End,
		Option: node.Option,
	}
}
