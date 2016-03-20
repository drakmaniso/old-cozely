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