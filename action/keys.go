// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

import (
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// A KeyCode designate a key by its physical position on the keyboard.
// It is not affected by the layout or any other language settings.
type KeyCode = internal.KeyCode

// KeyCode constants, following the USB HID usage table (page 0x07).
// The name correspond to the key label in that position on a US QWERTY
// layout.
const (
	KeyUnknown KeyCode = 0

	KeyA KeyCode = 4
	KeyB KeyCode = 5
	KeyC KeyCode = 6
	KeyD KeyCode = 7
	KeyE KeyCode = 8
	KeyF KeyCode = 9
	KeyG KeyCode = 10
	KeyH KeyCode = 11
	KeyI KeyCode = 12
	KeyJ KeyCode = 13
	KeyK KeyCode = 14
	KeyL KeyCode = 15
	KeyM KeyCode = 16
	KeyN KeyCode = 17
	KeyO KeyCode = 18
	KeyP KeyCode = 19
	KeyQ KeyCode = 20
	KeyR KeyCode = 21
	KeyS KeyCode = 22
	KeyT KeyCode = 23
	KeyU KeyCode = 24
	KeyV KeyCode = 25
	KeyW KeyCode = 26
	KeyX KeyCode = 27
	KeyY KeyCode = 28
	KeyZ KeyCode = 29

	Key1 KeyCode = 30
	Key2 KeyCode = 31
	Key3 KeyCode = 32
	Key4 KeyCode = 33
	Key5 KeyCode = 34
	Key6 KeyCode = 35
	Key7 KeyCode = 36
	Key8 KeyCode = 37
	Key9 KeyCode = 38
	Key0 KeyCode = 39

	KeyReturn    KeyCode = 40
	KeyEscape    KeyCode = 41
	KeyBackspace KeyCode = 42
	KeyTab       KeyCode = 43
	KeySpace     KeyCode = 44

	KeyMinus        KeyCode = 45
	KeyEquals       KeyCode = 46
	KeyLeftBracket  KeyCode = 47
	KeyRightBracket KeyCode = 48
	KeyBackSlash    KeyCode = 49
	KeyNonSHash     KeyCode = 50
	KeySemicolon    KeyCode = 51
	KeyApostrophe   KeyCode = 52
	KeyGrave        KeyCode = 53
	KeyComma        KeyCode = 54
	KeyPeriod       KeyCode = 55
	KeySlash        KeyCode = 56

	KeyCapsLock KeyCode = 57

	KeyF1  KeyCode = 58
	KeyF2  KeyCode = 59
	KeyF3  KeyCode = 60
	KeyF4  KeyCode = 61
	KeyF5  KeyCode = 62
	KeyF6  KeyCode = 63
	KeyF7  KeyCode = 64
	KeyF8  KeyCode = 65
	KeyF9  KeyCode = 66
	KeyF10 KeyCode = 67
	KeyF11 KeyCode = 68
	KeyF12 KeyCode = 69

	KeyPrintScreen KeyCode = 70
	KeyScrollLock  KeyCode = 71
	KeyPause       KeyCode = 72
	KeyInsert      KeyCode = 73
	KeyHome        KeyCode = 74
	KeyPageUp      KeyCode = 75
	KeyDelete      KeyCode = 76
	KeyEnd         KeyCode = 77
	KeyPageDown    KeyCode = 78
	KeyRight       KeyCode = 79
	KeyLeft        KeyCode = 80
	KeyDown        KeyCode = 81
	KeyUp          KeyCode = 82

	KeyNumLockClear KeyCode = 83
	KeyKPDivide     KeyCode = 84
	KeyKPMultiply   KeyCode = 85
	KeyKPMinus      KeyCode = 86
	KeyKPPlus       KeyCode = 87
	KeyKPEnter      KeyCode = 88
	KeyKP1          KeyCode = 89
	KeyKP2          KeyCode = 90
	KeyKP3          KeyCode = 91
	KeyKP4          KeyCode = 92
	KeyKP5          KeyCode = 93
	KeyKP6          KeyCode = 94
	KeyKP7          KeyCode = 95
	KeyKP8          KeyCode = 96
	KeyKP9          KeyCode = 97
	KeyKP0          KeyCode = 98
	KeyKPPeriod     KeyCode = 99

	KeyNonUSBackSlash KeyCode = 100
	KeyApplication    KeyCode = 101
	KeyPower          KeyCode = 102
	KeyKPEquals       KeyCode = 103
	KeyF13            KeyCode = 104
	KeyF14            KeyCode = 105
	KeyF15            KeyCode = 106
	KeyF16            KeyCode = 107
	KeyF17            KeyCode = 108
	KeyF18            KeyCode = 109
	KeyF19            KeyCode = 110
	KeyF20            KeyCode = 111
	KeyF21            KeyCode = 112
	KeyF22            KeyCode = 113
	KeyF23            KeyCode = 114
	KeyF24            KeyCode = 115
	KeyExecute        KeyCode = 116
	KeyHelp           KeyCode = 117
	KeyMenu           KeyCode = 118
	KeySelect         KeyCode = 119
	KeyStop           KeyCode = 120
	KeyAgain          KeyCode = 121
	KeyUndo           KeyCode = 122
	KeyCut            KeyCode = 123
	KeyCopy           KeyCode = 124
	KeyPaste          KeyCode = 125
	KeyFind           KeyCode = 126
	KeyMute           KeyCode = 127
	KeyVolumeUp       KeyCode = 128
	KeyVolumeDown     KeyCode = 129

	KeyKPComma       KeyCode = 133
	KeyKPEqualsAS400 KeyCode = 134

	KeyInternational1 KeyCode = 135
	KeyInternational2 KeyCode = 136
	KeyInternational3 KeyCode = 137
	KeyInternational4 KeyCode = 138
	KeyInternational5 KeyCode = 139
	KeyInternational6 KeyCode = 140
	KeyInternational7 KeyCode = 141
	KeyInternational8 KeyCode = 142
	KeyInternational9 KeyCode = 143
	KeyLang1          KeyCode = 144
	KeyLang2          KeyCode = 145
	KeyLang3          KeyCode = 146
	KeyLang4          KeyCode = 14
	KeyLang5          KeyCode = 148
	KeyLang6          KeyCode = 149
	KeyLang7          KeyCode = 150
	KeyLang8          KeyCode = 151
	KeyLang9          KeyCode = 152

	KeyAltErase   KeyCode = 153
	KeySysReq     KeyCode = 154
	KeyCancel     KeyCode = 155
	KeyClear      KeyCode = 156
	KeyPrior      KeyCode = 157
	KeyReturn2    KeyCode = 158
	KeySeparator  KeyCode = 159
	KeyOut        KeyCode = 160
	KeyOper       KeyCode = 161
	KeyClearAgain KeyCode = 162
	KeyCrSel      KeyCode = 163
	KeyExSel      KeyCode = 164

	KeyKP00               KeyCode = 176
	KeyKP000              KeyCode = 177
	KeyThousandsSeparator KeyCode = 178
	KeyDecimalSeparator   KeyCode = 179
	KeyCurrencyUnit       KeyCode = 180
	KeyCurrencySubUnit    KeyCode = 181
	KeyKPLeftParen        KeyCode = 182
	KeyKPRightParen       KeyCode = 183
	KeyKPLeftBrace        KeyCode = 184
	KeyKPRightBrace       KeyCode = 185
	KeyKPTab              KeyCode = 186
	KeyKPBackspace        KeyCode = 187
	KeyKPA                KeyCode = 188
	KeyKPB                KeyCode = 189
	KeyKPC                KeyCode = 190
	KeyKPD                KeyCode = 191
	KeyKPE                KeyCode = 192
	KeyKPF                KeyCode = 193
	KeyKPXor              KeyCode = 194
	KeyKPPower            KeyCode = 195
	KeyKPPercent          KeyCode = 196
	KeyKPLess             KeyCode = 197
	KeyKPGreater          KeyCode = 198
	KeyKPAmpersand        KeyCode = 199
	KeyKPDblAmpersand     KeyCode = 200
	KeyKPVerticalBar      KeyCode = 201
	KeyKPDblVerticalBar   KeyCode = 202
	KeyKPColon            KeyCode = 203
	KeyKPHash             KeyCode = 204
	KeyKPSpace            KeyCode = 205
	KeyKPAt               KeyCode = 206
	KeyKPExclam           KeyCode = 207
	KeyKPMemStore         KeyCode = 208
	KeyKPMemRecall        KeyCode = 209
	KeyKPMemClear         KeyCode = 210
	KeyKPMemAdd           KeyCode = 211
	KeyKPMemSubtract      KeyCode = 212
	KeyKPMemMultiply      KeyCode = 213
	KeyKPMemDivide        KeyCode = 214
	KeyKPPlusMinus        KeyCode = 215
	KeyKPClear            KeyCode = 216
	KeyKPClearEntry       KeyCode = 217
	KeyKPBinary           KeyCode = 218
	KeyKPOctal            KeyCode = 219
	KeyKPDecimal          KeyCode = 220
	KeyKPHexadecimal      KeyCode = 221

	KeyLCtrl  KeyCode = 224
	KeyLShift KeyCode = 225
	KeyLAlt   KeyCode = 226
	KeyLGUI   KeyCode = 227
	KeyRCtrl  KeyCode = 228
	KeyRShift KeyCode = 229
	KeyRAlt   KeyCode = 230
	KeyRGUI   KeyCode = 231
)

// Additional positions constants (from page 0x0C and ?)
const (
	KeyMode KeyCode = 257

	KeyAudioNext   KeyCode = 258
	KeyAudioPrev   KeyCode = 259
	KeyAudioStop   KeyCode = 260
	KeyAudioPlay   KeyCode = 261
	KeyAudioMute   KeyCode = 262
	KeyMediaSelect KeyCode = 263
	KeyWWW         KeyCode = 264
	KeyMail        KeyCode = 265
	KeyCalculator  KeyCode = 266
	KeyComputer    KeyCode = 267
	KeyACSearch    KeyCode = 268
	KeyACHome      KeyCode = 269
	KeyACBack      KeyCode = 270
	KeyACForward   KeyCode = 271
	KeyACStop      KeyCode = 272
	KeyACRefresh   KeyCode = 273
	KeyACBookmarks KeyCode = 274

	KeyBrightnessDown KeyCode = 275
	KeyBrightnessUp   KeyCode = 276
	KeyDisplaySwitch  KeyCode = 277
	KeyKbdIllumToggle KeyCode = 278
	KeyKbdIllumDown   KeyCode = 279
	KeyKbdIllumUp     KeyCode = 280
	KeyEject          KeyCode = 281
	KeySleep          KeyCode = 282

	KeyApp1 KeyCode = 283
	KeyApp2 KeyCode = 284
)

// MaxKey is the maximum valid position.
const MaxKey KeyCode = 512

//------------------------------------------------------------------------------
