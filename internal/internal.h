// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

#if defined(__WIN32)
	#include <SDL2/SDL.h>
#else
	#include <SDL.h>
#endif


#define PEEP_SIZE 128

SDL_Event events[PEEP_SIZE];

Uint8 *keystate;
SDL_Keymod keymod;

void initC();
int peepEvents();