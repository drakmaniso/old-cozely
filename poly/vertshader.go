// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

//
//

const vertshader = `
#version 450 core

layout(std140, binding = 0) uniform frameBlock {
	mat4 ProjectionView;
	mat4 Model;
	vec3 CameraPosition;
  float CameraExposure;
	vec3 SunIlluminance;
  float unused1;
} frame;

struct face {
	uint MatHiVert0Vert1;
	uint MatLoVert2Vert3;
};
layout(std430, binding = 0) buffer FaceBuffer {
	face []Faces;
} faceBuffer;

struct vec3tight {
	// Needed because vec3 arrays are not tighlty packed on some hardware
	float x, y, z;
};
layout(std430, binding = 1) buffer VertexBuffer {
	vec3tight []Vertices;
} vertexBuffer;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) flat out vec3 Normal; // in world space
	layout(location = 1) flat out uint Material;
	layout(location = 2) out vec3 SurfaceToCamera;
} vertex;

void main(void) {
	// Calculate index in face buffer
	uint faceID = gl_VertexID / 6;
	// Determine which face vertex this is
	const uint [6]triangulate = {0, 1, 2, 0, 2, 3};
	uint currVert = triangulate[gl_VertexID - (6 * faceID)];

	// Read the face buffer
	face f = faceBuffer.Faces[faceID];

	// Compute indices for the vertex buffer
	uint vi[4];
	vi[1] = f.MatHiVert0Vert1 & 0x3FFF;
	f.MatHiVert0Vert1 >>= 14;
	vi[0] = f.MatHiVert0Vert1 & 0x3FFF;
	f.MatHiVert0Vert1 >>= 14;
	vi[3] = f.MatLoVert2Vert3 & 0x3FFF;
	f.MatLoVert2Vert3 >>= 14;
	vi[2] = f.MatLoVert2Vert3 & 0x3FFF;
	f.MatLoVert2Vert3 >>= 14;
	vertex.Material = f.MatHiVert0Vert1 << 4 | f.MatLoVert2Vert3;

	// Read the vertex buffer
	vec3tight v[4];
	v[0] = vertexBuffer.Vertices[vi[0]];
	v[1] = vertexBuffer.Vertices[vi[1]];
	v[2] = vertexBuffer.Vertices[vi[2]];
	v[3] = vertexBuffer.Vertices[vi[3]];
	// Convert to vec3
	vec3 p[4];
	p[0] = vec3(v[0].x, v[0].y, v[0].z);
	p[1] = vec3(v[1].x, v[1].y, v[1].z);
	p[2] = vec3(v[2].x, v[2].y, v[2].z);
	p[3] = vec3(v[3].x, v[3].y, v[3].z);

	// Compute face normal
	vertex.Normal = cross(p[2] - p[0], p[3] - p[1]);
	// Transform normal to world space
	mat3 nm = mat3(frame.Model);
	nm = transpose(inverse(nm));
	vertex.Normal = (normalize(nm * vec3(vertex.Normal))).xyz;

	// Compute screen coordinates
	vec4 wp = frame.Model * vec4(p[currVert], 1);
	gl_Position = frame.ProjectionView * wp;

	//
	vertex.SurfaceToCamera = frame.CameraPosition - wp.xyz/wp.w;
}
`

//------------------------------------------------------------------------------
