package passport

import (
	"cmp"
	"fmt"
)

type SlicedValue[T cmp.Ordered] struct {
	err error
	key string
	val []T
}

func NewSlice[T cmp.Ordered](k string, v []T) *SlicedValue[T] {
	return &SlicedValue[T]{
		key: k,
		val: v,
	}
}

func (c *SlicedValue[T]) validate(ok bool, layout string, args ...any) *SlicedValue[T] {
	if c.err != nil || ok {
		return c
	}
	c.err = fmt.Errorf(layout, args...)
	return c
}

func (c *SlicedValue[T]) Err() error {
	return c.err
}

func (c *SlicedValue[T]) Required() *SlicedValue[T] {
	return c.validate(len(c.val) > 0, "%s is required", c.key)
}

func (c *SlicedValue[T]) LenEq(v int) *SlicedValue[T] {
	return c.validate(len(c.val) == v, "size of %s should equal %v", c.key, v)
}

func (c *SlicedValue[T]) LenGt(v int) *SlicedValue[T] {
	return c.validate(len(c.val) > v, "size of %s should great than %v", c.key, v)
}

func (c *SlicedValue[T]) LenGte(v int) *SlicedValue[T] {
	return c.validate(len(c.val) >= v, "size of %s should great or equal than %v", c.key, v)
}

func (c *SlicedValue[T]) LenLt(v int) *SlicedValue[T] {
	return c.validate(len(c.val) < v, "size of %s should less than %v", c.key, v)
}

func (c *SlicedValue[T]) LenLte(v int) *SlicedValue[T] {
	return c.validate(len(c.val) <= v, "size of %s should less or equal than %v", c.key, v)
}

func (c *SlicedValue[T]) Include(v T) *SlicedValue[T] {
	return c.validate(contains(c.val, v), "%s should contains %v", c.key, v)
}

func (c *SlicedValue[T]) Exclude(v T) *SlicedValue[T] {
	return c.validate(!contains(c.val, v), "%s should not contains %v", c.key, v)
}