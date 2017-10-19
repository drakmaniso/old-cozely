# Carol

**This is a work in progress:** The framework is *very* incomplete, and the API is not yet set in stone. **Use at your own risk!**

[![GoDoc](https://godoc.org/github.com/drakmaniso/carol?status.svg)](https://godoc.org/github.com/drakmaniso/carol)
[![Go Report Card](https://goreportcard.com/badge/github.com/drakmaniso/carol)](https://goreportcard.com/report/github.com/drakmaniso/carol)

Carol is (the beginning of) a *fantasy console* programmed in Go.

## Implemented Features

(none yet)

## Dependencies

The only dependancies are SDL 2 and OpenGL 4.5.

## License

The code is under a simplified BSD license (see LICENSE file). When a sub-package
is derived from another source, the directory contain the appropriate LICENSE file.

## Credits

Some implementations of the single-precision math functions are
derived from the [Go source code](https://github.com/golang/go) (BSD-style license).

The Perlin and Simplex noise functions are adapted from
["Simplex Noise Demystified"](http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf)
by Stefan Gustavson (code in the public domain).

The pixel font used by the MTX package was originally based on ["Pixel Operator
Mono"](https://notabug.org/HarvettFox96/ttf-pixeloperator) by Jayvee Enayas, but
has been so modified that it's now a completely different font. It is still
licensed under the SIL OFL.
