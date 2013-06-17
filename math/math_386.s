// Based on code from the Go standard library.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the ORIGINAL_LICENSE file.

//------------------------------------------------------------------------------

// SLOWER than the Go function
// func asmRound(s float32) float32
TEXT ·asmRound(SB),7,$0
	CVTSS2SL  x+0(FP), BX
	MOVL       BX, ret+4(FP)
	RET

//------------------------------------------------------------------------------
