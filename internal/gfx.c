// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

#include "glad.h"
#include "gfx.h"

GLuint CompileShader(const GLchar* b, GLenum t) {
	GLuint s = glCreateShader(t);
	if (s == 0) {
		return 0;
	}
	glShaderSource(s, 1, &b, NULL);
	glCompileShader(s);

	return s;
}

char* CheckCompileShaderError(GLuint s) {
    GLint status = GL_TRUE;
    glGetShaderiv (s, GL_COMPILE_STATUS, &status);
    if (status != GL_TRUE) {
        GLint length = 0;
        glGetShaderiv (s, GL_INFO_LOG_LENGTH, &length);
        char *message = calloc(length + 1, sizeof(char));
        glGetShaderInfoLog (s, length, &length, message);
        return message;
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