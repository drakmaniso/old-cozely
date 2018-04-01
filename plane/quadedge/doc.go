/*
Package quadedge provides a way to describe and manipulate the topology of any
subdivision of the (projective) plane.

It implements the quad-edge data structure described in:

  "Primitives for the Manipulation of General Subdivisions and the Computation
   of Voronoi Diagrams"

   P. Guibas, J. Stolfi, ACM TOG, April 1985

The main advantage of this data structure is that it represents at the same time
the primal and dual of the graph, as well as their mirrors. It even allows to
switch from one to another in constant time.

For example, a list of quad-edges that describe the topology a Delaunay
triangulation also describe the corresponding Voronoi diagram.

The structure does not hold (nor use) any geometrical information, but contains
32bit IDs that can be used to point to vertices and faces.
*/
package quadedge
