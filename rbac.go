/*
Package gorbac provides a lightweight role-based access
control implementation in Golang.

For the purposes of this package:

  - an identity has one or more roles.
  - a role requests access to a permission.
  - a permission is given to a role.

Thus, RBAC has the following model:

  - many to many relationship between identities and roles.
  - many to many relationship between roles and permissions.
  - roles can have parent roles.
*/
package gorbac

import (
	"errors"
	"sync"
)

var (
	// ErrRoleNotExist occurred if a role cann't be found
	ErrRoleNotExist = errors.New("Role does not exist")
	// ErrRoleExist occurred if a role shouldn't be found
	ErrRoleExist = errors.New("Role has already existed")
	empty        = struct{}{}
)

// AssertionFunc supplies more fine-grained permission controls.
type AssertionFunc[T comparable] func(*RBAC[T], T, Permission[T]) bool

// RBAC object, in most cases it should be used as a singleton.
type RBAC[T comparable] struct {
	mutex   sync.RWMutex
	roles   Roles[T]
	parents map[T]map[T]struct{}
}

// New returns a RBAC structure.
// The default role structure will be used.
func New[T comparable]() *RBAC[T] { _ = "STUB: not implemented"; return nil }

// SetParents bind `parents` to the role `id`.
// If the role or any of parents is not existing,
// an error will be returned.
func (rbac *RBAC[T]) SetParents(id T, parents []T) error { _ = "STUB: not implemented"; return nil }

// GetParents return `parents` of the role `id`.
// If the role is not existing, an error will be returned.
// Or the role doesn't have any parents,
// a nil slice will be returned.
func (rbac *RBAC[T]) GetParents(id T) ([]T, error) { _ = "STUB: not implemented"; return nil, nil }

// SetParent bind the `parent` to the role `id`.
// If the role or the parent is not existing,
// an error will be returned.
func (rbac *RBAC[T]) SetParent(id T, parent T) error { _ = "STUB: not implemented"; return nil }

// RemoveParent unbind the `parent` with the role `id`.
// If the role or the parent is not existing,
// an error will be returned.
func (rbac *RBAC[T]) RemoveParent(id T, parent T) error { _ = "STUB: not implemented"; return nil }

// Add a role `r`.
func (rbac *RBAC[T]) Add(r Role[T]) (err error) { _ = "STUB: not implemented"; return nil }

// Remove the role by `id`.
func (rbac *RBAC[T]) Remove(id T) (err error) { _ = "STUB: not implemented"; return nil }

// Get the role by `id` and a slice of its parents id.
func (rbac *RBAC[T]) Get(id T) (r Role[T], parents []T, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// IsGranted tests if the role `id` has Permission `p` with the condition `assert`.
func (rbac *RBAC[T]) IsGranted(id T, p Permission[T],
	assert AssertionFunc[T]) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (rbac *RBAC[T]) isGranted(id T, p Permission[T],
	assert AssertionFunc[T]) bool {
	_ = "STUB: not implemented"
	return false
}

func (rbac *RBAC[T]) recursionCheck(id T, p Permission[T]) bool {
	_ = "STUB: not implemented"
	return false
}
