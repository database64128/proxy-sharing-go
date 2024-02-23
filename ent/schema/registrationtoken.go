package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// RegistrationToken holds the schema definition for the RegistrationToken entity.
type RegistrationToken struct {
	ent.Schema
}

// Mixin of the RegistrationToken.
func (RegistrationToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the RegistrationToken.
func (RegistrationToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty().
			Comment("Name is the unique name of the registration token."),
		field.Bytes("token").
			Unique().
			NotEmpty().
			Comment("Token is the registration token."),
	}
}

// Edges of the RegistrationToken.
func (RegistrationToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("registrations", Account.Type).
			Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

// Indexes of the RegistrationToken.
func (RegistrationToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
		index.Fields("token").
			Unique(),
	}
}
