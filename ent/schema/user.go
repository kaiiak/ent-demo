package schema

import (
	"regexp"
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Positive(),
		field.String("name").
			Default("张三"),
		field.String("phone").
			Match(regexp.MustCompile("^1[3-9]\\d{9}$")).
			Unique().
			Optional().
			Nillable(),
		field.JSON("tags", []string{}).
			Optional(),
		field.Int64("created_at").
			Default(time.Now().Unix()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From(),
	}
}
