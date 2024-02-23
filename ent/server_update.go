// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/database64128/proxy-sharing-go/ent/account"
	"github.com/database64128/proxy-sharing-go/ent/node"
	"github.com/database64128/proxy-sharing-go/ent/predicate"
	"github.com/database64128/proxy-sharing-go/ent/server"
)

// ServerUpdate is the builder for updating Server entities.
type ServerUpdate struct {
	config
	hooks    []Hook
	mutation *ServerMutation
}

// Where appends a list predicates to the ServerUpdate builder.
func (su *ServerUpdate) Where(ps ...predicate.Server) *ServerUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *ServerUpdate) SetUpdateTime(t time.Time) *ServerUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetName sets the "name" field.
func (su *ServerUpdate) SetName(s string) *ServerUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *ServerUpdate) SetNillableName(s *string) *ServerUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (su *ServerUpdate) AddNodeIDs(ids ...int) *ServerUpdate {
	su.mutation.AddNodeIDs(ids...)
	return su
}

// AddNodes adds the "nodes" edges to the Node entity.
func (su *ServerUpdate) AddNodes(n ...*Node) *ServerUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return su.AddNodeIDs(ids...)
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (su *ServerUpdate) SetAccountID(id int) *ServerUpdate {
	su.mutation.SetAccountID(id)
	return su
}

// SetAccount sets the "account" edge to the Account entity.
func (su *ServerUpdate) SetAccount(a *Account) *ServerUpdate {
	return su.SetAccountID(a.ID)
}

// Mutation returns the ServerMutation object of the builder.
func (su *ServerUpdate) Mutation() *ServerMutation {
	return su.mutation
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (su *ServerUpdate) ClearNodes() *ServerUpdate {
	su.mutation.ClearNodes()
	return su
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (su *ServerUpdate) RemoveNodeIDs(ids ...int) *ServerUpdate {
	su.mutation.RemoveNodeIDs(ids...)
	return su
}

// RemoveNodes removes "nodes" edges to Node entities.
func (su *ServerUpdate) RemoveNodes(n ...*Node) *ServerUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return su.RemoveNodeIDs(ids...)
}

// ClearAccount clears the "account" edge to the Account entity.
func (su *ServerUpdate) ClearAccount() *ServerUpdate {
	su.mutation.ClearAccount()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ServerUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ServerUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ServerUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ServerUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ServerUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := server.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ServerUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := server.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Server.name": %w`, err)}
		}
	}
	if _, ok := su.mutation.AccountID(); su.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Server.account"`)
	}
	return nil
}

func (su *ServerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(server.Table, server.Columns, sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(server.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(server.FieldName, field.TypeString, value)
	}
	if su.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   server.NodesTable,
			Columns: []string{server.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedNodesIDs(); len(nodes) > 0 && !su.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   server.NodesTable,
			Columns: []string{server.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   server.NodesTable,
			Columns: []string{server.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   server.AccountTable,
			Columns: []string{server.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   server.AccountTable,
			Columns: []string{server.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{server.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ServerUpdateOne is the builder for updating a single Server entity.
type ServerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServerMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *ServerUpdateOne) SetUpdateTime(t time.Time) *ServerUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetName sets the "name" field.
func (suo *ServerUpdateOne) SetName(s string) *ServerUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *ServerUpdateOne) SetNillableName(s *string) *ServerUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (suo *ServerUpdateOne) AddNodeIDs(ids ...int) *ServerUpdateOne {
	suo.mutation.AddNodeIDs(ids...)
	return suo
}

// AddNodes adds the "nodes" edges to the Node entity.
func (suo *ServerUpdateOne) AddNodes(n ...*Node) *ServerUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return suo.AddNodeIDs(ids...)
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (suo *ServerUpdateOne) SetAccountID(id int) *ServerUpdateOne {
	suo.mutation.SetAccountID(id)
	return suo
}

// SetAccount sets the "account" edge to the Account entity.
func (suo *ServerUpdateOne) SetAccount(a *Account) *ServerUpdateOne {
	return suo.SetAccountID(a.ID)
}

// Mutation returns the ServerMutation object of the builder.
func (suo *ServerUpdateOne) Mutation() *ServerMutation {
	return suo.mutation
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (suo *ServerUpdateOne) ClearNodes() *ServerUpdateOne {
	suo.mutation.ClearNodes()
	return suo
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (suo *ServerUpdateOne) RemoveNodeIDs(ids ...int) *ServerUpdateOne {
	suo.mutation.RemoveNodeIDs(ids...)
	return suo
}

// RemoveNodes removes "nodes" edges to Node entities.
func (suo *ServerUpdateOne) RemoveNodes(n ...*Node) *ServerUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return suo.RemoveNodeIDs(ids...)
}

// ClearAccount clears the "account" edge to the Account entity.
func (suo *ServerUpdateOne) ClearAccount() *ServerUpdateOne {
	suo.mutation.ClearAccount()
	return suo
}

// Where appends a list predicates to the ServerUpdate builder.
func (suo *ServerUpdateOne) Where(ps ...predicate.Server) *ServerUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ServerUpdateOne) Select(field string, fields ...string) *ServerUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Server entity.
func (suo *ServerUpdateOne) Save(ctx context.Context) (*Server, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ServerUpdateOne) SaveX(ctx context.Context) *Server {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ServerUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ServerUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ServerUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := server.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ServerUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := server.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Server.name": %w`, err)}
		}
	}
	if _, ok := suo.mutation.AccountID(); suo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Server.account"`)
	}
	return nil
}

func (suo *ServerUpdateOne) sqlSave(ctx context.Context) (_node *Server, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(server.Table, server.Columns, sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Server.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, server.FieldID)
		for _, f := range fields {
			if !server.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != server.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(server.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(server.FieldName, field.TypeString, value)
	}
	if suo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   server.NodesTable,
			Columns: []string{server.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedNodesIDs(); len(nodes) > 0 && !suo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   server.NodesTable,
			Columns: []string{server.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   server.NodesTable,
			Columns: []string{server.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   server.AccountTable,
			Columns: []string{server.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   server.AccountTable,
			Columns: []string{server.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Server{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{server.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
