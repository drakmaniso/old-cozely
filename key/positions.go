package key

// #include "../engine/engine.h"
import "C"

// A Position designate a key by its physical position on the keyboard.
// It is not affected by the layout or any other language settings.
type Position uint32

// Position constants, following the USB HID usage table (page 0x07).
// The name correspond to the key label in that position on a US QWERTY
// layout.
const (
	PositionUnknown Position = 0

	PositionA Position = 4
	PositionB Position = 5
	PositionC Position = 6
	PositionD Position = 7
	PositionE Position = 8
	PositionF Position = 9
	PositionG Position = 10
	PositionH Position = 11
	PositionI Position = 12
	PositionJ Position = 13
	PositionK Position = 14
	PositionL Position = 15
	PositionM Position = 16
	PositionN Position = 17
	PositionO Position = 18
	PositionP Position = 19
	PositionQ Position = 20
	PositionR Position = 21
	PositionS Position = 22
	PositionT Position = 23
	PositionU Position = 24
	PositionV Position = 25
	PositionW Position = 26
	PositionX Position = 27
	PositionY Position = 28
	PositionZ Position = 29

	Position1 Position = 30
	Position2 Position = 31
	Position3 Position = 32
	Position4 Position = 33
	Position5 Position = 34
	Position6 Position = 35
	Position7 Position = 36
	Position8 Position = 37
	Position9 Position = 38
	Position0 Position = 39

	PositionReturn    Position = 40
	PositionEscape    Position = 41
	PositionBackspace Position = 42
	PositionTab       Position = 43
	PositionSpace     Position = 44

	PositionMinus        Position = 45
	PositionEquals       Position = 46
	PositionLeftBracket  Position = 47
	PositionRightBracket Position = 48
	PositionBackSlash    Position = 49
	PositionNonSHash     Position = 50
	PositionSemicolon    Position = 51
	PositionApostrophe   Position = 52
	PositionGrave        Position = 53
	PositionComma        Position = 54
	PositionPeriod       Position = 55
	PositionSlash        Position = 56

	PositionCapsLock Position = 57

	PositionF1  Position = 58
	PositionF2  Position = 59
	PositionF3  Position = 60
	PositionF4  Position = 61
	PositionF5  Position = 62
	PositionF6  Position = 63
	PositionF7  Position = 64
	PositionF8  Position = 65
	PositionF9  Position = 66
	PositionF10 Position = 67
	PositionF11 Position = 68
	PositionF12 Position = 69

	PositionPrintScreen Position = 70
	PositionScrollLock  Position = 71
	PositionPause       Position = 72
	PositionInsert      Position = 73
	PositionHome        Position = 74
	PositionPageUp      Position = 75
	PositionDelete      Position = 76
	PositionEnd         Position = 77
	PositionPageDown    Position = 78
	PositionRight       Position = 79
	PositionLeft        Position = 80
	PositionDown        Position = 81
	PositionUp          Position = 82

	PositionNumLockClear Position = 83
	PositionKPDivide     Position = 84
	PositionKPMultiply   Position = 85
	PositionKPMinus      Position = 86
	PositionKPPlus       Position = 87
	PositionKPEnter      Position = 88
	PositionKP1          Position = 89
	PositionKP2          Position = 90
	PositionKP3          Position = 91
	PositionKP4          Position = 92
	PositionKP5          Position = 93
	PositionKP6          Position = 94
	PositionKP7          Position = 95
	PositionKP8          Position = 96
	PositionKP9          Position = 97
	PositionKP0          Position = 98
	PositionKPPeriod     Position = 99

	PositionNonUSBackSlash Position = 100
	PositionApplication    Position = 101
	PositionPower          Position = 102
	PositionKPEquals       Position = 103
	PositionF13            Position = 104
	PositionF14            Position = 105
	PositionF15            Position = 106
	PositionF16            Position = 107
	PositionF17            Position = 108
	PositionF18            Position = 109
	PositionF19            Position = 110
	PositionF20            Position = 111
	PositionF21            Position = 112
	PositionF22            Position = 113
	PositionF23            Position = 114
	PositionF24            Position = 115
	PositionExecute        Position = 116
	PositionHelp           Position = 117
	PositionMenu           Position = 118
	PositionSelect         Position = 119
	PositionStop           Position = 120
	PositionAgain          Position = 121
	PositionUndo           Position = 122
	PositionCut            Position = 123
	PositionCopy           Position = 124
	PositionPaste          Position = 125
	PositionFind           Position = 126
	PositionMute           Position = 127
	PositionVolumeUp       Position = 128
	PositionVolumeDown     Position = 129

	PositionKPComma       Position = 133
	PositionKPEqualsAS400 Position = 134

	PositionInternational1 Position = 135
	PositionInternational2 Position = 136
	PositionInternational3 Position = 137
	PositionInternational4 Position = 138
	PositionInternational5 Position = 139
	PositionInternational6 Position = 140
	PositionInternational7 Position = 141
	PositionInternational8 Position = 142
	PositionInternational9 Position = 143
	PositionLang1          Position = 144
	PositionLang2          Position = 145
	PositionLang3          Position = 146
	PositionLang4          Position = 14
	PositionLang5          Position = 148
	PositionLang6          Position = 149
	PositionLang7          Position = 150
	PositionLang8          Position = 151
	PositionLang9          Position = 152

	PositionAltErase   Position = 153
	PositionSysReq     Position = 154
	PositionCancel     Position = 155
	PositionClear      Position = 156
	PositionPrior      Position = 157
	PositionReturn2    Position = 158
	PositionSeparator  Position = 159
	PositionOut        Position = 160
	PositionOper       Position = 161
	PositionClearAgain Position = 162
	PositionCrSel      Position = 163
	PositionExSel      Position = 164

	PositionKP00               Position = 176
	PositionKP000              Position = 177
	PositionThousandsSeparator Position = 178
	PositionDecimalSeparator   Position = 179
	PositionCurrencyUnit       Position = 180
	PositionCurrencySubUnit    Position = 181
	PositionKPLeftParen        Position = 182
	PositionKPRightParen       Position = 183
	PositionKPLeftBrace        Position = 184
	PositionKPRightBrace       Position = 185
	PositionKPTab              Position = 186
	PositionKPBackspace        Position = 187
	PositionKPA                Position = 188
	PositionKPB                Position = 189
	PositionKPC                Position = 190
	PositionKPD                Position = 191
	PositionKPE                Position = 192
	PositionKPF                Position = 193
	PositionKPXor              Position = 194
	PositionKPPower            Position = 195
	PositionKPPercent          Position = 196
	PositionKPLess             Position = 197
	PositionKPGreater          Position = 198
	PositionKPAmpersand        Position = 199
	PositionKPDblAmpersand     Position = 200
	PositionKPVerticalBar      Position = 201
	PositionKPDblVerticalBar   Position = 202
	PositionKPColon            Position = 203
	PositionKPHash             Position = 204
	PositionKPSpace            Position = 205
	PositionKPAt               Position = 206
	PositionKPExclam           Position = 207
	PositionKPMemStore         Position = 208
	PositionKPMemRecall        Position = 209
	PositionKPMemClear         Position = 210
	PositionKPMemAdd           Position = 211
	PositionKPMemSubtract      Position = 212
	PositionKPMemMultiply      Position = 213
	PositionKPMemDivide        Position = 214
	PositionKPPlusMinus        Position = 215
	PositionKPClear            Position = 216
	PositionKPClearEntry       Position = 217
	PositionKPBinary           Position = 218
	PositionKPOctal            Position = 219
	PositionKPDecimal          Position = 220
	PositionKPHexadecimal      Position = 221

	PositionLCtrl  Position = 224
	PositionLShift Position = 225
	PositionLAlt   Position = 226
	PositionLGUI   Position = 227
	PositionRCtrl  Position = 228
	PositionRShift Position = 229
	PositionRAlt   Position = 230
	PositionRGUI   Position = 231
)

// Additional positions constants (from page 0x0C and ?)
const (
	PositionMode Position = 257

	PositionAudioNext   Position = 258
	PositionAudioPrev   Position = 259
	PositionAudioStop   Position = 260
	PositionAudioPlay   Position = 261
	PositionAudioMute   Position = 262
	PositionMediaSelect Position = 263
	PositionWWW         Position = 264
	PositionMail        Position = 265
	PositionCalculator  Position = 266
	PositionComputer    Position = 267
	PositionACSearch    Position = 268
	PositionACHome      Position = 269
	PositionACBack      Position = 270
	PositionACForward   Position = 271
	PositionACStop      Position = 272
	PositionACRefresh   Position = 273
	PositionACBookmarks Position = 274

	PositionBrightnessDown Position = 275
	PositionBrightnessUp   Position = 276
	PositionDisplaySwitch  Position = 277
	PositionKbdillumToggle Position = 278
	PositionKbdillumDown   Position = 279
	PositionKbdillumUp     Position = 280
	PositionEject          Position = 281
	PositionSleep          Position = 282

	PositionApp1 Position = 283
	PositionApp2 Position = 284
)

// MaxPosition is the maximum valid position.
const MaxPosition Position = C.SDL_NUM_SCANCODES
