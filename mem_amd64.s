#include "textflag.h"

// func cas16b(dst, old, new *byte) (swapped bool)
//Compare RDX:RAX with m128. If equal, set ZF and load
//  CMPXCHG16B m128
//   RCX:RBX into m128. Else, clear ZF and load m128 into
//  RDX:RAX.

TEXT ·cas16b(SB), NOSPLIT, $0-25
    MOVQ  dst+0(FP), R8
    MOVQ  old+8(FP), R9
    MOVQ  new+16(FP), R10
    MOVQ  (R9), DX
    MOVQ  8(R9), AX
    MOVQ  (R10), CX
    MOVQ  8(R10), BX
    LOCK
    CMPXCHG16B (R8)
    SETEQ	ret+24(FP)
    RET
