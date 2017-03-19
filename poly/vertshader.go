// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

//
//

const vertshader = `
#version 450 core

layout(std140, binding = 0) uniform Frame {
	mat4 Projection;
	mat4 View;
	mat4 Model;
} frame;

struct facetight {
	uint mathivert0vert1;
	uint matlovert2vert3;
};
layout(std430, binding = 0) buffer FaceSSBO {
	facetight []faces;
} faceSSBO;

struct vec3tight {
	float x, y, z;
};
layout(std430, binding = 1) buffer VertexSSBO {
	vec3tight []vertices;
} vertexSSBO;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) flat out vec3 Normal;
	layout(location = 1) flat out uint Material;
} vertex;

const uint [6]triangulate = {0, 1, 2, 0, 2, 3};

void main(void) {
	uint faceID = gl_VertexID / 6;
	uint currVert = triangulate[gl_VertexID - (6 * faceID)];

	facetight face = faceSSBO.faces[faceID];
	uint vertID[4];
	vertID[1] = face.mathivert0vert1 & 0x3FFF;
	face.mathivert0vert1 >>= 14;
	vertID[0] = face.mathivert0vert1 & 0x3FFF;
	face.mathivert0vert1 >>= 14;
	vertID[3] = face.matlovert2vert3 & 0x3FFF;
	face.matlovert2vert3 >>= 14;
	vertID[2] = face.matlovert2vert3 & 0x3FFF;
	face.matlovert2vert3 >>= 14;
	vertex.Material = face.mathivert0vert1 << 4 | face.matlovert2vert3;

	vec3tight verts[4];
	verts[0] = vertexSSBO.vertices[vertID[0]];
	verts[1] = vertexSSBO.vertices[vertID[1]];
	verts[2] = vertexSSBO.vertices[vertID[2]];
	verts[3] = vertexSSBO.vertices[vertID[3]];
	vec3 p[4];
	p[0] = vec3(frame.Model * vec4(verts[0].x, verts[0].y, verts[0].z, 1.0));
	p[1] = vec3(frame.Model * vec4(verts[1].x, verts[1].y, verts[1].z, 1.0));
	p[2] = vec3(frame.Model * vec4(verts[2].x, verts[2].y, verts[2].z, 1.0));
	p[3] = vec3(frame.Model * vec4(verts[3].x, verts[3].y, verts[3].z, 1.0));

	vertex.Normal = normalize(cross(p[2] - p[0], p[3] - p[1]));

	gl_Position = frame.Projection * frame.View * vec4(p[currVert], 1);
}
`

//------------------------------------------------------------------------------
