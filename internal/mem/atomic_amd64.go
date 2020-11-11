package mem

// AtomicLoad16B atomically loads 16bytes from *addr.
// addr must be 16bytes aligned.
//
//go:noescape
func AtomicLoad16B(addr *byte) [16]byte

// AtomicLoad16B atomically stores 16bytes to *addr.
// addr must be 16bytes aligned.
//
//go:noescape
func AtomicStore16B(addr *byte, val [16]byte)

// AtomicCAS16B executes the compare-and-swap operation for a 16bytes value.
// addr must be 16bytes aligned.
//
// Using *byte saving memory copy.
//go:noescape
func AtomicCAS16B(addr, old, new *byte) (swapped bool)
