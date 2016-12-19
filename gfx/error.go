// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

// Err returns the first glam error since the previous call to Err().
func Err() error {
	err := stickyErr
	stickyErr = nil
	return err
}

//------------------------------------------------------------------------------

func setErr(err error) {
	// TODO: use two different functions and a *func variable
	// if internal.Debug {
	// 	//TODO: log?
	// 	fmt.Print("gxf error")
	// 	if stickyErr != nil {
	// 		fmt.Print(" (unchecked)")
	// 	}
	// 	fmt.Println(": ", err)
	// }
	if stickyErr == nil {
		stickyErr = err
	}
}

var stickyErr error

//------------------------------------------------------------------------------
