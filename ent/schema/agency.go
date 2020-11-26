package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Agency holds the schema definition for the Agency entity.
type Agency struct {
	ent.Schema
}

// Fields of the Agency.
func (Agency) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Positive(),
		field.String("name").
			NotEmpty(),
		field.Int64("owner_id").
			Positive(),
		field.Int64("member_limit").
			Default(100),
		field.Int64("created_at").
			Default(time.Now().Unix()),
	}
}

// Edges of the Agency.
func (Agency) Edges() []ent.Edge {
	return nil
}
