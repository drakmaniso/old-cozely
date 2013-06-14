//------------------------------------------------------------------------------

// func Sqrt(x float32) float32
TEXT ·Sqrt(SB),7,$0
	SQRTSS     x+0(FP), X0
	MOVSS      X0, ret+8(FP)
	RET

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
	
// Doesn't work (e.g. for -3.3)
//TEXT ·Floor(SB),7,$0
//	CVTTSS2SL  x+0(FP), BX
//	MOVL       x+0(FP), AX
//	SHRL       $31, AX
//	SUBL       AX, BX
//	MOVL       BX,ret+8(FP)
//	RET

//------------------------------------------------------------------------------
// Copyright (c) 2013 - Laurent Moussault <moussault.laurent@gmail.com>
