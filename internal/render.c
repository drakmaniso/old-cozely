#include "glad.h"
#include "sdl.h"
#include "render.h"

int InitOpenGL() {
	if(!gladLoadGLLoader(SDL_GL_GetProcAddress)) {
		return -1;
	}
	glClearColor (0.45, 0.31, 0.59, 1.0);

    glEnable (GL_DEPTH_TEST);
    glClearDepth (1.0);
    glDepthFunc (GL_LEQUAL);

    glEnable (GL_CULL_FACE);
    glCullFace (GL_BACK);

    glBlendFunc (GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA);

    glEnable(GL_FRAMEBUFFER_SRGB);
	
	return 0;
}

void Render(SDL_Window* w) {
	glClear (GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
	SDL_GL_SwapWindow(w);
}

GLenum GetOpenGLError() {
	return glGetError();
}