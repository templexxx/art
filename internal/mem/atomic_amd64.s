#include "textflag.h"

// func AtomicLoad16B(addr *byte) [16]byte
TEXT ·AtomicLoad16B(SB), NOSPLIT, $0
    MOVQ addr+0(FP), R8
	XORQ AX, AX
	XORQ DX, DX
	XORQ BX, BX
	XORQ CX, CX
	LOCK
	CMPXCHG16B (R8)
	MOVQ AX, val+8(FP)
	MOVQ DX, val+16(FP)
	RET

// func AtomicStore16B(addr *byte, val [16]byte)
TEXT ·AtomicStore16B(SB),NOSPLIT,$0
	MOVQ addr+0(FP), BP
	XORQ AX, AX
	XORQ DX, DX
	MOVQ val+8(FP), BX
	MOVQ val+16(FP), CX

loop:   // loop here, try to let RDX:RAX equal to m128.
	LOCK
	CMPXCHG16B (BP)
	JNE loop
	RET

// func AtomicCAS16B(addr, old, new *byte) (swapped bool)
//
// Compare RDX:RAX with m128.
// If equal, set ZF and load RCX:RBX into m128.
// Else, clear ZF and load m128 into RDX:RAX.
TEXT ·AtomicCAS16B(SB), NOSPLIT, $0-25
    MOVQ  addr+0(FP), R8
    MOVQ  old+8(FP), R9
    MOVQ  new+16(FP), R10
    MOVQ  (R9), AX
    MOVQ  8(R9), DX
    MOVQ  (R10), BX
    MOVQ  8(R10), CX
    LOCK
    CMPXCHG16B (R8)
    SETEQ	ret+24(FP)
    RET
