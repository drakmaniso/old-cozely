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
	uint []faces;
} faceSSBO;

layout(std430, binding = 1) buffer VertexSSBO {
	float []vertices;
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

	// faceID = 2;
	uvec2 face;
	face.x = faceSSBO.faces[faceID*2+0];
	face.y = faceSSBO.faces[faceID*2+1];
	vertex.Material = face.x & 0xFFFF;
	uint vert1 = face.x >> 16;
	uint vert2 = face.y & 0xFFFF;
	uint vert3 = face.y >> 16;

	vec3 p[3];
	p[0].x = vertexSSBO.vertices[vert1*3 + 0];
	p[0].y = vertexSSBO.vertices[vert1*3 + 1];
	p[0].z = vertexSSBO.vertices[vert1*3 + 2];
	p[1].x = vertexSSBO.vertices[vert2*3 + 0];
	p[1].y = vertexSSBO.vertices[vert2*3 + 1];
	p[1].z = vertexSSBO.vertices[vert2*3 + 2];
	p[2].x = vertexSSBO.vertices[vert3*3 + 0];
	p[2].y = vertexSSBO.vertices[vert3*3 + 1];
	p[2].z = vertexSSBO.vertices[vert3*3 + 2];

	// vec3 v1 = p[1] - p[0];
	// vec3 v2 = p[2] - p[0];
	// vertex.Normal = normalize(cross(v1, v2));

	gl_Position = frame.Projection * frame.View * frame.Model * vec4(p[currVert], 1);

	// gl_Position.x = 0.5*vertexSSBO.vertices[gl_VertexID * 3];
	// gl_Position.y = 0.5*vertexSSBO.vertices[gl_VertexID * 3 + 1];
	// gl_Position.z = 0.5*vertexSSBO.vertices[gl_VertexID * 3 + 2];
	// gl_Position.w = 1.0;
	// gl_Position = frame.Projection * frame.View * frame.Model * gl_Position;

	// const vec4 triangle[3] = vec4[3](
	// 	vec4(0, 0.65, 0.5, 1),
	// 	vec4(-0.65, -0.475, 0.5, 1),
	// 	vec4(0.65, -0.475, 0.5, 1)
	// );
	// gl_Position = triangle[gl_VertexID];
}
`

//------------------------------------------------------------------------------
