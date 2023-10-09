// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/hashequal"
)

// HashEqualCreate is the builder for creating a HashEqual entity.
type HashEqualCreate struct {
	config
	mutation *HashEqualMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOrigin sets the "origin" field.
func (hec *HashEqualCreate) SetOrigin(s string) *HashEqualCreate {
	hec.mutation.SetOrigin(s)
	return hec
}

// SetCollector sets the "collector" field.
func (hec *HashEqualCreate) SetCollector(s string) *HashEqualCreate {
	hec.mutation.SetCollector(s)
	return hec
}

// SetJustification sets the "justification" field.
func (hec *HashEqualCreate) SetJustification(s string) *HashEqualCreate {
	hec.mutation.SetJustification(s)
	return hec
}

// AddArtifactIDs adds the "artifacts" edge to the Artifact entity by IDs.
func (hec *HashEqualCreate) AddArtifactIDs(ids ...int) *HashEqualCreate {
	hec.mutation.AddArtifactIDs(ids...)
	return hec
}

// AddArtifacts adds the "artifacts" edges to the Artifact entity.
func (hec *HashEqualCreate) AddArtifacts(a ...*Artifact) *HashEqualCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return hec.AddArtifactIDs(ids...)
}

// Mutation returns the HashEqualMutation object of the builder.
func (hec *HashEqualCreate) Mutation() *HashEqualMutation {
	return hec.mutation
}

