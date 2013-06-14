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
