package util

import (
	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine-server/models"
)

func InputNode(node *model.InputNode) *models.Node {

	r := &models.Node{
		NodeID: node.NodeID,
		RuleID: node.RuleID,
		Debug:  false,
		End:    false,
		Option: node.Option,
	}

	if node.End != nil {
		r.End = *node.End
	}

	if node.Debug != nil {
		r.Debug = *node.Debug
	}

	return r
}
