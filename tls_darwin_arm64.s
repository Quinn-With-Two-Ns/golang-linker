//go:build darwin && arm64

#include "textflag.h"

TEXT ·tlsNext(SB), NOSPLIT, $0-4
	BL tls_next(SB)
	MOVW R0, ret+0(FP)
	RET
