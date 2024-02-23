// Code generated by ent, DO NOT EDIT.

package registrationtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the registrationtoken type in the database.
	Label = "registration_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// EdgeRegistrations holds the string denoting the registrations edge name in mutations.
	EdgeRegistrations = "registrations"
	// Table holds the table name of the registrationtoken in the database.
	Table = "registration_tokens"
	// RegistrationsTable is the table that holds the registrations relation/edge.
	RegistrationsTable = "accounts"
	// RegistrationsInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	RegistrationsInverseTable = "accounts"
	// RegistrationsColumn is the table column denoting the registrations relation/edge.
	RegistrationsColumn = "registration_token_registrations"
)

// Columns holds all SQL columns for registrationtoken fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldToken,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator func([]byte) error
)

// OrderOption defines the ordering options for the RegistrationToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByRegistrationsCount orders the results by registrations count.
func ByRegistrationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRegistrationsStep(), opts...)
	}
}

// ByRegistrations orders the results by registrations terms.
func ByRegistrations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRegistrationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRegistrationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RegistrationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RegistrationsTable, RegistrationsColumn),
	)
}
