package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Mixin of the Account.
func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			NotEmpty().
			Comment("Username is the unique username of the account."),
		field.Bytes("access_token").
			Unique().
			NotEmpty().
			Comment("AccessToken is the token used to access the account."),
		field.Bytes("refresh_token").
			Unique().
			NotEmpty().
			Comment("RefreshToken is the token used to refresh the account."),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servers", Server.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("nodes", Node.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("registration_token", RegistrationToken.Type).
			Ref("registrations").
			Unique().
			Required(),
	}
}

// Indexes of the Account.
func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username").
			Unique(),
		index.Fields("access_token").
			Unique(),
	}
}
