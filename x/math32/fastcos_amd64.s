// This code is adapted from:
// http://devmaster.net/forums/topic/4648-fast-and-accurate-sinecosine/
// Copyright Nicolas Capens

////////////////////////////////////////////////////////////////////////////////

#define PISLASHTWO 1.57079632679489661923132169164
#define PI 3.14159265358979323846264338328
#define TWOPI 6.283185307179586476925286766559

#define ABS 0x7FFFFFFF

#define B 1.27323954473516268615107010698
#define C -0.4052847345693510857755178528389

#define P 0.224008178776
#define Q 0.775991821224

// func FastCos(s float32) float32
TEXT ·FastCos(SB),7,$0
	MOVL		x+0(FP), X0	// X0 = x
	ADDSS		$PISLASHTWO, X0	// X0 = x + Pi/2
	MOVSS		$PI, X1
	UCOMISS		X1, X0
	JLS			inrange		// if x + Pi/2 < Pi jump to inrange
	MOVSS		$TWOPI, X1
	SUBSS		X1, X0		// X0 = x + Pi/2 - 2*Pi

/*	MOVSS		X0, X1		// X1 = x + Pi/2
	MOVSS		$PI, X2
	CMPSS		X2, X1, 5	// X1 = x + Pi/2 >= Pi
	MOVSS		$TWOPI, X2
	ANDPS		X2, X1		// X1 = (x + Pi/2 >= Pi) ? 2*Pi : 0
	SUBSS		X1, X0		// X0 = x = (x + Pi/2) modulo (2*Pi)*/

inrange:

	MOVL		X0, AX
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

	MOVSS		X0, ret+8(FP)
	RET

////////////////////////////////////////////////////////////////////////////////
