// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package window

import (
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// Events holds the callbacks for each window events.
//
// These callbacks can be modified at anytime, but should always contain valid
// functions (i.e., non nil). The change will take effect at the next frame.
var Events = struct {
	Resize  func()
	Hide    func()
	Show    func()
	Focus   func()
	Unfocus func()
	Quit    func()
}{
	Resize:  func() {},
	Hide:    func() {},
	Show:    func() {},
	Focus:   func() {},
	Unfocus: func() {},
	Quit:    func() { internal.QuitRequested = true },
}
