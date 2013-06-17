// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

//------------------------------------------------------------------------------

// SLOWER than the Go function
// func asmRound(s float32) float32
TEXT ·asmRound(SB),7,$0
	CVTSS2SL  x+0(FP), BX
	MOVL       BX, ret+4(FP)
	RET

//------------------------------------------------------------------------------
