# Glam

**This is a work in progress! The framework is very incomplete, and even the
existing API is not yet set in stone. Use at your own risk!**

[![GoDoc](https://godoc.org/github.com/drakmaniso/glam?status.svg)](https://godoc.org/github.com/drakmaniso/glam)
[![Go Report Card](https://goreportcard.com/badge/github.com/drakmaniso/glam)](https://goreportcard.com/report/github.com/drakmaniso/glam)

Glam is a minimalist framework for making games in Go. It provides simple
abstractions to access to the hardware, and tries to find a balance between
simplicity and versatility.

By order of priority, the API aims for:

- simplicity of implementation (keep it small and manageable, avoid unnecessary abstractions),
- ease of use (minimize boiler-plate code),
- and, only when not in contradiction with the first two points, efficiency.

## Implemented Features

As of 14/01/2017, Glam provides:

- a game loop with basic support for mouse, keyboard and window events,
- vectors, matrices and efficient float32 math functions (incomplete),
- simple abstractions over a modern subset of OpenGL (*very* incomplete).

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
