
package codes

const( 
    UNKNOWN = 0
    BACKSPACE = 8
    TAB = 9
    CLEAR = 12
    RETURN = 13
    PAUSE = 19
    ESCAPE = 27
    SPACE = 32
    QUOTEDDOUBLE = 34
    HASH = 35
    DOLLAR = 36
    PERCENT = 37
    AMPERSAND = 38
    QUOTE = 39
    LEFTPARENTHESIS = 40
    RIGHTPARENTHESIS = 41
    ASTERISK = 42
    PLUS = 43
    COMMA = 44
    MINUS = 45
    PERIOD = 46
    SLASH = 47
    ZERO = 48
    ONE = 49
    TWO = 50
    THREE = 51
    FOUR = 52
    FIVE = 53
    SIX = 54
    SEVEN = 55
    EIGHT = 56
    NINE = 57
    COLON = 58
    SEMICOLON = 59
    LESSTHAN = 60
    EQUALS = 61
    GREATERTHAN = 62
    QUESTION = 63
    AT = 64
    LEFTBRACKET = 91
    BACKSLASH = 92
    RIGHTBRACKET = 93
    CARET = 94
    UNDERSCORE = 95
    BACKQUOTE = 96
    A = 97
    B = 98
    C = 99
    D = 100
    E = 101
    F = 102
    G = 103
    H = 104
    I = 105
    J = 106
    K = 107
    L = 108
    M = 109
    N = 110
    O = 111
    P = 112
    Q = 113
    R = 114
    S = 115
    T = 116
    U = 117
    V = 118
    W = 119
    X = 120
    Y = 121
    Z = 122
    LEFTCURLY = 123
    PIPE = 124
    RIGHTCURLY = 125
    TILDE = 126
    DELETE = 127
    KEYPADZERO = 256
    KEYPADONE = 257
    KEYPADTWO = 258
    KEYPADTHREE = 259
    KEYPADFOUR = 260
    KEYPADFIVE = 261
    KEYPADSIX = 262
    KEYPADSEVEN = 263
    KEYPADEIGHT = 264
    KEYPADNINE = 265
    KEYPADPERIOD = 266
    KEYPADDIVIDE = 267
    KEYPADMULTIPLY = 268
    KEYPADMINUS = 269
    KEYPADPLUS = 270
    KEYPADENTER = 271
    KEYPADEQUALS = 272
    UP = 273
    DOWN = 274
    RIGHT = 275
    LEFT = 276
    INSERT = 277
    HOME = 278
    END = 279
    PAGEUP = 280
    PAGEDOWN = 281
    LEFTSHIFT = 304
    RIGHTSHIFT = 303
    LEFTMETA = 310
    RIGHTMETA = 309
    LEFTALT = 308
    RIGHTALT = 307
    LEFTCONTROL = 306
    RIGHTCONTROL = 305
    CAPSLOCK = 301
    NUMLOCK = 300
    SCROLLLOCK = 302
    LEFTSUPER = 311
    RIGHTSUPER = 312
    MODE = 313
    COMPOSE = 314
    HELP = 315
    PRINT = 316
    SYSREQ = 317
    BREAK = 318
    MENU = 319
    POWER = 320
    EURO = 321
    UNDO = 322
    F1 = 282
    F2 = 283
    F3 = 284
    F4 = 285
    F5 = 286
    F6 = 287
    F7 = 288
    F8 = 289
    F9 = 290
    F10 = 291
    F11 = 292
    F12 = 293
    F13 = 294
    F14 = 295
    F15 = 296
    WORLD0 = 160
    WORLD1 = 161
    WORLD2 = 162
    WORLD3 = 163
    WORLD4 = 164
    WORLD5 = 165
    WORLD6 = 166
    WORLD7 = 167
    WORLD8 = 168
    WORLD9 = 169
    WORLD10 = 170
    WORLD11 = 171
    WORLD12 = 172
    WORLD13 = 173
    WORLD14 = 174
    WORLD15 = 175
    WORLD16 = 176
    WORLD17 = 177
    WORLD18 = 178
    WORLD19 = 179
    WORLD20 = 180
    WORLD21 = 181
    WORLD22 = 182
    WORLD23 = 183
    WORLD24 = 184
    WORLD25 = 185
    WORLD26 = 186
    WORLD27 = 187
    WORLD28 = 188
    WORLD29 = 189
    WORLD30 = 190
    WORLD31 = 191
    WORLD32 = 192
    WORLD33 = 193
    WORLD34 = 194
    WORLD35 = 195
    WORLD36 = 196
    WORLD37 = 197
    WORLD38 = 198
    WORLD39 = 199
    WORLD40 = 200
    WORLD41 = 201
    WORLD42 = 202
    WORLD43 = 203
    WORLD44 = 204
    WORLD45 = 205
    WORLD46 = 206
    WORLD47 = 207
    WORLD48 = 208
    WORLD49 = 209
    WORLD50 = 210
    WORLD51 = 211
    WORLD52 = 212
    WORLD53 = 213
    WORLD54 = 214
    WORLD55 = 215
    WORLD56 = 216
    WORLD57 = 217
    WORLD58 = 218
    WORLD59 = 219
    WORLD60 = 220
    WORLD61 = 221
    WORLD62 = 222
    WORLD63 = 223
    WORLD64 = 224
    WORLD65 = 225
    WORLD66 = 226
    WORLD67 = 227
    WORLD68 = 228
    WORLD69 = 229
    WORLD70 = 230
    WORLD71 = 231
    WORLD72 = 232
    WORLD73 = 233
    WORLD74 = 234
    WORLD75 = 235
    WORLD76 = 236
    WORLD77 = 237
    WORLD78 = 238
    WORLD79 = 239
    WORLD80 = 240
    WORLD81 = 241
    WORLD82 = 242
    WORLD83 = 243
    WORLD84 = 244
    WORLD85 = 245
    WORLD86 = 246
    WORLD87 = 247
    WORLD88 = 248
    WORLD89 = 249
    WORLD90 = 250
    WORLD91 = 251
    WORLD92 = 252
    WORLD93 = 253
    WORLD94 = 254
    WORLD95 = 255
    BUTTONX = 1000
    BUTTONY = 1001
    BUTTONA = 1002
    BUTTONB = 1003
    BUTTONR1 = 1004
    BUTTONL1 = 1005
    BUTTONR2 = 1006
    BUTTONL2 = 1007
    BUTTONR3 = 1008
    BUTTONL3 = 1009
    BUTTONSTART = 1010
    BUTTONSELECT = 1011
    DPADLEFT = 1012
    DPADRIGHT = 1013
    DPADUP = 1014
    DPADDOWN = 1015
    THUMBSTICK1 = 1016
    THUMBSTICK2 = 1017
)
var keys = map[string]int{
    "unknown":UNKNOWN,
    "backspace":BACKSPACE,
    "tab":TAB,
    "clear":CLEAR,
    "return":RETURN,
    "pause":PAUSE,
    "escape":ESCAPE,
    "space":SPACE,
    "quoteddouble":QUOTEDDOUBLE,
    "hash":HASH,
    "dollar":DOLLAR,
    "percent":PERCENT,
    "ampersand":AMPERSAND,
    "quote":QUOTE,
    "leftparenthesis":LEFTPARENTHESIS,
    "rightparenthesis":RIGHTPARENTHESIS,
    "asterisk":ASTERISK,
    "plus":PLUS,
    "comma":COMMA,
    "minus":MINUS,
    "period":PERIOD,
    "slash":SLASH,
    "zero":ZERO,
    "one":ONE,
    "two":TWO,
    "three":THREE,
    "four":FOUR,
    "five":FIVE,
    "six":SIX,
    "seven":SEVEN,
    "eight":EIGHT,
    "nine":NINE,
    "colon":COLON,
    "semicolon":SEMICOLON,
    "lessthan":LESSTHAN,
    "equals":EQUALS,
    "greaterthan":GREATERTHAN,
    "question":QUESTION,
    "at":AT,
    "leftbracket":LEFTBRACKET,
    "backslash":BACKSLASH,
    "rightbracket":RIGHTBRACKET,
    "caret":CARET,
    "underscore":UNDERSCORE,
    "backquote":BACKQUOTE,
    "a":A,
    "b":B,
    "c":C,
    "d":D,
    "e":E,
    "f":F,
    "g":G,
    "h":H,
    "i":I,
    "j":J,
    "k":K,
    "l":L,
    "m":M,
    "n":N,
    "o":O,
    "p":P,
    "q":Q,
    "r":R,
    "s":S,
    "t":T,
    "u":U,
    "v":V,
    "w":W,
    "x":X,
    "y":Y,
    "z":Z,
    "leftcurly":LEFTCURLY,
    "pipe":PIPE,
    "rightcurly":RIGHTCURLY,
    "tilde":TILDE,
    "delete":DELETE,
    "keypadzero":KEYPADZERO,
    "keypadone":KEYPADONE,
    "keypadtwo":KEYPADTWO,
    "keypadthree":KEYPADTHREE,
    "keypadfour":KEYPADFOUR,
    "keypadfive":KEYPADFIVE,
    "keypadsix":KEYPADSIX,
    "keypadseven":KEYPADSEVEN,
    "keypadeight":KEYPADEIGHT,
    "keypadnine":KEYPADNINE,
    "keypadperiod":KEYPADPERIOD,
    "keypaddivide":KEYPADDIVIDE,
    "keypadmultiply":KEYPADMULTIPLY,
    "keypadminus":KEYPADMINUS,
    "keypadplus":KEYPADPLUS,
    "keypadenter":KEYPADENTER,
    "keypadequals":KEYPADEQUALS,
    "up":UP,
    "down":DOWN,
    "right":RIGHT,
    "left":LEFT,
    "insert":INSERT,
    "home":HOME,
    "end":END,
    "pageup":PAGEUP,
    "pagedown":PAGEDOWN,
    "leftshift":LEFTSHIFT,
    "rightshift":RIGHTSHIFT,
    "leftmeta":LEFTMETA,
    "rightmeta":RIGHTMETA,
    "leftalt":LEFTALT,
    "rightalt":RIGHTALT,
    "leftcontrol":LEFTCONTROL,
    "rightcontrol":RIGHTCONTROL,
    "capslock":CAPSLOCK,
    "numlock":NUMLOCK,
    "scrolllock":SCROLLLOCK,
    "leftsuper":LEFTSUPER,
    "rightsuper":RIGHTSUPER,
    "mode":MODE,
    "compose":COMPOSE,
    "help":HELP,
    "print":PRINT,
    "sysreq":SYSREQ,
    "break":BREAK,
    "menu":MENU,
    "power":POWER,
    "euro":EURO,
    "undo":UNDO,
    "f1":F1,
    "f2":F2,
    "f3":F3,
    "f4":F4,
    "f5":F5,
    "f6":F6,
    "f7":F7,
    "f8":F8,
    "f9":F9,
    "f10":F10,
    "f11":F11,
    "f12":F12,
    "f13":F13,
    "f14":F14,
    "f15":F15,
    "world0":WORLD0,
    "world1":WORLD1,
    "world2":WORLD2,
    "world3":WORLD3,
    "world4":WORLD4,
    "world5":WORLD5,
    "world6":WORLD6,
    "world7":WORLD7,
    "world8":WORLD8,
    "world9":WORLD9,
    "world10":WORLD10,
    "world11":WORLD11,
    "world12":WORLD12,
    "world13":WORLD13,
    "world14":WORLD14,
    "world15":WORLD15,
    "world16":WORLD16,
    "world17":WORLD17,
    "world18":WORLD18,
    "world19":WORLD19,
    "world20":WORLD20,
    "world21":WORLD21,
    "world22":WORLD22,
    "world23":WORLD23,
    "world24":WORLD24,
    "world25":WORLD25,
    "world26":WORLD26,
    "world27":WORLD27,
    "world28":WORLD28,
    "world29":WORLD29,
    "world30":WORLD30,
    "world31":WORLD31,
    "world32":WORLD32,
    "world33":WORLD33,
    "world34":WORLD34,
    "world35":WORLD35,
    "world36":WORLD36,
    "world37":WORLD37,
    "world38":WORLD38,
    "world39":WORLD39,
    "world40":WORLD40,
    "world41":WORLD41,
    "world42":WORLD42,
    "world43":WORLD43,
    "world44":WORLD44,
    "world45":WORLD45,
    "world46":WORLD46,
    "world47":WORLD47,
    "world48":WORLD48,
    "world49":WORLD49,
    "world50":WORLD50,
    "world51":WORLD51,
    "world52":WORLD52,
    "world53":WORLD53,
    "world54":WORLD54,
    "world55":WORLD55,
    "world56":WORLD56,
    "world57":WORLD57,
    "world58":WORLD58,
    "world59":WORLD59,
    "world60":WORLD60,
    "world61":WORLD61,
    "world62":WORLD62,
    "world63":WORLD63,
    "world64":WORLD64,
    "world65":WORLD65,
    "world66":WORLD66,
    "world67":WORLD67,
    "world68":WORLD68,
    "world69":WORLD69,
    "world70":WORLD70,
    "world71":WORLD71,
    "world72":WORLD72,
    "world73":WORLD73,
    "world74":WORLD74,
    "world75":WORLD75,
    "world76":WORLD76,
    "world77":WORLD77,
    "world78":WORLD78,
    "world79":WORLD79,
    "world80":WORLD80,
    "world81":WORLD81,
    "world82":WORLD82,
    "world83":WORLD83,
    "world84":WORLD84,
    "world85":WORLD85,
    "world86":WORLD86,
    "world87":WORLD87,
    "world88":WORLD88,
    "world89":WORLD89,
    "world90":WORLD90,
    "world91":WORLD91,
    "world92":WORLD92,
    "world93":WORLD93,
    "world94":WORLD94,
    "world95":WORLD95,
    "buttonx":BUTTONX,
    "buttony":BUTTONY,
    "buttona":BUTTONA,
    "buttonb":BUTTONB,
    "buttonr1":BUTTONR1,
    "buttonl1":BUTTONL1,
    "buttonr2":BUTTONR2,
    "buttonl2":BUTTONL2,
    "buttonr3":BUTTONR3,
    "buttonl3":BUTTONL3,
    "buttonstart":BUTTONSTART,
    "buttonselect":BUTTONSELECT,
    "dpadleft":DPADLEFT,
    "dpadright":DPADRIGHT,
    "dpadup":DPADUP,
    "dpaddown":DPADDOWN,
    "thumbstick1":THUMBSTICK1,
    "thumbstick2":THUMBSTICK2,
}

func GetKey(name string)int{
    if k, ok := keys[name]; ok{
        return k
    }
    return -1
}
