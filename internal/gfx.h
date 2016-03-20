// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

#include "glad.h"

GLuint CompileShader(const GLchar* b, GLenum t);
char* CompileShaderError(GLuint s);
GLuint LinkProgram(GLuint vs, GLuint fs);
char* LinkProgramError(GLuint s);
