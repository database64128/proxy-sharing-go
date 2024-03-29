// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/database64128/proxy-sharing-go/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/database64128/proxy-sharing-go/ent/account"
	"github.com/database64128/proxy-sharing-go/ent/node"
	"github.com/database64128/proxy-sharing-go/ent/registrationtoken"
	"github.com/database64128/proxy-sharing-go/ent/server"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Node is the client for interacting with the Node builders.
	Node *NodeClient
	// RegistrationToken is the client for interacting with the RegistrationToken builders.
	RegistrationToken *RegistrationTokenClient
	// Server is the client for interacting with the Server builders.
	Server *ServerClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Node = NewNodeClient(c.config)
	c.RegistrationToken = NewRegistrationTokenClient(c.config)
	c.Server = NewServerClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		Account:           NewAccountClient(cfg),
		Node:              NewNodeClient(cfg),
		RegistrationToken: NewRegistrationTokenClient(cfg),
		Server:            NewServerClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		Account:           NewAccountClient(cfg),
		Node:              NewNodeClient(cfg),
		RegistrationToken: NewRegistrationTokenClient(cfg),
		Server:            NewServerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Node.Use(hooks...)
	c.RegistrationToken.Use(hooks...)
	c.Server.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Account.Intercept(interceptors...)
	c.Node.Intercept(interceptors...)
	c.RegistrationToken.Intercept(interceptors...)
	c.Server.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AccountMutation:
		return c.Account.mutate(ctx, m)
	case *NodeMutation:
		return c.Node.mutate(ctx, m)
	case *RegistrationTokenMutation:
		return c.RegistrationToken.mutate(ctx, m)
	case *ServerMutation:
		return c.Server.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `account.Intercept(f(g(h())))`.
