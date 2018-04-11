// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// A keyCode designate a key by its physical position on the keyboard.
// It is not affected by the layout or any other language settings.
type keyCode = internal.KeyCode

// keyCode constants, following the USB HID usage table (page 0x07).
// The name correspond to the key label in that position on a US QWERTY
// layout.
const (
	keyUnknown keyCode = 0

	keyA keyCode = 4
	keyB keyCode = 5
	keyC keyCode = 6
	keyD keyCode = 7
	keyE keyCode = 8
	keyF keyCode = 9
	keyG keyCode = 10
	keyH keyCode = 11
	keyI keyCode = 12
	keyJ keyCode = 13
	keyK keyCode = 14
	keyL keyCode = 15
	keyM keyCode = 16
	keyN keyCode = 17
	keyO keyCode = 18
	keyP keyCode = 19
	keyQ keyCode = 20
	keyR keyCode = 21
	keyS keyCode = 22
	keyT keyCode = 23
	keyU keyCode = 24
	keyV keyCode = 25
	keyW keyCode = 26
	keyX keyCode = 27
	keyY keyCode = 28
	keyZ keyCode = 29

	key1 keyCode = 30
	key2 keyCode = 31
	key3 keyCode = 32
	key4 keyCode = 33
	key5 keyCode = 34
	key6 keyCode = 35
	key7 keyCode = 36
	key8 keyCode = 37
	key9 keyCode = 38
	key0 keyCode = 39

	keyReturn    keyCode = 40
	keyEscape    keyCode = 41
	keyBackspace keyCode = 42
	keyTab       keyCode = 43
	keySpace     keyCode = 44

	keyMinus        keyCode = 45
	keyEquals       keyCode = 46
	keyLeftBracket  keyCode = 47
	keyRightBracket keyCode = 48
	keyBackSlash    keyCode = 49
	keyNonSHash     keyCode = 50
	keySemicolon    keyCode = 51
	keyApostrophe   keyCode = 52
	keyGrave        keyCode = 53
	keyComma        keyCode = 54
	keyPeriod       keyCode = 55
	keySlash        keyCode = 56

	keyCapsLock keyCode = 57

	keyF1  keyCode = 58
	keyF2  keyCode = 59
	keyF3  keyCode = 60
	keyF4  keyCode = 61
	keyF5  keyCode = 62
	keyF6  keyCode = 63
	keyF7  keyCode = 64
	keyF8  keyCode = 65
	keyF9  keyCode = 66
	keyF10 keyCode = 67
	keyF11 keyCode = 68
	keyF12 keyCode = 69

	keyPrintScreen keyCode = 70
	keyScrollLock  keyCode = 71
	keyPause       keyCode = 72
	keyInsert      keyCode = 73
	keyHome        keyCode = 74
	keyPageUp      keyCode = 75
	keyDelete      keyCode = 76
	keyEnd         keyCode = 77
	keyPageDown    keyCode = 78
	keyRight       keyCode = 79
	keyLeft        keyCode = 80
	keyDown        keyCode = 81
	keyUp          keyCode = 82

	keyNumLockClear keyCode = 83
	keyKPDivide     keyCode = 84
	keyKPMultiply   keyCode = 85
	keyKPMinus      keyCode = 86
	keyKPPlus       keyCode = 87
	keyKPEnter      keyCode = 88
	keyKP1          keyCode = 89
	keyKP2          keyCode = 90
	keyKP3          keyCode = 91
	keyKP4          keyCode = 92
	keyKP5          keyCode = 93
	keyKP6          keyCode = 94
	keyKP7          keyCode = 95
	keyKP8          keyCode = 96
	keyKP9          keyCode = 97
	keyKP0          keyCode = 98
	keyKPPeriod     keyCode = 99

	keyNonUSBackSlash keyCode = 100
	keyApplication    keyCode = 101
	keyPower          keyCode = 102
	keyKPEquals       keyCode = 103
	keyF13            keyCode = 104
	keyF14            keyCode = 105
	keyF15            keyCode = 106
	keyF16            keyCode = 107
	keyF17            keyCode = 108
	keyF18            keyCode = 109
	keyF19            keyCode = 110
	keyF20            keyCode = 111
	keyF21            keyCode = 112
	keyF22            keyCode = 113
	keyF23            keyCode = 114
	keyF24            keyCode = 115
	keyExecute        keyCode = 116
	keyHelp           keyCode = 117
	keyMenu           keyCode = 118
	keySelect         keyCode = 119
	keyStop           keyCode = 120
	keyAgain          keyCode = 121
	keyUndo           keyCode = 122
	keyCut            keyCode = 123
	keyCopy           keyCode = 124
	keyPaste          keyCode = 125
	keyFind           keyCode = 126
	keyMute           keyCode = 127
	keyVolumeUp       keyCode = 128
	keyVolumeDown     keyCode = 129

	keyKPComma       keyCode = 133
	keyKPEqualsAS400 keyCode = 134

	keyInternational1 keyCode = 135
	keyInternational2 keyCode = 136
	keyInternational3 keyCode = 137
	keyInternational4 keyCode = 138
	keyInternational5 keyCode = 139
	keyInternational6 keyCode = 140
	keyInternational7 keyCode = 141
	keyInternational8 keyCode = 142
	keyInternational9 keyCode = 143
	keyLang1          keyCode = 144
	keyLang2          keyCode = 145
	keyLang3          keyCode = 146
	keyLang4          keyCode = 14
	keyLang5          keyCode = 148
	keyLang6          keyCode = 149
	keyLang7          keyCode = 150
	keyLang8          keyCode = 151
	keyLang9          keyCode = 152

	keyAltErase   keyCode = 153
	keySysReq     keyCode = 154
	keyCancel     keyCode = 155
	keyClear      keyCode = 156
	keyPrior      keyCode = 157
	keyReturn2    keyCode = 158
	keySeparator  keyCode = 159
	keyOut        keyCode = 160
	keyOper       keyCode = 161
	keyClearAgain keyCode = 162
	keyCrSel      keyCode = 163
	keyExSel      keyCode = 164

	keyKP00               keyCode = 176
	keyKP000              keyCode = 177
	keyThousandsSeparator keyCode = 178
	keyDecimalSeparator   keyCode = 179
	keyCurrencyUnit       keyCode = 180
	keyCurrencySubUnit    keyCode = 181
	keyKPLeftParen        keyCode = 182
	keyKPRightParen       keyCode = 183
	keyKPLeftBrace        keyCode = 184
	keyKPRightBrace       keyCode = 185
	keyKPTab              keyCode = 186
	keyKPBackspace        keyCode = 187
	keyKPA                keyCode = 188
	keyKPB                keyCode = 189
	keyKPC                keyCode = 190
	keyKPD                keyCode = 191
	keyKPE                keyCode = 192
	keyKPF                keyCode = 193
	keyKPXor              keyCode = 194
	keyKPPower            keyCode = 195
	keyKPPercent          keyCode = 196
	keyKPLess             keyCode = 197
	keyKPGreater          keyCode = 198
	keyKPAmpersand        keyCode = 199
	keyKPDblAmpersand     keyCode = 200
	keyKPVerticalBar      keyCode = 201
	keyKPDblVerticalBar   keyCode = 202
	keyKPColon            keyCode = 203
	keyKPHash             keyCode = 204
	keyKPSpace            keyCode = 205
	keyKPAt               keyCode = 206
	keyKPExclam           keyCode = 207
	keyKPMemStore         keyCode = 208
	keyKPMemRecall        keyCode = 209
	keyKPMemClear         keyCode = 210
	keyKPMemAdd           keyCode = 211
	keyKPMemSubtract      keyCode = 212
	keyKPMemMultiply      keyCode = 213
	keyKPMemDivide        keyCode = 214
	keyKPPlusMinus        keyCode = 215
	keyKPClear            keyCode = 216
	keyKPClearEntry       keyCode = 217
	keyKPBinary           keyCode = 218
	keyKPOctal            keyCode = 219
	keyKPDecimal          keyCode = 220
	keyKPHexadecimal      keyCode = 221

	keyLCtrl  keyCode = 224
	keyLShift keyCode = 225
	keyLAlt   keyCode = 226
	keyLGUI   keyCode = 227
	keyRCtrl  keyCode = 228
	keyRShift keyCode = 229
	keyRAlt   keyCode = 230
	keyRGUI   keyCode = 231
)

