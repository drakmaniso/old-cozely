//------------------------------------------------------------------------------

// func Sqrt(x float32) float32	
TEXT ·Sqrt(SB),7,$0
	FMOVF   x+0(FP),F0
	FSQRT
	FMOVFP  F0,ret+4(FP)
	RET

//------------------------------------------------------------------------------

// func Floor(s float32) float32
TEXT ·Floor(SB),7,$0
	FMOVF   x+0(FP), F0  // F0=x
	FSTCW   -2(SP)       // save old Control Word
	MOVL    -2(SP), AX
	ANDW    $0xf3ff, AX
	ORW     $0x0400, AX  // Rounding Control set to -Inf
	MOVW    AX, -4(SP)   // store new Control Word
	FLDCW   -4(SP)       // load new Control Word
	FRNDINT              // F0=Floor(x)
	FLDCW   -2(SP)       // load old Control Word
	FMOVFP  F0, ret+4(FP)
	RET


//------------------------------------------------------------------------------
// Copyright (c) 2013 - Laurent Moussault <moussault.laurent@gmail.com>
