// Based on code from the Go standard library.
// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the ORIGINALLICENSE file.

//------------------------------------------------------------------------------

// func Abs(x float32) float32
TEXT ·abs_asm(SB),7,$0
	MOVL   $(1<<31), BX
	MOVL   BX, X0 // movsd $(-0.0), x0
	MOVSS  x+0(FP), X1
	ANDNPS X1, X0
	MOVSS  X0, ret+4(FP)
	RET

//------------------------------------------------------------------------------