// Additional positions constants (from page 0x0C and ?)
const (
	keyMode keyCode = 257

	keyAudioNext   keyCode = 258
	keyAudioPrev   keyCode = 259
	keyAudioStop   keyCode = 260
	keyAudioPlay   keyCode = 261
	keyAudioMute   keyCode = 262
	keyMediaSelect keyCode = 263
	keyWWW         keyCode = 264
	keyMail        keyCode = 265
	keyCalculator  keyCode = 266
	keyComputer    keyCode = 267
	keyACSearch    keyCode = 268
	keyACHome      keyCode = 269
	keyACBack      keyCode = 270
	keyACForward   keyCode = 271
	keyACStop      keyCode = 272
	keyACRefresh   keyCode = 273
	keyACBookmarks keyCode = 274

	keyBrightnessDown keyCode = 275
	keyBrightnessUp   keyCode = 276
	keyDisplaySwitch  keyCode = 277
	keyKbdIllumToggle keyCode = 278
	keyKbdIllumDown   keyCode = 279
	keyKbdIllumUp     keyCode = 280
	keyEject          keyCode = 281
	keySleep          keyCode = 282

	keyApp1 keyCode = 283
	keyApp2 keyCode = 284
)

// maxKey is the maximum valid position.
const maxKey keyCode = 512

////////////////////////////////////////////////////////////////////////////////
