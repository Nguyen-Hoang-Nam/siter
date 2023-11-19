package utils

var controlFunctionOneCharacter = map[string]rune{
	"NUL":   0x00,
	"BEL":   0x07,
	"BS":    0x08,
	"HT":    0x09,
	"LF":    0x0a,
	"VT":    0x0b,
	"FF":    0x0c,
	"CR":    0x0d,
	"IND":   0x84,
	"NEL":   0x85,
	"HTS":   0x88,
	"RI":    0x8d,
	"SS2":   0x8e,
	"SS3":   0x8f,
	"SPA":   0x96,
	"EDA":   0x97,
	"SOS":   0x98,
	"DECID": 0x9a,
}

var controlFunctionEsc7Bit = map[string]rune{
	"DECBI":   '6',
	"DECSC":   '7',
	"DECRC":   '8',
	"DECFI":   '9',
	"IND":     'D',
	"NEL":     'E',
	"HTS":     'H',
	"RI":      'M',
	"SS2":     'N',
	"SS3":     'O',
	"SPA":     'V',
	"EPA":     'W',
	"SOS":     'X',
	"DECID":   'Z',
	"RIS":     'c',
	"DECKPAM": '=',
}

var controlCharacterToEscapeSequence = map[rune][]rune{
	0x9b: {0x1b, '['},
	0x9d: {0x1b, ']'},
	0x90: {0x1b, 'P'},
	0x9e: {0x1b, '^'},
	0x9f: {0x1b, '_'},
}

var controlSequences = map[rune]map[rune]map[int32]string{
	0x20: {
		'F': {'F': "S7C1T"},
		'G': {'G': "S8C1T"},
	},
	'#': {
		'3': {'3': "DECDHLT"},
		'4': {'4': "DECDHLB"},
		'5': {'5': "DECSWL"},
		'6': {'6': "DECDWL"},
	},
	'[': {
		-1: {
			'@': "ICH",
			'A': "CUU",
			'B': "CUD",
			'C': "CUF",
			'D': "CUB",
			'E': "CNL",
			'F': "CPL",
			'G': "CHA",
			'H': "CUP",
			'J': "CHT",
			'K': "EL",
			'L': "IL",
			'M': "DL",
			'P': "DCH",
			'S': "SU",
			'T': "SD",
			'X': "ECH",
			'Z': "CBT",
			'`': "HPA",
			'b': "REP",
			'c': "DA1",
			'd': "VPA",
			'f': "HVP",
			'g': "TBC",
			'h': "SM",
			'i': "MC",
			'l': "RM",
			'm': "SGR",
			'n': "DSR",
			'p': "DECSCL",
			'q': "DECSCA",
			's': "ANSI.SYS",
			'u': "ANSI.SYS",
			'w': "DECEFR",
		},
		'?': {
			'J': "DECSED",
			'K': "DECSEL",
			'h': "DECSET",
			'i': "MC",
			'l': "DECRST",
			'n': "DSR",
		},
		'>': {
			'c': "DA2",
		},
		'=': {
			'c': "DA3",
		},
		'!': {
			'p': "DECSTR",
		},
	},
	']': {
		-1: {
			0x07: "OSC",
			0x9c: "OSC",
		},
	},
	'P': {
		-1: {
			0x9c: "DCS",
		},
	},
	'_': {
		-1: {
			0x9c: "APC",
		},
	},
	'^': {
		-1: {
			0x9c: "PM",
		},
	},
}

func IsControlCharacter(r rune) bool {
	return (r < 0x20) || (r > 0x7e && r < 0xa0)
}

func ControlFunctionOneCharacter(r rune) (name string, isEnd bool) {
	if isEscapeSequence(r) {
		return "", false
	}

	for k, v := range controlFunctionOneCharacter {
		if v == r {
			return k, true
		}
	}

	return "", true
}

func ControlFunctionEsc7Bit(rs []rune) (name string, isEnd bool) {
	if is8BitControlCharacter(rs[0]) {
		return "", false
	}

	if rs[0] != 0x1b {
		return "", true
	}

	for k, v := range controlFunctionEsc7Bit {
		if v == rs[1] {
			return k, true
		}
	}

	return "", false
}

func ControlFunctionEscSequence(rs []rune) (name string, isEnd bool) {
	rs = convert8BitControlCharacterTo7BitEscapeSequence(rs)

	if val, ok := controlSequences[rs[1]]; ok {
		if val1, ok := val[rs[2]]; ok {
			if val2, ok := val1[rs[len(rs)-1:][0]]; ok {
				return val2, true
			}

		}

		if val1, ok := val[-1]; ok {
			if val2, ok := val1[rs[len(rs)-1:][0]]; ok {
				return val2, true
			}
		}

	}

	return "", false
}

func isEscapeSequence(r rune) bool {
	return r == 0x1b || is8BitControlCharacter(r)
}

func is8BitControlCharacter(r rune) bool {
	return r == 0x9b || r == 0x9d || r == 0x90 || r == 0x9e || r == 0x9f
}

func convert8BitControlCharacterTo7BitEscapeSequence(rs []rune) []rune {
	if is8BitControlCharacter(rs[0]) {
		return append(controlCharacterToEscapeSequence[rs[0]], rs[1:]...)
	}

	return rs
}
