// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/migrate/entv2/conversion"
	"github.com/facebook/ent/entc/integration/migrate/entv2/predicate"
	"github.com/facebook/ent/schema/field"
)

// ConversionUpdate is the builder for updating Conversion entities.
type ConversionUpdate struct {
	config
	hooks    []Hook
	mutation *ConversionMutation
}

// Where adds a new predicate for the builder.
func (cu *ConversionUpdate) Where(ps ...predicate.Conversion) *ConversionUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetName sets the name field.
func (cu *ConversionUpdate) SetName(s string) *ConversionUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetInt8ToString sets the int8_to_string field.
func (cu *ConversionUpdate) SetInt8ToString(s string) *ConversionUpdate {
	cu.mutation.SetInt8ToString(s)
	return cu
}

// SetUint8ToString sets the uint8_to_string field.
func (cu *ConversionUpdate) SetUint8ToString(s string) *ConversionUpdate {
	cu.mutation.SetUint8ToString(s)
	return cu
}

// SetInt16ToString sets the int16_to_string field.
func (cu *ConversionUpdate) SetInt16ToString(s string) *ConversionUpdate {
	cu.mutation.SetInt16ToString(s)
	return cu
}

// SetUint16ToString sets the uint16_to_string field.
func (cu *ConversionUpdate) SetUint16ToString(s string) *ConversionUpdate {
	cu.mutation.SetUint16ToString(s)
	return cu
}

// SetInt32ToString sets the int32_to_string field.
func (cu *ConversionUpdate) SetInt32ToString(s string) *ConversionUpdate {
	cu.mutation.SetInt32ToString(s)
	return cu
}

// SetUint32ToString sets the uint32_to_string field.
func (cu *ConversionUpdate) SetUint32ToString(s string) *ConversionUpdate {
	cu.mutation.SetUint32ToString(s)
	return cu
}

// SetInt64ToString sets the int64_to_string field.
func (cu *ConversionUpdate) SetInt64ToString(s string) *ConversionUpdate {
	cu.mutation.SetInt64ToString(s)
	return cu
}

// SetUint64ToString sets the uint64_to_string field.
func (cu *ConversionUpdate) SetUint64ToString(s string) *ConversionUpdate {
	cu.mutation.SetUint64ToString(s)
	return cu
}

// Mutation returns the ConversionMutation object of the builder.
func (cu *ConversionUpdate) Mutation() *ConversionMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConversionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConversionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConversionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConversionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConversionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConversionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   conversion.Table,
			Columns: conversion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: conversion.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldName,
		})
	}
	if value, ok := cu.mutation.Int8ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldInt8ToString,
		})
	}
	if value, ok := cu.mutation.Uint8ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldUint8ToString,
		})
	}
	if value, ok := cu.mutation.Int16ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldInt16ToString,
		})
	}
	if value, ok := cu.mutation.Uint16ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldUint16ToString,
		})
	}
	if value, ok := cu.mutation.Int32ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldInt32ToString,
		})
	}
	if value, ok := cu.mutation.Uint32ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldUint32ToString,
		})
	}
	if value, ok := cu.mutation.Int64ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldInt64ToString,
		})
	}
	if value, ok := cu.mutation.Uint64ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldUint64ToString,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversion.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ConversionUpdateOne is the builder for updating a single Conversion entity.
type ConversionUpdateOne struct {
	config
	hooks    []Hook
	mutation *ConversionMutation
}

// SetName sets the name field.
func (cuo *ConversionUpdateOne) SetName(s string) *ConversionUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetInt8ToString sets the int8_to_string field.
func (cuo *ConversionUpdateOne) SetInt8ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetInt8ToString(s)
	return cuo
}

// SetUint8ToString sets the uint8_to_string field.
func (cuo *ConversionUpdateOne) SetUint8ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetUint8ToString(s)
	return cuo
}

// SetInt16ToString sets the int16_to_string field.
func (cuo *ConversionUpdateOne) SetInt16ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetInt16ToString(s)
	return cuo
}

// SetUint16ToString sets the uint16_to_string field.
func (cuo *ConversionUpdateOne) SetUint16ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetUint16ToString(s)
	return cuo
}

// SetInt32ToString sets the int32_to_string field.
func (cuo *ConversionUpdateOne) SetInt32ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetInt32ToString(s)
	return cuo
}

// SetUint32ToString sets the uint32_to_string field.
func (cuo *ConversionUpdateOne) SetUint32ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetUint32ToString(s)
	return cuo
}

// SetInt64ToString sets the int64_to_string field.
func (cuo *ConversionUpdateOne) SetInt64ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetInt64ToString(s)
	return cuo
}

// SetUint64ToString sets the uint64_to_string field.
func (cuo *ConversionUpdateOne) SetUint64ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetUint64ToString(s)
	return cuo
}

// Mutation returns the ConversionMutation object of the builder.
func (cuo *ConversionUpdateOne) Mutation() *ConversionMutation {
	return cuo.mutation
}

// Save executes the query and returns the updated entity.
func (cuo *ConversionUpdateOne) Save(ctx context.Context) (*Conversion, error) {
	var (
		err  error
		node *Conversion
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConversionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConversionUpdateOne) SaveX(ctx context.Context) *Conversion {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConversionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConversionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConversionUpdateOne) sqlSave(ctx context.Context) (_node *Conversion, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   conversion.Table,
			Columns: conversion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: conversion.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Conversion.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldName,
		})
	}
	if value, ok := cuo.mutation.Int8ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldInt8ToString,
		})
	}
	if value, ok := cuo.mutation.Uint8ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldUint8ToString,
		})
	}
	if value, ok := cuo.mutation.Int16ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldInt16ToString,
		})
	}
	if value, ok := cuo.mutation.Uint16ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldUint16ToString,
		})
	}
	if value, ok := cuo.mutation.Int32ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldInt32ToString,
		})
	}
	if value, ok := cuo.mutation.Uint32ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldUint32ToString,
		})
	}
	if value, ok := cuo.mutation.Int64ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldInt64ToString,
		})
	}
	if value, ok := cuo.mutation.Uint64ToString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversion.FieldUint64ToString,
		})
	}
	_node = &Conversion{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversion.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
