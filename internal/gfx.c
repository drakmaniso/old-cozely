// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

#include "glad.h"

GLuint CompileShader(GLenum t, const GLchar* b) {
	GLuint s = glCreateShaderProgramv(t, 1, &b);
	if (s == 0) {
		return 0;
	}

	return s;
}

GLuint NewPipeline() {
	GLuint p;
	glCreateProgramPipelines(1, &p);
	return p;
}

GLuint CreateVAO() {
	GLuint vao;
	glCreateVertexArrays(1, &vao);
	return vao;
}

void PipelineUseShader(GLuint p, GLenum stages, GLuint shader) {
	glUseProgramStages(p, stages, shader);
}

char* ShaderLinkError(GLuint p) {
    GLint ok = GL_TRUE;
    glGetProgramiv (p, GL_LINK_STATUS, &ok);
    if (ok != GL_TRUE)
    {
        GLint l = 0;
        glGetProgramiv (p, GL_INFO_LOG_LENGTH, &l);
        char *m = calloc(l + 1, sizeof(char));
        glGetProgramInfoLog (p, l, &l, m);
        return m;
    }
	
	return NULL;
}

void VertexAttribute(
	GLuint vao, 
	GLuint index, 
	GLuint binding,
	GLint size,
	GLenum type, 
	GLboolean normalized,
	GLuint relativeOffset
) {
	glVertexArrayAttribFormat(vao, index, size, type, normalized, relativeOffset);
	glVertexArrayAttribBinding(vao, index, binding);
	glEnableVertexArrayAttrib(vao, index);
}

GLuint NewBuffer(GLsizeiptr size, void* data, GLenum flags) {
	GLuint b;
	glCreateBuffers(1, &b);
	glNamedBufferStorage(b, size, data, flags);
	return b;
}

void ClosePipeline(GLuint p, GLuint vao) {
	glDeleteVertexArrays(1, &vao);
	glDeleteProgramPipelines(1, &p);
}
