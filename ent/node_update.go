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

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	hooks    []Hook
	mutation *NodeMutation
}

// Where appends a list predicates to the NodeUpdate builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdateTime sets the "update_time" field.
func (nu *NodeUpdate) SetUpdateTime(t time.Time) *NodeUpdate {
	nu.mutation.SetUpdateTime(t)
	return nu
}

// SetName sets the "name" field.
func (nu *NodeUpdate) SetName(s string) *NodeUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableName(s *string) *NodeUpdate {
	if s != nil {
		nu.SetName(*s)
	}
	return nu
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (nu *NodeUpdate) SetAccountID(id int) *NodeUpdate {
	nu.mutation.SetAccountID(id)
	return nu
}

// SetAccount sets the "account" edge to the Account entity.
func (nu *NodeUpdate) SetAccount(a *Account) *NodeUpdate {
	return nu.SetAccountID(a.ID)
}

// SetServerID sets the "server" edge to the Server entity by ID.
func (nu *NodeUpdate) SetServerID(id int) *NodeUpdate {
	nu.mutation.SetServerID(id)
	return nu
}

// SetServer sets the "server" edge to the Server entity.
func (nu *NodeUpdate) SetServer(s *Server) *NodeUpdate {
	return nu.SetServerID(s.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nu *NodeUpdate) Mutation() *NodeMutation {
	return nu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (nu *NodeUpdate) ClearAccount() *NodeUpdate {
	nu.mutation.ClearAccount()
	return nu
}

// ClearServer clears the "server" edge to the Server entity.
func (nu *NodeUpdate) ClearServer() *NodeUpdate {
	nu.mutation.ClearServer()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NodeUpdate) defaults() {
	if _, ok := nu.mutation.UpdateTime(); !ok {
		v := node.UpdateDefaultUpdateTime()
		nu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NodeUpdate) check() error {
	if v, ok := nu.mutation.Name(); ok {
		if err := node.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Node.name": %w`, err)}
		}
	}
	if _, ok := nu.mutation.AccountID(); nu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Node.account"`)
	}
	if _, ok := nu.mutation.ServerID(); nu.mutation.ServerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Node.server"`)
	}
	return nil
}

func (nu *NodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(node.Table, node.Columns, sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdateTime(); ok {
		_spec.SetField(node.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.SetField(node.FieldName, field.TypeString, value)
	}
	if nu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.AccountTable,
			Columns: []string{node.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.AccountTable,
			Columns: []string{node.AccountColumn},
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
	if nu.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.ServerTable,
			Columns: []string{node.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ServerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.ServerTable,
			Columns: []string{node.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeMutation
}

// SetUpdateTime sets the "update_time" field.
func (nuo *NodeUpdateOne) SetUpdateTime(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetUpdateTime(t)
	return nuo
}

// SetName sets the "name" field.
func (nuo *NodeUpdateOne) SetName(s string) *NodeUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableName(s *string) *NodeUpdateOne {
	if s != nil {
		nuo.SetName(*s)
	}
	return nuo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (nuo *NodeUpdateOne) SetAccountID(id int) *NodeUpdateOne {
	nuo.mutation.SetAccountID(id)
	return nuo
}

// SetAccount sets the "account" edge to the Account entity.
func (nuo *NodeUpdateOne) SetAccount(a *Account) *NodeUpdateOne {
	return nuo.SetAccountID(a.ID)
}

// SetServerID sets the "server" edge to the Server entity by ID.
func (nuo *NodeUpdateOne) SetServerID(id int) *NodeUpdateOne {
	nuo.mutation.SetServerID(id)
	return nuo
}

// SetServer sets the "server" edge to the Server entity.
func (nuo *NodeUpdateOne) SetServer(s *Server) *NodeUpdateOne {
	return nuo.SetServerID(s.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nuo *NodeUpdateOne) Mutation() *NodeMutation {
	return nuo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (nuo *NodeUpdateOne) ClearAccount() *NodeUpdateOne {
	nuo.mutation.ClearAccount()
	return nuo
}

// ClearServer clears the "server" edge to the Server entity.
func (nuo *NodeUpdateOne) ClearServer() *NodeUpdateOne {
	nuo.mutation.ClearServer()
	return nuo
}

// Where appends a list predicates to the NodeUpdate builder.
func (nuo *NodeUpdateOne) Where(ps ...predicate.Node) *NodeUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NodeUpdateOne) Select(field string, fields ...string) *NodeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Node entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NodeUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdateTime(); !ok {
		v := node.UpdateDefaultUpdateTime()
		nuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NodeUpdateOne) check() error {
	if v, ok := nuo.mutation.Name(); ok {
		if err := node.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Node.name": %w`, err)}
		}
	}
	if _, ok := nuo.mutation.AccountID(); nuo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Node.account"`)
	}
	if _, ok := nuo.mutation.ServerID(); nuo.mutation.ServerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Node.server"`)
	}
	return nil
}

func (nuo *NodeUpdateOne) sqlSave(ctx context.Context) (_node *Node, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(node.Table, node.Columns, sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Node.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, node.FieldID)
		for _, f := range fields {
			if !node.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != node.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdateTime(); ok {
		_spec.SetField(node.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.SetField(node.FieldName, field.TypeString, value)
	}
	if nuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.AccountTable,
			Columns: []string{node.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.AccountTable,
			Columns: []string{node.AccountColumn},
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
	if nuo.mutation.ServerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.ServerTable,
			Columns: []string{node.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ServerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.ServerTable,
			Columns: []string{node.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Node{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
