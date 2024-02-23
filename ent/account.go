// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/database64128/proxy-sharing-go/ent/account"
	"github.com/database64128/proxy-sharing-go/ent/registrationtoken"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Username is the unique username of the account.
	Username string `json:"username,omitempty"`
	// AccessToken is the token used to access the account.
	AccessToken []byte `json:"access_token,omitempty"`
	// RefreshToken is the token used to refresh the account.
	RefreshToken []byte `json:"refresh_token,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges                            AccountEdges `json:"edges"`
	registration_token_registrations *int
	selectValues                     sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Servers holds the value of the servers edge.
	Servers []*Server `json:"servers,omitempty"`
	// Nodes holds the value of the nodes edge.
	Nodes []*Node `json:"nodes,omitempty"`
	// RegistrationToken holds the value of the registration_token edge.
	RegistrationToken *RegistrationToken `json:"registration_token,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ServersOrErr returns the Servers value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) ServersOrErr() ([]*Server, error) {
	if e.loadedTypes[0] {
		return e.Servers, nil
	}
	return nil, &NotLoadedError{edge: "servers"}
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) NodesOrErr() ([]*Node, error) {
	if e.loadedTypes[1] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// RegistrationTokenOrErr returns the RegistrationToken value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) RegistrationTokenOrErr() (*RegistrationToken, error) {
	if e.loadedTypes[2] {
		if e.RegistrationToken == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: registrationtoken.Label}
		}
		return e.RegistrationToken, nil
	}
	return nil, &NotLoadedError{edge: "registration_token"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldAccessToken, account.FieldRefreshToken:
			values[i] = new([]byte)
		case account.FieldID:
			values[i] = new(sql.NullInt64)
		case account.FieldUsername:
			values[i] = new(sql.NullString)
		case account.FieldCreateTime, account.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case account.ForeignKeys[0]: // registration_token_registrations
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case account.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				a.CreateTime = value.Time
			}
		case account.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				a.UpdateTime = value.Time
			}
		case account.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				a.Username = value.String
			}
		case account.FieldAccessToken:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value != nil {
				a.AccessToken = *value
			}
		case account.FieldRefreshToken:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value != nil {
				a.RefreshToken = *value
			}
		case account.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field registration_token_registrations", value)
			} else if value.Valid {
				a.registration_token_registrations = new(int)
				*a.registration_token_registrations = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryServers queries the "servers" edge of the Account entity.
func (a *Account) QueryServers() *ServerQuery {
	return NewAccountClient(a.config).QueryServers(a)
}

// QueryNodes queries the "nodes" edge of the Account entity.
func (a *Account) QueryNodes() *NodeQuery {
	return NewAccountClient(a.config).QueryNodes(a)
}

// QueryRegistrationToken queries the "registration_token" edge of the Account entity.
func (a *Account) QueryRegistrationToken() *RegistrationTokenQuery {
	return NewAccountClient(a.config).QueryRegistrationToken(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("create_time=")
	builder.WriteString(a.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(a.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(a.Username)
	builder.WriteString(", ")
	builder.WriteString("access_token=")
	builder.WriteString(fmt.Sprintf("%v", a.AccessToken))
	builder.WriteString(", ")
	builder.WriteString("refresh_token=")
	builder.WriteString(fmt.Sprintf("%v", a.RefreshToken))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account
