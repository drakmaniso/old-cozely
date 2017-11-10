package gpu

/*
#include "glad.h"

static inline void SetupQuadPipeline(GLuint *program, GLuint*vao) {
	*program = glCreateProgram();
	glCreateVertexArrays(1, vao);
}
*/
import "C"

//------------------------------------------------------------------------------

// A Quad is...
//
//TODO
type Quad struct {
	// First word
	X1, Y1 int16

	// Second word
	X2, Y2 int16

	// Third word
	Xtex, Ytex int16

	// Fourth word
	TexIndex   uint16
	Palette    uint8
	ZoomOrient byte
}

//------------------------------------------------------------------------------

var QuadPipeline struct {
	program C.GLuint
	vao     C.GLuint
}

func SetupQuadPipeline() {
	C.SetupQuadPipeline(&QuadPipeline.program, &QuadPipeline.vao)
}
