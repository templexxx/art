package artree

// cas16b executes the compare-and-swap operation for a 16bytes value.
// dst must be 16bytes aligned.
//go:noescape
func cas16b(dst, old, new *byte) (swapped bool)
