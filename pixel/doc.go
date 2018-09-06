// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

/*
Package pixel provides a drwaing canvas specialized for pixel art.

It implements a GPU pipeline that works entirely in indexed colors at a chosen,
fixed, resolution (usually lower than the screen resolution).

It does not provide any anti-aliasing, nor alpha-transparency, since the goal is
to offer an easy way to work within the stricter definition of pixel art
(limited color palette, no mixed-resolution or "mixels", and so on).
*/
package pixel
