package key

// #include "../engine/engine.h"
import "C"

// A Modifier represents one or more key modifiers (e.g. Shift, Ctrl,
// Caps Lock...)
type Modifier uint32

// Modifiers constants
const (
	ModifierNone     Modifier = C.KMOD_NONE
	ModifierLShift   Modifier = C.KMOD_LSHIFT
	ModifierRShift   Modifier = C.KMOD_RSHIFT
	ModifierLCtrl    Modifier = C.KMOD_LCTRL
	ModifierRCtrl    Modifier = C.KMOD_RCTRL
	ModifierLAlt     Modifier = C.KMOD_LALT
	ModifierRAlt     Modifier = C.KMOD_RALT
	ModifierLGUI     Modifier = C.KMOD_LGUI
	ModifierRGUI     Modifier = C.KMOD_RGUI
	ModifierNum      Modifier = C.KMOD_NUM
	ModifierCaps     Modifier = C.KMOD_CAPS
	ModifierMode     Modifier = C.KMOD_MODE
	ModifierReserved Modifier = C.KMOD_RESERVED

	ModifierCtrl  Modifier = C.KMOD_CTRL
	ModifierShift Modifier = C.KMOD_SHIFT
	ModifierAlt   Modifier = C.KMOD_ALT
	ModifierGUI   Modifier = C.KMOD_GUI
)
