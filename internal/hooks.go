// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

// Package gl ///////////////////////////////////////////////////////////////

// GLSetup hook
var GLSetup = func() error { return nil }

// GLCleanup hook
var GLCleanup = func() error { return nil }

// GLErr hook
var GLErr = func() error { return nil }

// Package input ///////////////////////////////////////////////////////////////

// InputSetup hook
var InputSetup = func() error { return nil }

// InputNewFrame hook
var InputNewFrame = func() error { return nil }

// InputCleanup hook
var InputCleanup = func() error { return nil }

// InputErr hook
var InputErr = func() error { return nil }

// Package color ///////////////////////////////////////////////////////////////

// ColorSetup hook
var ColorSetup = func() error { return nil }

// ColorUpload hook
var ColorUpload = func() error { return nil }

// ColorCleanup hook
var ColorCleanup = func() error { return nil }

// ColorErr hook
var ColorErr = func() error { return nil }

// Package pixel ///////////////////////////////////////////////////////////////

// PixelSetup hook
var PixelSetup = func() error { return nil }

// PixelResize hook
var PixelResize = func() {}

// PixelCleanup hook
var PixelCleanup = func() error { return nil }

// PixelErr hook
var PixelErr = func() error { return nil }

// Package poly ////////////////////////////////////////////////////////////////

// PolySetup hook
var PolySetup = func() error { return nil }

// PolyCleanup hook
var PolyCleanup = func() error { return nil }

// PolyErr hook
var PolyErr = func() error { return nil }

// Package vector //////////////////////////////////////////////////////////////

// VectorSetup hook
var VectorSetup = func() error { return nil }

// VectorDraw hook
var VectorDraw = func() error { return nil }

// VectorCleanup hook
var VectorCleanup = func() error { return nil }

// VectorErr hook
var VectorErr = func() error { return nil }