// Save creates the HashEqual in the database.
func (hec *HashEqualCreate) Save(ctx context.Context) (*HashEqual, error) {
	return withHooks(ctx, hec.sqlSave, hec.mutation, hec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hec *HashEqualCreate) SaveX(ctx context.Context) *HashEqual {
	v, err := hec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hec *HashEqualCreate) Exec(ctx context.Context) error {
	_, err := hec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hec *HashEqualCreate) ExecX(ctx context.Context) {
	if err := hec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hec *HashEqualCreate) check() error {
	if _, ok := hec.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "HashEqual.origin"`)}
	}
	if _, ok := hec.mutation.Collector(); !ok {
		return &ValidationError{Name: "collector", err: errors.New(`ent: missing required field "HashEqual.collector"`)}
	}
	if _, ok := hec.mutation.Justification(); !ok {
		return &ValidationError{Name: "justification", err: errors.New(`ent: missing required field "HashEqual.justification"`)}
	}
	if len(hec.mutation.ArtifactsIDs()) == 0 {
		return &ValidationError{Name: "artifacts", err: errors.New(`ent: missing required edge "HashEqual.artifacts"`)}
	}
	return nil
}

func (hec *HashEqualCreate) sqlSave(ctx context.Context) (*HashEqual, error) {
	if err := hec.check(); err != nil {
		return nil, err
	}
	_node, _spec := hec.createSpec()
	if err := sqlgraph.CreateNode(ctx, hec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hec.mutation.id = &_node.ID
	hec.mutation.done = true
	return _node, nil
}

func (hec *HashEqualCreate) createSpec() (*HashEqual, *sqlgraph.CreateSpec) {
	var (
		_node = &HashEqual{config: hec.config}
		_spec = sqlgraph.NewCreateSpec(hashequal.Table, sqlgraph.NewFieldSpec(hashequal.FieldID, field.TypeInt))
	)
	_spec.OnConflict = hec.conflict
	if value, ok := hec.mutation.Origin(); ok {
		_spec.SetField(hashequal.FieldOrigin, field.TypeString, value)
		_node.Origin = value
	}
	if value, ok := hec.mutation.Collector(); ok {
		_spec.SetField(hashequal.FieldCollector, field.TypeString, value)
		_node.Collector = value
	}
	if value, ok := hec.mutation.Justification(); ok {
		_spec.SetField(hashequal.FieldJustification, field.TypeString, value)
		_node.Justification = value
	}
	if nodes := hec.mutation.ArtifactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hashequal.ArtifactsTable,
			Columns: hashequal.ArtifactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.HashEqual.Create().
//		SetOrigin(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HashEqualUpsert) {
//			SetOrigin(v+v).
//		}).
//		Exec(ctx)
func (hec *HashEqualCreate) OnConflict(opts ...sql.ConflictOption) *HashEqualUpsertOne {
	hec.conflict = opts
	return &HashEqualUpsertOne{
		create: hec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.HashEqual.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hec *HashEqualCreate) OnConflictColumns(columns ...string) *HashEqualUpsertOne {
	hec.conflict = append(hec.conflict, sql.ConflictColumns(columns...))
	return &HashEqualUpsertOne{
		create: hec,
	}
}

type (
	// HashEqualUpsertOne is the builder for "upsert"-ing
	//  one HashEqual node.
	HashEqualUpsertOne struct {
		create *HashEqualCreate
	}

	// HashEqualUpsert is the "OnConflict" setter.
	HashEqualUpsert struct {
		*sql.UpdateSet
	}
)

// SetOrigin sets the "origin" field.
func (u *HashEqualUpsert) SetOrigin(v string) *HashEqualUpsert {
	u.Set(hashequal.FieldOrigin, v)
	return u
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *HashEqualUpsert) UpdateOrigin() *HashEqualUpsert {
	u.SetExcluded(hashequal.FieldOrigin)
	return u
}

// SetCollector sets the "collector" field.
func (u *HashEqualUpsert) SetCollector(v string) *HashEqualUpsert {
	u.Set(hashequal.FieldCollector, v)
	return u
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *HashEqualUpsert) UpdateCollector() *HashEqualUpsert {
	u.SetExcluded(hashequal.FieldCollector)
	return u
}

// SetJustification sets the "justification" field.
func (u *HashEqualUpsert) SetJustification(v string) *HashEqualUpsert {
	u.Set(hashequal.FieldJustification, v)
	return u
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *HashEqualUpsert) UpdateJustification() *HashEqualUpsert {
	u.SetExcluded(hashequal.FieldJustification)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.HashEqual.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *HashEqualUpsertOne) UpdateNewValues() *HashEqualUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.HashEqual.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *HashEqualUpsertOne) Ignore() *HashEqualUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HashEqualUpsertOne) DoNothing() *HashEqualUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HashEqualCreate.OnConflict
// documentation for more info.
func (u *HashEqualUpsertOne) Update(set func(*HashEqualUpsert)) *HashEqualUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HashEqualUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrigin sets the "origin" field.
func (u *HashEqualUpsertOne) SetOrigin(v string) *HashEqualUpsertOne {
	return u.Update(func(s *HashEqualUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *HashEqualUpsertOne) UpdateOrigin() *HashEqualUpsertOne {
	return u.Update(func(s *HashEqualUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *HashEqualUpsertOne) SetCollector(v string) *HashEqualUpsertOne {
	return u.Update(func(s *HashEqualUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *HashEqualUpsertOne) UpdateCollector() *HashEqualUpsertOne {
	return u.Update(func(s *HashEqualUpsert) {
		s.UpdateCollector()
	})
}

// SetJustification sets the "justification" field.
func (u *HashEqualUpsertOne) SetJustification(v string) *HashEqualUpsertOne {
	return u.Update(func(s *HashEqualUpsert) {
		s.SetJustification(v)
	})
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *HashEqualUpsertOne) UpdateJustification() *HashEqualUpsertOne {
	return u.Update(func(s *HashEqualUpsert) {
		s.UpdateJustification()
	})
}

// Exec executes the query.
func (u *HashEqualUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HashEqualCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HashEqualUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *HashEqualUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *HashEqualUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// HashEqualCreateBulk is the builder for creating many HashEqual entities in bulk.
type HashEqualCreateBulk struct {
	config
	err      error
	builders []*HashEqualCreate
	conflict []sql.ConflictOption
}

// Save creates the HashEqual entities in the database.
func (hecb *HashEqualCreateBulk) Save(ctx context.Context) ([]*HashEqual, error) {
	if hecb.err != nil {
		return nil, hecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hecb.builders))
	nodes := make([]*HashEqual, len(hecb.builders))
	mutators := make([]Mutator, len(hecb.builders))
	for i := range hecb.builders {
		func(i int, root context.Context) {
			builder := hecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HashEqualMutation)
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
					_, err = mutators[i+1].Mutate(root, hecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = hecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hecb *HashEqualCreateBulk) SaveX(ctx context.Context) []*HashEqual {
	v, err := hecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hecb *HashEqualCreateBulk) Exec(ctx context.Context) error {
	_, err := hecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hecb *HashEqualCreateBulk) ExecX(ctx context.Context) {
	if err := hecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.HashEqual.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HashEqualUpsert) {
//			SetOrigin(v+v).
//		}).
//		Exec(ctx)
func (hecb *HashEqualCreateBulk) OnConflict(opts ...sql.ConflictOption) *HashEqualUpsertBulk {
	hecb.conflict = opts
	return &HashEqualUpsertBulk{
		create: hecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.HashEqual.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hecb *HashEqualCreateBulk) OnConflictColumns(columns ...string) *HashEqualUpsertBulk {
	hecb.conflict = append(hecb.conflict, sql.ConflictColumns(columns...))
	return &HashEqualUpsertBulk{
		create: hecb,
	}
}

// HashEqualUpsertBulk is the builder for "upsert"-ing
// a bulk of HashEqual nodes.
type HashEqualUpsertBulk struct {
	create *HashEqualCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.HashEqual.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *HashEqualUpsertBulk) UpdateNewValues() *HashEqualUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.HashEqual.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *HashEqualUpsertBulk) Ignore() *HashEqualUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HashEqualUpsertBulk) DoNothing() *HashEqualUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HashEqualCreateBulk.OnConflict
// documentation for more info.
func (u *HashEqualUpsertBulk) Update(set func(*HashEqualUpsert)) *HashEqualUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HashEqualUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrigin sets the "origin" field.
func (u *HashEqualUpsertBulk) SetOrigin(v string) *HashEqualUpsertBulk {
	return u.Update(func(s *HashEqualUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *HashEqualUpsertBulk) UpdateOrigin() *HashEqualUpsertBulk {
	return u.Update(func(s *HashEqualUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *HashEqualUpsertBulk) SetCollector(v string) *HashEqualUpsertBulk {
	return u.Update(func(s *HashEqualUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *HashEqualUpsertBulk) UpdateCollector() *HashEqualUpsertBulk {
	return u.Update(func(s *HashEqualUpsert) {
		s.UpdateCollector()
	})
}

// SetJustification sets the "justification" field.
func (u *HashEqualUpsertBulk) SetJustification(v string) *HashEqualUpsertBulk {
	return u.Update(func(s *HashEqualUpsert) {
		s.SetJustification(v)
	})
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *HashEqualUpsertBulk) UpdateJustification() *HashEqualUpsertBulk {
	return u.Update(func(s *HashEqualUpsert) {
		s.UpdateJustification()
	})
}

// Exec executes the query.
func (u *HashEqualUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the HashEqualCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HashEqualCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HashEqualUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}