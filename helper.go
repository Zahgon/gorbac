package gorbac

import "fmt"

// WalkHandler is a function defined by user to handle role
type WalkHandler[T comparable] func(Role[T], []T) error

// Walk passes each Role to WalkHandler
func Walk[T comparable](rbac *RBAC[T], h WalkHandler[T]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// InherCircle returns an error when detecting any circle inheritance.
func InherCircle[T comparable](rbac *RBAC[T]) (err error) { _ = "STUB: not implemented"; return nil }

var (
	ErrFoundCircle = fmt.Errorf("Found circle")
)

// https://en.wikipedia.org/wiki/Depth-first_search
func dfs[T comparable](rbac *RBAC[T], id T, skipped map[T]struct{},
	stack []T) error {
	_ = "STUB: not implemented"
	return nil
}

// AnyGranted checks if any role has the permission.
func AnyGranted[T comparable](rbac *RBAC[T], roles []T,
	permission Permission[T], assert AssertionFunc[T]) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// AllGranted checks if all roles have the permission.
func AllGranted[T comparable](rbac *RBAC[T], roles []T,
	permission Permission[T], assert AssertionFunc[T]) (ok bool) {
	_ = "STUB: not implemented"
	return false
}
