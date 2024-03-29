// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/database64128/proxy-sharing-go/ent/account"
	"github.com/database64128/proxy-sharing-go/ent/node"
	"github.com/database64128/proxy-sharing-go/ent/server"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	mutation *NodeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (nc *NodeCreate) SetCreateTime(t time.Time) *NodeCreate {
	nc.mutation.SetCreateTime(t)
	return nc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (nc *NodeCreate) SetNillableCreateTime(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetCreateTime(*t)
	}
	return nc
}

// SetUpdateTime sets the "update_time" field.
func (nc *NodeCreate) SetUpdateTime(t time.Time) *NodeCreate {
	nc.mutation.SetUpdateTime(t)
	return nc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (nc *NodeCreate) SetNillableUpdateTime(t *time.Time) *NodeCreate {
	if t != nil {
		nc.SetUpdateTime(*t)
	}
	return nc
}

// SetName sets the "name" field.
func (nc *NodeCreate) SetName(s string) *NodeCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (nc *NodeCreate) SetAccountID(id int) *NodeCreate {
	nc.mutation.SetAccountID(id)
	return nc
}

// SetAccount sets the "account" edge to the Account entity.
func (nc *NodeCreate) SetAccount(a *Account) *NodeCreate {
	return nc.SetAccountID(a.ID)
}

// SetServerID sets the "server" edge to the Server entity by ID.
func (nc *NodeCreate) SetServerID(id int) *NodeCreate {
	nc.mutation.SetServerID(id)
	return nc
}

// SetServer sets the "server" edge to the Server entity.
func (nc *NodeCreate) SetServer(s *Server) *NodeCreate {
	return nc.SetServerID(s.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nc *NodeCreate) Mutation() *NodeMutation {
	return nc.mutation
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NodeCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NodeCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NodeCreate) defaults() {
	if _, ok := nc.mutation.CreateTime(); !ok {
		v := node.DefaultCreateTime()
		nc.mutation.SetCreateTime(v)
	}
	if _, ok := nc.mutation.UpdateTime(); !ok {
		v := node.DefaultUpdateTime()
		nc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NodeCreate) check() error {
	if _, ok := nc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Node.create_time"`)}
	}
	if _, ok := nc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Node.update_time"`)}
	}
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Node.name"`)}
	}
	if v, ok := nc.mutation.Name(); ok {
		if err := node.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Node.name": %w`, err)}
		}
	}
	if _, ok := nc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "Node.account"`)}
	}
	if _, ok := nc.mutation.ServerID(); !ok {
		return &ValidationError{Name: "server", err: errors.New(`ent: missing required edge "Node.server"`)}
	}
	return nil
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NodeCreate) createSpec() (*Node, *sqlgraph.CreateSpec) {
	var (
		_node = &Node{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(node.Table, sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt))
	)
	if value, ok := nc.mutation.CreateTime(); ok {
		_spec.SetField(node.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := nc.mutation.UpdateTime(); ok {
		_spec.SetField(node.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := nc.mutation.Name(); ok {
		_spec.SetField(node.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := nc.mutation.AccountIDs(); len(nodes) > 0 {
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
		_node.account_nodes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ServerIDs(); len(nodes) > 0 {
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
		_node.server_nodes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	err      error
	builders []*NodeCreate
}

// Save creates the Node entities in the database.
func (ncb *NodeCreateBulk) Save(ctx context.Context) ([]*Node, error) {
	if ncb.err != nil {
		return nil, ncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Node, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NodeCreateBulk) SaveX(ctx context.Context) []*Node {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NodeCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
