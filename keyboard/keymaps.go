package keyboard

import "github.com/jbdemonte/virtual-device/linux"

// https://kbdlayout.info/kbdusx
var qwertyKeyMap = KeyMap{
	// Letters
	'a': {linux.KEY_A, false, false}, 'A': {linux.KEY_A, true, false},
	'b': {linux.KEY_B, false, false}, 'B': {linux.KEY_B, true, false},
	'c': {linux.KEY_C, false, false}, 'C': {linux.KEY_C, true, false},
	'd': {linux.KEY_D, false, false}, 'D': {linux.KEY_D, true, false},
	'e': {linux.KEY_E, false, false}, 'E': {linux.KEY_E, true, false},
	'f': {linux.KEY_F, false, false}, 'F': {linux.KEY_F, true, false},
	'g': {linux.KEY_G, false, false}, 'G': {linux.KEY_G, true, false},
	'h': {linux.KEY_H, false, false}, 'H': {linux.KEY_H, true, false},
	'i': {linux.KEY_I, false, false}, 'I': {linux.KEY_I, true, false},
	'j': {linux.KEY_J, false, false}, 'J': {linux.KEY_J, true, false},
	'k': {linux.KEY_K, false, false}, 'K': {linux.KEY_K, true, false},
	'l': {linux.KEY_L, false, false}, 'L': {linux.KEY_L, true, false},
	'm': {linux.KEY_M, false, false}, 'M': {linux.KEY_M, true, false},
	'n': {linux.KEY_N, false, false}, 'N': {linux.KEY_N, true, false},
	'o': {linux.KEY_O, false, false}, 'O': {linux.KEY_O, true, false},
	'p': {linux.KEY_P, false, false}, 'P': {linux.KEY_P, true, false},
	'q': {linux.KEY_Q, false, false}, 'Q': {linux.KEY_Q, true, false},
	'r': {linux.KEY_R, false, false}, 'R': {linux.KEY_R, true, false},
	's': {linux.KEY_S, false, false}, 'S': {linux.KEY_S, true, false},
	't': {linux.KEY_T, false, false}, 'T': {linux.KEY_T, true, false},
	'u': {linux.KEY_U, false, false}, 'U': {linux.KEY_U, true, false},
	'v': {linux.KEY_V, false, false}, 'V': {linux.KEY_V, true, false},
	'w': {linux.KEY_W, false, false}, 'W': {linux.KEY_W, true, false},
	'x': {linux.KEY_X, false, false}, 'X': {linux.KEY_X, true, false},
	'y': {linux.KEY_Y, false, false}, 'Y': {linux.KEY_Y, true, false},
	'z': {linux.KEY_Z, false, false}, 'Z': {linux.KEY_Z, true, false},

	// Digits and their shifted symbols
	'1': {linux.KEY_1, false, false}, '!': {linux.KEY_1, true, false},
	'2': {linux.KEY_2, false, false}, '@': {linux.KEY_2, true, false},
	'3': {linux.KEY_3, false, false}, '#': {linux.KEY_3, true, false},
	'4': {linux.KEY_4, false, false}, '$': {linux.KEY_4, true, false},
	'5': {linux.KEY_5, false, false}, '%': {linux.KEY_5, true, false},
	'6': {linux.KEY_6, false, false}, '^': {linux.KEY_6, true, false},
	'7': {linux.KEY_7, false, false}, '&': {linux.KEY_7, true, false},
	'8': {linux.KEY_8, false, false}, '*': {linux.KEY_8, true, false},
	'9': {linux.KEY_9, false, false}, '(': {linux.KEY_9, true, false},
	'0': {linux.KEY_0, false, false}, ')': {linux.KEY_0, true, false},

	// Special characters
	'-': {linux.KEY_MINUS, false, false}, '_': {linux.KEY_MINUS, true, false},
	'=': {linux.KEY_EQUAL, false, false}, '+': {linux.KEY_EQUAL, true, false},
	'[': {linux.KEY_LEFTBRACE, false, false}, '{': {linux.KEY_LEFTBRACE, true, false},
	']': {linux.KEY_RIGHTBRACE, false, false}, '}': {linux.KEY_RIGHTBRACE, true, false},
	'\\': {linux.KEY_BACKSLASH, false, false}, '|': {linux.KEY_BACKSLASH, true, false},
	';': {linux.KEY_SEMICOLON, false, false}, ':': {linux.KEY_SEMICOLON, true, false},
	'\'': {linux.KEY_APOSTROPHE, false, false}, '"': {linux.KEY_APOSTROPHE, true, false},
	',': {linux.KEY_COMMA, false, false}, '<': {linux.KEY_COMMA, true, false},
	'.': {linux.KEY_DOT, false, false}, '>': {linux.KEY_DOT, true, false},
	'/': {linux.KEY_SLASH, false, false}, '?': {linux.KEY_SLASH, true, false},
	'`': {linux.KEY_GRAVE, false, false}, '~': {linux.KEY_GRAVE, true, false},

	// Space and Enter
	' ':  {linux.KEY_SPACE, false, false},
	'\n': {linux.KEY_ENTER, false, false},
}

