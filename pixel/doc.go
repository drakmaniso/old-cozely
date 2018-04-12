// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

/*
Package pixel provides types and drawing functions specialized for pixel art.

It implements a GPU pipeline that work entirely in indexed colors at a chosen,
fixed, resolution.

It does not provide any anti-aliasing, nor alpha-transparency, since the goal is
to offer an easy way to work within the stricter definition of pixel art
(limited color color, no mixed-resolution or "mixels", and so on).
*/
package pixel
