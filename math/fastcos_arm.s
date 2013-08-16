// This code is adapted from: 
// http://devmaster.net/forums/topic/4648-fast-and-accurate-sinecosine/
// Copyright Nicolas Capens

//------------------------------------------------------------------------------

// func FastCos(s float32) float32
TEXT ·FastCos(SB),7,$0
	B ·fastCos(SB)

//------------------------------------------------------------------------------
