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

layout(std430, binding = 0) buffer FaceSSBO {
	uvec2 []faces;
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

void main(void) {
	uint faceID = gl_VertexID / 3;
	uint currVert = gl_VertexID - (3 * faceID);

	uvec2 face = faceSSBO.faces[faceID];
	vertex.Material = face.x & 0xFFFF;
	uint vertID[3];
	vertID[0] = face.x >> 16;
	vertID[1] = face.y & 0xFFFF;
	vertID[2] = face.y >> 16;

	vec3tight verts[3];
	verts[0] = vertexSSBO.vertices[vertID[0]];
	verts[1] = vertexSSBO.vertices[vertID[1]];
	verts[2] = vertexSSBO.vertices[vertID[2]];
	vec3 p[3];
	p[0] = vec3(frame.Model * vec4(verts[0].x, verts[0].y, verts[0].z, 1.0));
	p[1] = vec3(frame.Model * vec4(verts[1].x, verts[1].y, verts[1].z, 1.0));
	p[2] = vec3(frame.Model * vec4(verts[2].x, verts[2].y, verts[2].z, 1.0));

	vertex.Normal = normalize(cross(p[1] - p[0], p[2] - p[0]));

	gl_Position = frame.Projection * frame.View * vec4(p[currVert], 1);
}
`

//------------------------------------------------------------------------------
