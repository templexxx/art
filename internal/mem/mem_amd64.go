package mem

// CAS16B executes the compare-and-swap operation for a 16bytes value.
// dst must be 16bytes aligned.
//
// Using *byte saving memory copy.
//go:noescape
func CAS16B(dst, old, new *byte) (swapped bool)
