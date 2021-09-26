package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BinaryFile holds the schema definition for the BinaryFile entity.
type BinaryFile struct {
	ent.Schema
}

// Fields of the BinaryFile.
func (BinaryFile) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").
			NotEmpty().
			Unique(),
		field.Bytes("body").
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the BinaryFile.
func (BinaryFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Unique().
			Required().
			Ref("owned"),
	}
}
