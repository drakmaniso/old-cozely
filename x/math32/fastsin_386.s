// This code is adapted from:
// http://devmaster.net/forums/topic/4648-fast-and-accurate-sinecosine/
// Copyright Nicolas Capens

////////////////////////////////////////////////////////////////////////////////

#define ABS 0x7FFFFFFF

#define B 1.27323954473516268615107010698
#define C -0.4052847345693510857755178528389

#define P 0.224008178776
#define Q 0.775991821224

// func FastSin(s float32) float32
TEXT ·FastSin(SB),7,$0
	MOVL		x+0(FP), AX
	MOVL		AX, X0		// X0 = x
	ANDL		$ABS, AX
	MOVL		AX, X1		// X1 = |x|
	MULSS		X0, X1		// X1 = x * |x|
	MULSS		$B, X0		// X0 = B * x
	MULSS		$C, X1		// X1 = C * x * |x|
	ADDSS		X1, X0		// X0 = y = C * x * |x| + B * x

	MOVL		X0, AX
	ANDL		$ABS, AX
	MOVL		AX, X1		// X1 = |y|
	MULSS		X0, X1		// X1 = y * |y|
	MULSS		$Q, X0		// X0 = Q * y
	MULSS		$P, X1		// X1 = P * y * |y|
	ADDSS		X1, X0		// X0 = P * y * |y| + Q * y

	MOVSS		X0, ret+4(FP)
	RET

////////////////////////////////////////////////////////////////////////////////
