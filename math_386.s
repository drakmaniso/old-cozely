// func Sqrt(x float32) float32	
TEXT ·Sqrt(SB),7,$0
	FMOVF   x+0(FP),F0
	FSQRT
	FMOVFP  F0,ret+4(FP)
	RET

//------------------------------------------------------------------------------
// Copyright (c) 2013 - Laurent Moussault <moussault.laurent@gmail.com>
