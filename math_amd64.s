// func Sqrt(x float32) float32
TEXT ·Sqrt(SB),7,$0
	SQRTSS x+0(FP), X0
	MOVSS X0, ret+8(FP)
	RET

//------------------------------------------------------------------------------
// Copyright (c) 2013 - Laurent Moussault <moussault.laurent@gmail.com>