func (c *AccountClient) Intercept(interceptors ...Interceptor) {
	c.inters.Account = append(c.inters.Account, interceptors...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AccountClient) MapCreateBulk(slice any, setFunc func(*AccountCreate, int)) *AccountCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AccountCreateBulk{err: fmt.Errorf("calling to AccountClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AccountCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id int) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id int) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAccount},
		inters: c.Interceptors(),
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id int) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id int) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryServers queries the servers edge of a Account.
func (c *AccountClient) QueryServers(a *Account) *ServerQuery {
	query := (&ServerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.ServersTable, account.ServersColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNodes queries the nodes edge of a Account.
func (c *AccountClient) QueryNodes(a *Account) *NodeQuery {
	query := (&NodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.NodesTable, account.NodesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRegistrationToken queries the registration_token edge of a Account.
func (c *AccountClient) QueryRegistrationToken(a *Account) *RegistrationTokenQuery {
	query := (&RegistrationTokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(registrationtoken.Table, registrationtoken.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, account.RegistrationTokenTable, account.RegistrationTokenColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// Interceptors returns the client interceptors.
func (c *AccountClient) Interceptors() []Interceptor {
	return c.inters.Account
}

func (c *AccountClient) mutate(ctx context.Context, m *AccountMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AccountCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AccountDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Account mutation op: %q", m.Op())
	}
}

// NodeClient is a client for the Node schema.
type NodeClient struct {
	config
}

// NewNodeClient returns a client for the Node from the given config.
func NewNodeClient(c config) *NodeClient {
	return &NodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `node.Hooks(f(g(h())))`.
func (c *NodeClient) Use(hooks ...Hook) {
	c.hooks.Node = append(c.hooks.Node, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `node.Intercept(f(g(h())))`.
func (c *NodeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Node = append(c.inters.Node, interceptors...)
}

// Create returns a builder for creating a Node entity.
func (c *NodeClient) Create() *NodeCreate {
	mutation := newNodeMutation(c.config, OpCreate)
	return &NodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Node entities.
func (c *NodeClient) CreateBulk(builders ...*NodeCreate) *NodeCreateBulk {
	return &NodeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *NodeClient) MapCreateBulk(slice any, setFunc func(*NodeCreate, int)) *NodeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &NodeCreateBulk{err: fmt.Errorf("calling to NodeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*NodeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &NodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Node.
func (c *NodeClient) Update() *NodeUpdate {
	mutation := newNodeMutation(c.config, OpUpdate)
	return &NodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NodeClient) UpdateOne(n *Node) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNode(n))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NodeClient) UpdateOneID(id int) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNodeID(id))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Node.
func (c *NodeClient) Delete() *NodeDelete {
	mutation := newNodeMutation(c.config, OpDelete)
	return &NodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NodeClient) DeleteOne(n *Node) *NodeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *NodeClient) DeleteOneID(id int) *NodeDeleteOne {
	builder := c.Delete().Where(node.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NodeDeleteOne{builder}
}

// Query returns a query builder for Node.
func (c *NodeClient) Query() *NodeQuery {
	return &NodeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeNode},
		inters: c.Interceptors(),
	}
}

// Get returns a Node entity by its id.
func (c *NodeClient) Get(ctx context.Context, id int) (*Node, error) {
	return c.Query().Where(node.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NodeClient) GetX(ctx context.Context, id int) *Node {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a Node.
func (c *NodeClient) QueryAccount(n *Node) *AccountQuery {
	query := (&AccountClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, node.AccountTable, node.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryServer queries the server edge of a Node.
func (c *NodeClient) QueryServer(n *Node) *ServerQuery {
	query := (&ServerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, node.ServerTable, node.ServerColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NodeClient) Hooks() []Hook {
	return c.hooks.Node
}

// Interceptors returns the client interceptors.
func (c *NodeClient) Interceptors() []Interceptor {
	return c.inters.Node
}

func (c *NodeClient) mutate(ctx context.Context, m *NodeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&NodeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&NodeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&NodeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Node mutation op: %q", m.Op())
	}
}

// RegistrationTokenClient is a client for the RegistrationToken schema.
type RegistrationTokenClient struct {
	config
}

// NewRegistrationTokenClient returns a client for the RegistrationToken from the given config.
func NewRegistrationTokenClient(c config) *RegistrationTokenClient {
	return &RegistrationTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `registrationtoken.Hooks(f(g(h())))`.
func (c *RegistrationTokenClient) Use(hooks ...Hook) {
	c.hooks.RegistrationToken = append(c.hooks.RegistrationToken, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `registrationtoken.Intercept(f(g(h())))`.
func (c *RegistrationTokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.RegistrationToken = append(c.inters.RegistrationToken, interceptors...)
}

// Create returns a builder for creating a RegistrationToken entity.
func (c *RegistrationTokenClient) Create() *RegistrationTokenCreate {
	mutation := newRegistrationTokenMutation(c.config, OpCreate)
	return &RegistrationTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RegistrationToken entities.
func (c *RegistrationTokenClient) CreateBulk(builders ...*RegistrationTokenCreate) *RegistrationTokenCreateBulk {
	return &RegistrationTokenCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RegistrationTokenClient) MapCreateBulk(slice any, setFunc func(*RegistrationTokenCreate, int)) *RegistrationTokenCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RegistrationTokenCreateBulk{err: fmt.Errorf("calling to RegistrationTokenClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RegistrationTokenCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RegistrationTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RegistrationToken.
func (c *RegistrationTokenClient) Update() *RegistrationTokenUpdate {
	mutation := newRegistrationTokenMutation(c.config, OpUpdate)
	return &RegistrationTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RegistrationTokenClient) UpdateOne(rt *RegistrationToken) *RegistrationTokenUpdateOne {
	mutation := newRegistrationTokenMutation(c.config, OpUpdateOne, withRegistrationToken(rt))
	return &RegistrationTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RegistrationTokenClient) UpdateOneID(id int) *RegistrationTokenUpdateOne {
	mutation := newRegistrationTokenMutation(c.config, OpUpdateOne, withRegistrationTokenID(id))
	return &RegistrationTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RegistrationToken.
func (c *RegistrationTokenClient) Delete() *RegistrationTokenDelete {
	mutation := newRegistrationTokenMutation(c.config, OpDelete)
	return &RegistrationTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RegistrationTokenClient) DeleteOne(rt *RegistrationToken) *RegistrationTokenDeleteOne {
	return c.DeleteOneID(rt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RegistrationTokenClient) DeleteOneID(id int) *RegistrationTokenDeleteOne {
	builder := c.Delete().Where(registrationtoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RegistrationTokenDeleteOne{builder}
}

// Query returns a query builder for RegistrationToken.
func (c *RegistrationTokenClient) Query() *RegistrationTokenQuery {
	return &RegistrationTokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRegistrationToken},
		inters: c.Interceptors(),
	}
}

// Get returns a RegistrationToken entity by its id.
func (c *RegistrationTokenClient) Get(ctx context.Context, id int) (*RegistrationToken, error) {
	return c.Query().Where(registrationtoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RegistrationTokenClient) GetX(ctx context.Context, id int) *RegistrationToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRegistrations queries the registrations edge of a RegistrationToken.
func (c *RegistrationTokenClient) QueryRegistrations(rt *RegistrationToken) *AccountQuery {
	query := (&AccountClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(registrationtoken.Table, registrationtoken.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, registrationtoken.RegistrationsTable, registrationtoken.RegistrationsColumn),
		)
		fromV = sqlgraph.Neighbors(rt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RegistrationTokenClient) Hooks() []Hook {
	return c.hooks.RegistrationToken
}

// Interceptors returns the client interceptors.
func (c *RegistrationTokenClient) Interceptors() []Interceptor {
	return c.inters.RegistrationToken
}

func (c *RegistrationTokenClient) mutate(ctx context.Context, m *RegistrationTokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RegistrationTokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RegistrationTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RegistrationTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RegistrationTokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RegistrationToken mutation op: %q", m.Op())
	}
}

// ServerClient is a client for the Server schema.
type ServerClient struct {
	config
}

// NewServerClient returns a client for the Server from the given config.
func NewServerClient(c config) *ServerClient {
	return &ServerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `server.Hooks(f(g(h())))`.
func (c *ServerClient) Use(hooks ...Hook) {
	c.hooks.Server = append(c.hooks.Server, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `server.Intercept(f(g(h())))`.
func (c *ServerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Server = append(c.inters.Server, interceptors...)
}

// Create returns a builder for creating a Server entity.
func (c *ServerClient) Create() *ServerCreate {
	mutation := newServerMutation(c.config, OpCreate)
	return &ServerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Server entities.
func (c *ServerClient) CreateBulk(builders ...*ServerCreate) *ServerCreateBulk {
	return &ServerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ServerClient) MapCreateBulk(slice any, setFunc func(*ServerCreate, int)) *ServerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ServerCreateBulk{err: fmt.Errorf("calling to ServerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ServerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ServerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Server.
func (c *ServerClient) Update() *ServerUpdate {
	mutation := newServerMutation(c.config, OpUpdate)
	return &ServerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServerClient) UpdateOne(s *Server) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServer(s))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServerClient) UpdateOneID(id int) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServerID(id))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Server.
func (c *ServerClient) Delete() *ServerDelete {
	mutation := newServerMutation(c.config, OpDelete)
	return &ServerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServerClient) DeleteOne(s *Server) *ServerDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServerClient) DeleteOneID(id int) *ServerDeleteOne {
	builder := c.Delete().Where(server.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServerDeleteOne{builder}
}

// Query returns a query builder for Server.
func (c *ServerClient) Query() *ServerQuery {
	return &ServerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeServer},
		inters: c.Interceptors(),
	}
}

// Get returns a Server entity by its id.
func (c *ServerClient) Get(ctx context.Context, id int) (*Server, error) {
	return c.Query().Where(server.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServerClient) GetX(ctx context.Context, id int) *Server {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryNodes queries the nodes edge of a Server.
func (c *ServerClient) QueryNodes(s *Server) *NodeQuery {
	query := (&NodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, server.NodesTable, server.NodesColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccount queries the account edge of a Server.
func (c *ServerClient) QueryAccount(s *Server) *AccountQuery {
	query := (&AccountClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, server.AccountTable, server.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServerClient) Hooks() []Hook {
	return c.hooks.Server
}

// Interceptors returns the client interceptors.
func (c *ServerClient) Interceptors() []Interceptor {
	return c.inters.Server
}

func (c *ServerClient) mutate(ctx context.Context, m *ServerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ServerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ServerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ServerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Server mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Account, Node, RegistrationToken, Server []ent.Hook
	}
	inters struct {
		Account, Node, RegistrationToken, Server []ent.Interceptor
	}
)
