// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sample/ent/binaryfile"
	"sample/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BinaryFileCreate is the builder for creating a BinaryFile entity.
type BinaryFileCreate struct {
	config
	mutation *BinaryFileMutation
	hooks    []Hook
}

// SetFilename sets the "filename" field.
func (bfc *BinaryFileCreate) SetFilename(s string) *BinaryFileCreate {
	bfc.mutation.SetFilename(s)
	return bfc
}

// SetBody sets the "body" field.
func (bfc *BinaryFileCreate) SetBody(b []byte) *BinaryFileCreate {
	bfc.mutation.SetBody(b)
	return bfc
}

// SetCreatedAt sets the "created_at" field.
func (bfc *BinaryFileCreate) SetCreatedAt(t time.Time) *BinaryFileCreate {
	bfc.mutation.SetCreatedAt(t)
	return bfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bfc *BinaryFileCreate) SetNillableCreatedAt(t *time.Time) *BinaryFileCreate {
	if t != nil {
		bfc.SetCreatedAt(*t)
	}
	return bfc
}

// SetUpdatedAt sets the "updated_at" field.
func (bfc *BinaryFileCreate) SetUpdatedAt(t time.Time) *BinaryFileCreate {
	bfc.mutation.SetUpdatedAt(t)
	return bfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bfc *BinaryFileCreate) SetNillableUpdatedAt(t *time.Time) *BinaryFileCreate {
	if t != nil {
		bfc.SetUpdatedAt(*t)
	}
	return bfc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (bfc *BinaryFileCreate) SetOwnerID(id int) *BinaryFileCreate {
	bfc.mutation.SetOwnerID(id)
	return bfc
}

// SetOwner sets the "owner" edge to the User entity.
func (bfc *BinaryFileCreate) SetOwner(u *User) *BinaryFileCreate {
	return bfc.SetOwnerID(u.ID)
}

// Mutation returns the BinaryFileMutation object of the builder.
func (bfc *BinaryFileCreate) Mutation() *BinaryFileMutation {
	return bfc.mutation
}

// Save creates the BinaryFile in the database.
func (bfc *BinaryFileCreate) Save(ctx context.Context) (*BinaryFile, error) {
	var (
		err  error
		node *BinaryFile
	)
	bfc.defaults()
	if len(bfc.hooks) == 0 {
		if err = bfc.check(); err != nil {
			return nil, err
		}
		node, err = bfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BinaryFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bfc.check(); err != nil {
				return nil, err
			}
			bfc.mutation = mutation
			if node, err = bfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bfc.hooks) - 1; i >= 0; i-- {
			if bfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bfc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bfc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bfc *BinaryFileCreate) SaveX(ctx context.Context) *BinaryFile {
	v, err := bfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bfc *BinaryFileCreate) Exec(ctx context.Context) error {
	_, err := bfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bfc *BinaryFileCreate) ExecX(ctx context.Context) {
	if err := bfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bfc *BinaryFileCreate) defaults() {
	if _, ok := bfc.mutation.CreatedAt(); !ok {
		v := binaryfile.DefaultCreatedAt()
		bfc.mutation.SetCreatedAt(v)
	}
	if _, ok := bfc.mutation.UpdatedAt(); !ok {
		v := binaryfile.DefaultUpdatedAt()
		bfc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bfc *BinaryFileCreate) check() error {
	if _, ok := bfc.mutation.Filename(); !ok {
		return &ValidationError{Name: "filename", err: errors.New(`ent: missing required field "filename"`)}
	}
	if v, ok := bfc.mutation.Filename(); ok {
		if err := binaryfile.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf(`ent: validator failed for field "filename": %w`, err)}
		}
	}
	if _, ok := bfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := bfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := bfc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New("ent: missing required edge \"owner\"")}
	}
	return nil
}

func (bfc *BinaryFileCreate) sqlSave(ctx context.Context) (*BinaryFile, error) {
	_node, _spec := bfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bfc *BinaryFileCreate) createSpec() (*BinaryFile, *sqlgraph.CreateSpec) {
	var (
		_node = &BinaryFile{config: bfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: binaryfile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: binaryfile.FieldID,
			},
		}
	)
	if value, ok := bfc.mutation.Filename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: binaryfile.FieldFilename,
		})
		_node.Filename = value
	}
	if value, ok := bfc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: binaryfile.FieldBody,
		})
		_node.Body = &value
	}
	if value, ok := bfc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: binaryfile.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bfc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: binaryfile.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := bfc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   binaryfile.OwnerTable,
			Columns: []string{binaryfile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_owned = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BinaryFileCreateBulk is the builder for creating many BinaryFile entities in bulk.
type BinaryFileCreateBulk struct {
	config
	builders []*BinaryFileCreate
}

// Save creates the BinaryFile entities in the database.
func (bfcb *BinaryFileCreateBulk) Save(ctx context.Context) ([]*BinaryFile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bfcb.builders))
	nodes := make([]*BinaryFile, len(bfcb.builders))
	mutators := make([]Mutator, len(bfcb.builders))
	for i := range bfcb.builders {
		func(i int, root context.Context) {
			builder := bfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BinaryFileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bfcb *BinaryFileCreateBulk) SaveX(ctx context.Context) []*BinaryFile {
	v, err := bfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bfcb *BinaryFileCreateBulk) Exec(ctx context.Context) error {
	_, err := bfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bfcb *BinaryFileCreateBulk) ExecX(ctx context.Context) {
	if err := bfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
