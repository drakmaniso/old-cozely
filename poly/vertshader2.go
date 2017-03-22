// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

//
//

const vertshader2 = `
#version 450 core

//------------------------------------------------------------------------------

layout(std140, binding = 0) uniform glam_Camera {
	mat4  glam_ProjectionView;
	vec3  glam_CameraPosition;
  float glam_CameraExposure;
};

layout(std140, binding = 1) uniform Misc {
	mat4  glam_Model;
	vec3  glam_SunIlluminance;
  float glam_unused1;
};

//------------------------------------------------------------------------------

struct glam_Face {
	uint MatHiVert0Vert1;
	uint MatLoVert2Vert3;
};
layout(std430, binding = 0) buffer glam_FaceBuffer {
	glam_Face []glam_Faces;
};

struct glam_Vec3Tight {
	// Needed because vec3 arrays are not tighlty packed on some hardware
	float x, y, z;
};
layout(std430, binding = 1) buffer glam_VertexBuffer {
	glam_Vec3Tight []glam_Vertices;
};

//------------------------------------------------------------------------------

out gl_PerVertex {
	vec4 gl_Position;
};
out glam_PerVertex {
	layout(location = 0) flat vec3 glam_Normal; // in world space
	layout(location = 1)      vec3 glam_SurfaceToCamera;
	layout(location = 2) flat uint glam_Material;
};

//------------------------------------------------------------------------------

void glam_PrepareVertex() {
	// Calculate index in face buffer
	uint faceID = gl_VertexID / 6;
	// Determine which face vertex this is
	const uint [6]triangulate = {0, 1, 2, 0, 2, 3};
	uint currVert = triangulate[gl_VertexID - (6 * faceID)];

	// Read the face buffer
	glam_Face f = glam_Faces[faceID];

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
	glam_Material = f.MatHiVert0Vert1 << 4 | f.MatLoVert2Vert3;

	// Read the vertex buffer
	glam_Vec3Tight v[4];
	v[0] = glam_Vertices[vi[0]];
	v[1] = glam_Vertices[vi[1]];
	v[2] = glam_Vertices[vi[2]];
	v[3] = glam_Vertices[vi[3]];
	// Convert to vec3
	vec3 p[4];
	p[0] = vec3(v[0].x, v[0].y, v[0].z);
	p[1] = vec3(v[1].x, v[1].y, v[1].z);
	p[2] = vec3(v[2].x, v[2].y, v[2].z);
	p[3] = vec3(v[3].x, v[3].y, v[3].z);

	// Compute face normal
	const uint [4][4]tri = {
		{0, 1, 2, 3},
		{1, 2, 3, 0},
		{2, 3, 0, 1},
		{3, 0, 1, 2}
	};
	vec3 n1 = cross(p[tri[currVert][1]] - p[tri[currVert][0]], p[tri[currVert][2]] - p[tri[currVert][0]]);
	vec3 n2 = cross(p[tri[currVert][2]] - p[tri[currVert][0]], p[tri[currVert][3]] - p[tri[currVert][0]]);
	vec3 n3 = cross(p[tri[currVert][1]] - p[tri[currVert][0]], p[tri[currVert][3]] - p[tri[currVert][0]]);
	glam_Normal = normalize(n1+n2+n3);

	// Transform normal to world space
	mat3 nm = mat3(glam_Model);
	nm = transpose(inverse(nm));
	glam_Normal = (normalize(nm * vec3(glam_Normal))).xyz;

	// Compute screen coordinates
	vec4 wp = glam_Model * vec4(p[currVert], 1);
	gl_Position = glam_ProjectionView * wp;

	//
	glam_SurfaceToCamera = glam_CameraPosition - wp.xyz/wp.w;
}
`

//------------------------------------------------------------------------------
