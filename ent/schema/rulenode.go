package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// RuleNode holds the schema definition for the RuleNode entity.
type RuleNode struct {
	ent.Schema
}

// Fields of the RuleNode.
func (RuleNode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("chain_id").Optional(),
		field.String("node_id").NotEmpty().MaxLen(20),
		field.JSON("option", map[string]interface{}{}).Default(map[string]interface{}{}),
		field.Bool("debug").Default(false),
		field.Bool("end").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the RuleNode.
func (RuleNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chain", RuleChan.Type).
			Ref("rules").
			Field("chain_id").
			Unique(),
	}
}

func (RuleNode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rule_nodes"},
	}
}
