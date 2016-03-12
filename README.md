Glam: a minimalist framework for making games in Go
===================================================


**WARNING: This is a work in progress, very incomplete and not yet functional.**

[![GoDoc](https://godoc.org/github.com/drakmaniso/glam?status.svg)](https://godoc.org/github.com/drakmaniso/glam)

Goals
-----

This project is currently for my personal use. It's not meant as a general 
purpose engine, but may become useful to others at some point.

The main inspiration behind the design is the Lua framework Löve. The goal is to
provide thin abstractions above OpenGL and SDL.

By order of priority, the API aims for:

- simplicity of implementation (i.e. keep it small and manageable),
- ease of use (e.g. avoid unnecessary abstractions, inconvenient APIs),
- and, only when not in contradiction with the first two points, efficiency.


Implemented Features
--------------------

- Package engine: game loop and platform-dependent features (just started)
- Package geom: vectors and matrices (incomplete).
- Package key: provides support for the keyboard (usable but incomplete).
- Package mouse: provides suport for the mouse (incomplete).
- Package math: efficient single-precision math.
- Package noise: perlin noise.


License
-------

The code is under a simplified BSD license (see LICENSE file). When a sub-package
is derived from anothe source, the directory contain the appropriate LICENSE file. 


Credits
-------

Some implementations of the single-precision math functions are
derived from the Go source code.

The Perlin and Simplex noise functions are adapted from
["Simplex Noise Demystified"](http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf)
by Stefan Gustavson (code in the public domain).


Author
------

Laurent Moussault <moussault.laurent@gmail.com>
