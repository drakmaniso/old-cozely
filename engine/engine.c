// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

#include "engine.h"

SDL_Event events[PEEP_SIZE];

Uint8 *keystate;
SDL_Keymod keymod;

int mouseX, mouseY;
Uint32 mouseButtons;

void initC() {
	keystate = (Uint8*)SDL_GetKeyboardState(NULL);
}

int peepEvents() {
	SDL_PumpEvents();
	int n = SDL_PeepEvents(events, PEEP_SIZE, SDL_GETEVENT, SDL_FIRSTEVENT, SDL_LASTEVENT);
	keymod = SDL_GetModState();
	mouseButtons = SDL_GetMouseState(&mouseX, &mouseY);
	return n;
}