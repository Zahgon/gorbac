package gorbac

import (
	"sync"
)

// Roles is a map
type Roles[T comparable] map[T]Role[T]

// NewStdRole is the default role factory function.
// It matches the declaration to RoleFactoryFunc.
func NewRole[T comparable](id T) Role[T] { _ = "STUB: not implemented"; return nil }

// StdRole is the default role implement.
// You can combine this struct into your own Role implement.
// T is the type of ID
type Role[T comparable] struct {
	sync.RWMutex
	// ID is the serialisable identity of role
	ID          T `json:"id"`
	permissions Permissions[T]
}

// Assign a permission to the role.
func (role *Role[T]) Assign(p Permission[T]) error { _ = "STUB: not implemented"; return nil }

// Permit returns true if the role has specific permission.
func (role *Role[T]) Permit(p Permission[T]) (ok bool) { _ = "STUB: not implemented"; return false }

// Revoke the specific permission.
func (role *Role[T]) Revoke(p Permission[T]) error { _ = "STUB: not implemented"; return nil }

// Permissions returns all permissions into a slice.
func (role *Role[T]) Permissions() []Permission[T] { _ = "STUB: not implemented"; return nil }
