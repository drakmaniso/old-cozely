# glam

**This is a work in progress!**

The framework is *very* incomplete, and the API is not yet set in stone.
**Use at your own risk!**

Also note that I'm writing this framework alongside a personal project that uses
it, and the development of the former is mainly driven by the needs of the
latter. So the order in which features are added is a bit chaotic.

[![GoDoc](https://godoc.org/github.com/drakmaniso/glam?status.svg)](https://godoc.org/github.com/drakmaniso/glam)
[![Go Report Card](https://goreportcard.com/badge/github.com/drakmaniso/glam)](https://goreportcard.com/report/github.com/drakmaniso/glam)

This is a minimalist framework for making games in Go. It provides simple
abstractions to access the hardware, and tries to find a balance between
simplicity and versatility.

By order of priority, the API aims for:

- simplicity of implementation (i.e. keep it maintainable),
- ease of use (i.e. minimize boiler-plate code),
- and, when not in contradiction with the first two points, efficiency.

## Implemented Features

As of 6/03/2017, glam provides:

- A game loop with basic support for mouse, keyboard and window events.
- Simple abstractions over a modern subset of OpenGL (usable but incomplete).
- Vectors, matrices and efficient float32 math functions (some function still
  missing or unoptimized).
- A "text mode" overlay, for development and debugging purposes.

## Dependencies

The only dependancies are SDL 2 and OpenGL (with hardware capable of OpenGL
4.5).

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
