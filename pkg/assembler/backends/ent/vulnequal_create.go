// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/vulnequal"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/vulnerabilityid"
)

// VulnEqualCreate is the builder for creating a VulnEqual entity.
type VulnEqualCreate struct {
	config
	mutation *VulnEqualMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetJustification sets the "justification" field.
func (vec *VulnEqualCreate) SetJustification(s string) *VulnEqualCreate {
	vec.mutation.SetJustification(s)
	return vec
}

// SetOrigin sets the "origin" field.
func (vec *VulnEqualCreate) SetOrigin(s string) *VulnEqualCreate {
	vec.mutation.SetOrigin(s)
	return vec
}

// SetCollector sets the "collector" field.
func (vec *VulnEqualCreate) SetCollector(s string) *VulnEqualCreate {
	vec.mutation.SetCollector(s)
	return vec
}

// AddVulnerabilityIDIDs adds the "vulnerability_ids" edge to the VulnerabilityID entity by IDs.
func (vec *VulnEqualCreate) AddVulnerabilityIDIDs(ids ...int) *VulnEqualCreate {
	vec.mutation.AddVulnerabilityIDIDs(ids...)
	return vec
}

// AddVulnerabilityIds adds the "vulnerability_ids" edges to the VulnerabilityID entity.
func (vec *VulnEqualCreate) AddVulnerabilityIds(v ...*VulnerabilityID) *VulnEqualCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vec.AddVulnerabilityIDIDs(ids...)
}

// Mutation returns the VulnEqualMutation object of the builder.
func (vec *VulnEqualCreate) Mutation() *VulnEqualMutation {
	return vec.mutation
}

