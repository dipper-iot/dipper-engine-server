package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dipper-iot/dipper-engine-server/constants"
	"time"
)

// RuleChan holds the schema definition for the RuleChan entity.
type RuleChan struct {
	ent.Schema
}

// Fields of the RuleChan.
func (RuleChan) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().Immutable(),
		field.String("name").NotEmpty().MaxLen(256),
		field.String("description").Nillable().MaxLen(1000),
		field.String("root_node").NotEmpty().MaxLen(256),
		field.Enum("status").Values("activated", "deactivated"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the RuleChan.
func (RuleChan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rules", RuleNode.Type),
		edge.To("sessions", Session.Type),
	}
}

func (RuleChan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constants.ChanTable},
	}
}
