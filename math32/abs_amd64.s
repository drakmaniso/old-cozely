// Based on code from the Go standard library.
// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the ORIGINALLICENSE file.

//------------------------------------------------------------------------------

#include "textflag.h"

// func absAsm(x float32) float32
TEXT ·absAsm(SB),NOSPLIT,$0
	MOVL   x+0(FP), AX
	SHLL   $1, AX
	SHRL   $1, AX
	MOVL   AX, ret+8(FP)
	RET
//	MOVL   $(1<<31), BX
//	MOVL   BX, X0 // movss $(-0.0), x0
//	MOVSS  x+0(FP), X1
//	ANDNPS X1, X0
//	MOVSS  X0, ret+8(FP)
//	RET

//------------------------------------------------------------------------------
