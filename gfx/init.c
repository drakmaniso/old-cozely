// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

#include "glad.h"
#include "_cgo_export.h"

void errorCallback(
	GLenum source,
	GLenum type,
	GLuint id,
   	GLenum severity,
	GLsizei length,
	const GLchar* message,
	const void* userParam
) {
	logGLError(source, type, id, severity, length, (char*)message);
}

int InitOpenGL(int debug) {
	if(!gladLoadGL()) {
		return -1;
	}
	
	if(debug) {
		glEnable(GL_DEBUG_OUTPUT);
		glDebugMessageCallback(errorCallback, NULL);
	}
	
	return 0;
}
