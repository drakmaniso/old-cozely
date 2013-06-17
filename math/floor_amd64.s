// Based on code from the Go standard library.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the ORIGINAL_LICENSE file.

//------------------------------------------------------------------------------

// func Floor(s float32) float32
TEXT ·Floor(SB),7,$0
	MOVL       x+0(FP), AX
	MOVL       AX, X0 // X0 = x
	CVTTSS2SL  X0, AX // AX = int(x)
	CVTSL2SS   AX, X1 // X1 = float(int(x))
	CMPSS      X1, X0, 1 // compare LT; X0 = 0xffffffffffffffff or 0
	MOVSS      $(-1.0), X2
	ANDPS      X2, X0 // if x < float(int(x)) {X0 = -1} else {X0 = 0}
	ADDSS      X1, X0
	MOVSS      X0, ret+8(FP)
	RET
	
//------------------------------------------------------------------------------
