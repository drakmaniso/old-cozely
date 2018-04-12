/*
Package quadedge provides a way to describe and manipulate the topology of
planar maps (i.e., the topology of any closed polygonal mesh).

It implements the quad-edge data structure described in:

  "Primitives for the Manipulation of General Subdivisions and the Computation
   of Voronoi Diagrams"

   P. Guibas, J. Stolfi, ACM TOG, April 1985

This structure represents at the same time the primal and dual of the map, as
well as their mirror images. It allows to switch from one to another in constant
time.

For example, a list of quad-edges that describe the topology a Delaunay
triangulation also describe the corresponding Voronoi diagram.

The structure does not hold (nor use) any geometrical information, but contains
32bit IDs that can be used to identify vertices and faces.

Note that in this package, each (undirected) edge of the map is represented by
four directed Edge objects.
*/
package quadedge
