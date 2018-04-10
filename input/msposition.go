// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type msPosition struct {
}

func (a msPosition) bind(c Context, target Action)   {}
func (a msPosition) device() Device                  { return noDevice }
func (a msPosition) action() Action                  { return nil }
func (a msPosition) asBool() (just bool, value bool) { return false, false }
