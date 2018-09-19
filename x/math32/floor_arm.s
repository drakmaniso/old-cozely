// Based on code from the Go standard library.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the ORIGINAL_LICENSE file.

//------------------------------------------------------------------------------

// func Floor(s float32) float32
TEXT ·Floor(SB),7,$0
	B ·floor(SB)
	
//------------------------------------------------------------------------------
