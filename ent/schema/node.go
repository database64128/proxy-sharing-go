package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Mixin of the Node.
func (Node) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty().
			Comment("Name is the unique name of the node."),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("nodes").
			Unique().
			Required(),
		edge.From("server", Server.Type).
			Ref("nodes").
			Unique().
			Required(),
	}
}

// Indexes of the Node.
func (Node) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
		index.Edges("account"),
		index.Edges("server"),
	}
}
