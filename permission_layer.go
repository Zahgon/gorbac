package gorbac

// NewLayerPermission returns an instance of layered permission with `id`
func NewLayerPermission(id, sep string) LayerPermission {
	_ = "STUB: not implemented"
	return *new(LayerPermission)
}

// LayerPermission uses string as a layered ID.
// Each layer splits by "/".
type LayerPermission struct {
	SID string `json:"id"`
	Sep string `json:"sep"`
}

// ID returns id
func (p LayerPermission) ID() string {
	_ = "STUB: not implemented"

	// Match another permission
	return ""
}

func (p LayerPermission) Match(parent Permission[string]) bool {
	_ = "STUB: not implemented"
	return false
}

// layer counts of q should be less than that of p