// Save creates the VulnEqual in the database.
func (vec *VulnEqualCreate) Save(ctx context.Context) (*VulnEqual, error) {
	return withHooks(ctx, vec.sqlSave, vec.mutation, vec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vec *VulnEqualCreate) SaveX(ctx context.Context) *VulnEqual {
	v, err := vec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vec *VulnEqualCreate) Exec(ctx context.Context) error {
	_, err := vec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vec *VulnEqualCreate) ExecX(ctx context.Context) {
	if err := vec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vec *VulnEqualCreate) check() error {
	if _, ok := vec.mutation.Justification(); !ok {
		return &ValidationError{Name: "justification", err: errors.New(`ent: missing required field "VulnEqual.justification"`)}
	}
	if _, ok := vec.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "VulnEqual.origin"`)}
	}
	if _, ok := vec.mutation.Collector(); !ok {
		return &ValidationError{Name: "collector", err: errors.New(`ent: missing required field "VulnEqual.collector"`)}
	}
	if len(vec.mutation.VulnerabilityIdsIDs()) == 0 {
		return &ValidationError{Name: "vulnerability_ids", err: errors.New(`ent: missing required edge "VulnEqual.vulnerability_ids"`)}
	}
	return nil
}

func (vec *VulnEqualCreate) sqlSave(ctx context.Context) (*VulnEqual, error) {
	if err := vec.check(); err != nil {
		return nil, err
	}
	_node, _spec := vec.createSpec()
	if err := sqlgraph.CreateNode(ctx, vec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vec.mutation.id = &_node.ID
	vec.mutation.done = true
	return _node, nil
}

func (vec *VulnEqualCreate) createSpec() (*VulnEqual, *sqlgraph.CreateSpec) {
	var (
		_node = &VulnEqual{config: vec.config}
		_spec = sqlgraph.NewCreateSpec(vulnequal.Table, sqlgraph.NewFieldSpec(vulnequal.FieldID, field.TypeInt))
	)
	_spec.OnConflict = vec.conflict
	if value, ok := vec.mutation.Justification(); ok {
		_spec.SetField(vulnequal.FieldJustification, field.TypeString, value)
		_node.Justification = value
	}
	if value, ok := vec.mutation.Origin(); ok {
		_spec.SetField(vulnequal.FieldOrigin, field.TypeString, value)
		_node.Origin = value
	}
	if value, ok := vec.mutation.Collector(); ok {
		_spec.SetField(vulnequal.FieldCollector, field.TypeString, value)
		_node.Collector = value
	}
	if nodes := vec.mutation.VulnerabilityIdsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   vulnequal.VulnerabilityIdsTable,
			Columns: vulnequal.VulnerabilityIdsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vulnerabilityid.FieldID, field.TypeInt),
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
//	client.VulnEqual.Create().
//		SetJustification(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VulnEqualUpsert) {
//			SetJustification(v+v).
//		}).
//		Exec(ctx)
func (vec *VulnEqualCreate) OnConflict(opts ...sql.ConflictOption) *VulnEqualUpsertOne {
	vec.conflict = opts
	return &VulnEqualUpsertOne{
		create: vec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.VulnEqual.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vec *VulnEqualCreate) OnConflictColumns(columns ...string) *VulnEqualUpsertOne {
	vec.conflict = append(vec.conflict, sql.ConflictColumns(columns...))
	return &VulnEqualUpsertOne{
		create: vec,
	}
}

type (
	// VulnEqualUpsertOne is the builder for "upsert"-ing
	//  one VulnEqual node.
	VulnEqualUpsertOne struct {
		create *VulnEqualCreate
	}

	// VulnEqualUpsert is the "OnConflict" setter.
	VulnEqualUpsert struct {
		*sql.UpdateSet
	}
)

// SetJustification sets the "justification" field.
func (u *VulnEqualUpsert) SetJustification(v string) *VulnEqualUpsert {
	u.Set(vulnequal.FieldJustification, v)
	return u
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *VulnEqualUpsert) UpdateJustification() *VulnEqualUpsert {
	u.SetExcluded(vulnequal.FieldJustification)
	return u
}

// SetOrigin sets the "origin" field.
func (u *VulnEqualUpsert) SetOrigin(v string) *VulnEqualUpsert {
	u.Set(vulnequal.FieldOrigin, v)
	return u
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *VulnEqualUpsert) UpdateOrigin() *VulnEqualUpsert {
	u.SetExcluded(vulnequal.FieldOrigin)
	return u
}

// SetCollector sets the "collector" field.
func (u *VulnEqualUpsert) SetCollector(v string) *VulnEqualUpsert {
	u.Set(vulnequal.FieldCollector, v)
	return u
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *VulnEqualUpsert) UpdateCollector() *VulnEqualUpsert {
	u.SetExcluded(vulnequal.FieldCollector)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.VulnEqual.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *VulnEqualUpsertOne) UpdateNewValues() *VulnEqualUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.VulnEqual.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *VulnEqualUpsertOne) Ignore() *VulnEqualUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VulnEqualUpsertOne) DoNothing() *VulnEqualUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VulnEqualCreate.OnConflict
// documentation for more info.
func (u *VulnEqualUpsertOne) Update(set func(*VulnEqualUpsert)) *VulnEqualUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VulnEqualUpsert{UpdateSet: update})
	}))
	return u
}

// SetJustification sets the "justification" field.
func (u *VulnEqualUpsertOne) SetJustification(v string) *VulnEqualUpsertOne {
	return u.Update(func(s *VulnEqualUpsert) {
		s.SetJustification(v)
	})
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *VulnEqualUpsertOne) UpdateJustification() *VulnEqualUpsertOne {
	return u.Update(func(s *VulnEqualUpsert) {
		s.UpdateJustification()
	})
}

// SetOrigin sets the "origin" field.
func (u *VulnEqualUpsertOne) SetOrigin(v string) *VulnEqualUpsertOne {
	return u.Update(func(s *VulnEqualUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *VulnEqualUpsertOne) UpdateOrigin() *VulnEqualUpsertOne {
	return u.Update(func(s *VulnEqualUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *VulnEqualUpsertOne) SetCollector(v string) *VulnEqualUpsertOne {
	return u.Update(func(s *VulnEqualUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *VulnEqualUpsertOne) UpdateCollector() *VulnEqualUpsertOne {
	return u.Update(func(s *VulnEqualUpsert) {
		s.UpdateCollector()
	})
}

// Exec executes the query.
func (u *VulnEqualUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VulnEqualCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VulnEqualUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *VulnEqualUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *VulnEqualUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// VulnEqualCreateBulk is the builder for creating many VulnEqual entities in bulk.
type VulnEqualCreateBulk struct {
	config
	err      error
	builders []*VulnEqualCreate
	conflict []sql.ConflictOption
}

// Save creates the VulnEqual entities in the database.
func (vecb *VulnEqualCreateBulk) Save(ctx context.Context) ([]*VulnEqual, error) {
	if vecb.err != nil {
		return nil, vecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vecb.builders))
	nodes := make([]*VulnEqual, len(vecb.builders))
	mutators := make([]Mutator, len(vecb.builders))
	for i := range vecb.builders {
		func(i int, root context.Context) {
			builder := vecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VulnEqualMutation)
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
					_, err = mutators[i+1].Mutate(root, vecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = vecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vecb *VulnEqualCreateBulk) SaveX(ctx context.Context) []*VulnEqual {
	v, err := vecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vecb *VulnEqualCreateBulk) Exec(ctx context.Context) error {
	_, err := vecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vecb *VulnEqualCreateBulk) ExecX(ctx context.Context) {
	if err := vecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.VulnEqual.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VulnEqualUpsert) {
//			SetJustification(v+v).
//		}).
//		Exec(ctx)
func (vecb *VulnEqualCreateBulk) OnConflict(opts ...sql.ConflictOption) *VulnEqualUpsertBulk {
	vecb.conflict = opts
	return &VulnEqualUpsertBulk{
		create: vecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.VulnEqual.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vecb *VulnEqualCreateBulk) OnConflictColumns(columns ...string) *VulnEqualUpsertBulk {
	vecb.conflict = append(vecb.conflict, sql.ConflictColumns(columns...))
	return &VulnEqualUpsertBulk{
		create: vecb,
	}
}

// VulnEqualUpsertBulk is the builder for "upsert"-ing
// a bulk of VulnEqual nodes.
type VulnEqualUpsertBulk struct {
	create *VulnEqualCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.VulnEqual.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *VulnEqualUpsertBulk) UpdateNewValues() *VulnEqualUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.VulnEqual.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *VulnEqualUpsertBulk) Ignore() *VulnEqualUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VulnEqualUpsertBulk) DoNothing() *VulnEqualUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VulnEqualCreateBulk.OnConflict
// documentation for more info.
func (u *VulnEqualUpsertBulk) Update(set func(*VulnEqualUpsert)) *VulnEqualUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VulnEqualUpsert{UpdateSet: update})
	}))
	return u
}

// SetJustification sets the "justification" field.
func (u *VulnEqualUpsertBulk) SetJustification(v string) *VulnEqualUpsertBulk {
	return u.Update(func(s *VulnEqualUpsert) {
		s.SetJustification(v)
	})
}

// UpdateJustification sets the "justification" field to the value that was provided on create.
func (u *VulnEqualUpsertBulk) UpdateJustification() *VulnEqualUpsertBulk {
	return u.Update(func(s *VulnEqualUpsert) {
		s.UpdateJustification()
	})
}

// SetOrigin sets the "origin" field.
func (u *VulnEqualUpsertBulk) SetOrigin(v string) *VulnEqualUpsertBulk {
	return u.Update(func(s *VulnEqualUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *VulnEqualUpsertBulk) UpdateOrigin() *VulnEqualUpsertBulk {
	return u.Update(func(s *VulnEqualUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *VulnEqualUpsertBulk) SetCollector(v string) *VulnEqualUpsertBulk {
	return u.Update(func(s *VulnEqualUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *VulnEqualUpsertBulk) UpdateCollector() *VulnEqualUpsertBulk {
	return u.Update(func(s *VulnEqualUpsert) {
		s.UpdateCollector()
	})
}

// Exec executes the query.
func (u *VulnEqualUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the VulnEqualCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VulnEqualCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VulnEqualUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
