// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

#include "engine.h"

SDL_Event events[PEEP_SIZE];

int peepEvents() {
	SDL_PumpEvents();
	int n = SDL_PeepEvents(events, PEEP_SIZE, SDL_GETEVENT, SDL_FIRSTEVENT, SDL_LASTEVENT);
	return n;
}