var azertyKeyMap = KeyMap{
	// Letters
	'a': {linux.KEY_Q, false, false}, 'A': {linux.KEY_Q, true, false},
	'b': {linux.KEY_B, false, false}, 'B': {linux.KEY_B, true, false},
	'c': {linux.KEY_C, false, false}, 'C': {linux.KEY_C, true, false},
	'd': {linux.KEY_D, false, false}, 'D': {linux.KEY_D, true, false},
	'e': {linux.KEY_E, false, false}, 'E': {linux.KEY_E, true, false},
	'f': {linux.KEY_F, false, false}, 'F': {linux.KEY_F, true, false},
	'g': {linux.KEY_G, false, false}, 'G': {linux.KEY_G, true, false},
	'h': {linux.KEY_H, false, false}, 'H': {linux.KEY_H, true, false},
	'i': {linux.KEY_I, false, false}, 'I': {linux.KEY_I, true, false},
	'j': {linux.KEY_J, false, false}, 'J': {linux.KEY_J, true, false},
	'k': {linux.KEY_K, false, false}, 'K': {linux.KEY_K, true, false},
	'l': {linux.KEY_L, false, false}, 'L': {linux.KEY_L, true, false},
	'm': {linux.KEY_SEMICOLON, false, false}, 'M': {linux.KEY_SEMICOLON, true, false},
	'n': {linux.KEY_N, false, false}, 'N': {linux.KEY_N, true, false},
	'o': {linux.KEY_O, false, false}, 'O': {linux.KEY_O, true, false},
	'p': {linux.KEY_P, false, false}, 'P': {linux.KEY_P, true, false},
	'q': {linux.KEY_A, false, false}, 'Q': {linux.KEY_A, true, false},
	'r': {linux.KEY_R, false, false}, 'R': {linux.KEY_R, true, false},
	's': {linux.KEY_S, false, false}, 'S': {linux.KEY_S, true, false},
	't': {linux.KEY_T, false, false}, 'T': {linux.KEY_T, true, false},
	'u': {linux.KEY_U, false, false}, 'U': {linux.KEY_U, true, false},
	'v': {linux.KEY_V, false, false}, 'V': {linux.KEY_V, true, false},
	'w': {linux.KEY_Z, false, false}, 'W': {linux.KEY_Z, true, false},
	'x': {linux.KEY_X, false, false}, 'X': {linux.KEY_X, true, false},
	'y': {linux.KEY_Y, false, false}, 'Y': {linux.KEY_Y, true, false},
	'z': {linux.KEY_W, false, false}, 'Z': {linux.KEY_W, true, false},

	// Digits and their shifted symbols
	'1': {linux.KEY_1, true, false}, '&': {linux.KEY_1, false, false},
	'2': {linux.KEY_2, true, false}, 'é': {linux.KEY_2, false, false},
	'3': {linux.KEY_3, true, false}, '"': {linux.KEY_3, false, false},
	'4': {linux.KEY_4, true, false}, '\'': {linux.KEY_4, false, false},
	'5': {linux.KEY_5, true, false}, '(': {linux.KEY_5, false, false},
	'6': {linux.KEY_6, true, false}, '-': {linux.KEY_6, false, false},
	'7': {linux.KEY_7, true, false}, 'è': {linux.KEY_7, false, false},
	'8': {linux.KEY_8, true, false}, '_': {linux.KEY_8, false, false},
	'9': {linux.KEY_9, true, false}, 'ç': {linux.KEY_9, false, false},
	'0': {linux.KEY_0, true, false}, 'à': {linux.KEY_0, false, false},
	')': {linux.KEY_MINUS, false, false}, '°': {linux.KEY_MINUS, true, false},

	// Special characters
	',': {linux.KEY_M, false, false}, '?': {linux.KEY_M, true, false},
	';': {linux.KEY_COMMA, false, false}, '.': {linux.KEY_COMMA, true, false},
	':': {linux.KEY_DOT, false, false}, '/': {linux.KEY_DOT, true, false},
	'!': {linux.KEY_SLASH, false, false}, '§': {linux.KEY_SLASH, true, false},
	'=': {linux.KEY_EQUAL, false, false}, '+': {linux.KEY_EQUAL, true, false},
	'[': {linux.KEY_5, false, true}, ']': {linux.KEY_MINUS, false, false},
	'{': {linux.KEY_4, false, true}, '}': {linux.KEY_EQUAL, false, true},
	'\\': {linux.KEY_8, false, true}, '|': {linux.KEY_6, false, true},
	'<': {linux.KEY_102ND, false, false}, '>': {linux.KEY_102ND, true, false},
	'`': {linux.KEY_7, false, true}, '~': {linux.KEY_2, false, true},
	'@': {linux.KEY_0, false, true}, '#': {linux.KEY_3, false, true},
	'^': {linux.KEY_LEFTBRACE, false, false}, '¨': {linux.KEY_LEFTBRACE, true, false},
	'*': {linux.KEY_BACKSLASH, false, false}, 'µ': {linux.KEY_BACKSLASH, true, false},
	'%': {linux.KEY_APOSTROPHE, false, false}, 'ù': {linux.KEY_BACKSLASH, true, false},

	// Space and Enter
	' ':  {linux.KEY_SPACE, false, false},
	'\n': {linux.KEY_ENTER, false, false},

	// Additional symbols
	'€': {linux.KEY_E, false, true},
	'$': {linux.KEY_RIGHTBRACE, false, false}, '£': {linux.KEY_RIGHTBRACE, true, false},
}
