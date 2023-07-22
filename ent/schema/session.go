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

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().Immutable(),
		field.Uint64("chain_id").Optional(),
		field.Bool("is_test").Default(false),
		field.Bool("infinite").Default(false),
		field.JSON("data", map[string]interface{}{}).Default(map[string]interface{}{}),
		field.JSON("result", map[string]interface{}{}).Default(map[string]interface{}{}),
		field.Int("end_count").Default(0),
		field.Int("timeout").Default(30),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chain", RuleChan.Type).
			Ref("sessions").
			Field("chain_id").
			Unique(),
	}
}

func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constants.SessionTable},
	}
}
