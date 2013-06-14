// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Based on code from the Go standard library.
// Original license:
//
// Copyright (c) 2012 The Go Authors. All rights reserved.
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
// 
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
	
//------------------------------------------------------------------------------

// SLOWER than the Go function
// func FastFloor(s float32) int32
//TEXT ·FastFloor(SB),7,$0
//	CVTTSS2SL  x+0(FP), BX
//	MOVL       x+0(FP), AX
//	SHRL       $31, AX
//	SUBL       AX, BX
//	MOVL       BX,ret+8(FP)
//	RET

//------------------------------------------------------------------------------
