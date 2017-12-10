// Based on code from the Go standard library.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the ORIGINAL_LICENSE file.

package math32

//------------------------------------------------------------------------------

// Abs returns the absolute value of x.
func Abs(x float32) float32 {
	// TODO: rewrite when golang.org/issue/13095 is fixed
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

//------------------------------------------------------------------------------
