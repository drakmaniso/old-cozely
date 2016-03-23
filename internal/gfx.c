// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

#include "glad.h"

GLuint CompileShader(const GLchar* b, GLenum t) {
	GLuint s = glCreateShader(t);
	if (s == 0) {
		return 0;
	}
	glShaderSource(s, 1, &b, NULL);
	glCompileShader(s);

	return s;
}

char* CompileShaderError(GLuint s) {
    GLint ok = GL_TRUE;
    glGetShaderiv (s, GL_COMPILE_STATUS, &ok);
    if (ok != GL_TRUE) {
        GLint l = 0;
        glGetShaderiv (s, GL_INFO_LOG_LENGTH, &l);
        char *m = calloc(l + 1, sizeof(char));
        glGetShaderInfoLog (s, l, &l, m);
        return m;
    }
	
	return NULL;
}

GLuint LinkProgram(GLuint vs, GLuint fs) {
	GLuint p = glCreateProgram();
	glAttachShader(p, vs);
	glAttachShader(p, fs);
	glLinkProgram(p);	
	return p;
}

char* LinkProgramError(GLuint p) {
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

GLuint SetupVAO() {
	GLuint vao;
	glCreateVertexArrays(1, &vao);
	return vao;
}

void CreateAttributeBinding(
	GLuint vao, 
	GLuint index, 
	GLuint binding,
	GLint size,
	GLenum type, 
	GLboolean normalized,
	GLuint relativeOffset
) {
	glEnableVertexArrayAttrib(vao, index);
	glVertexArrayAttribBinding(vao, index, binding);
	glVertexArrayAttribFormat(vao, index, size, type, normalized, relativeOffset);
}

GLuint CreateBufferFrom(GLsizeiptr size, const GLvoid* data) {
	GLuint b;
	glCreateBuffers(1, &b);
	glNamedBufferStorage(b, size, data, 0);
	return b;
}

void ClosePipeline(GLuint p, GLuint vao) {
	glDeleteVertexArrays(1, &vao);
	glDeleteProgram(p);
}
