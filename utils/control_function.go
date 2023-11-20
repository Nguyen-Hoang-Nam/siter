package utils

var controlCharacter = map[rune]string{
	0x00: "NUL",
	0x01: "SOH",
	0x02: "STX",
	0x03: "ETX",
	0x04: "EOT",
	0x05: "ENQ",
	0x06: "ACK",
	0x07: "BEL",
	0x08: "BS",
	0x09: "HT",
	0x0a: "LF",
	0x0b: "VT",
	0x0c: "FF",
	0x0d: "CR",
	0x0e: "SO",
	0x0f: "SI",
	0x10: "DLE",
	0x11: "DC1",
	0x12: "DC2",
	0x13: "DC3",
	0x14: "DC4",
	0x15: "NAK",
	0x16: "SYN",
	0x17: "ETB",
	0x18: "CAN",
	0x19: "EM",
	0x1a: "SUB",
	0x1c: "FS",
	0x1d: "GS",
	0x1e: "RS",
	0x1f: "US",
	0x84: "IND",
	0x85: "NEL",
	0x88: "HTS",
	0x8d: "RI",
	0x8e: "SS2",
	0x8f: "SS3",
	0x96: "SPA",
	0x97: "EDA",
	0x98: "SOS",
	0x9a: "DECID",
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
	'D': {
		-1: {'D': "IND"},
	},
	'E': {
		-1: {'E': "NEL"},
	},
	'H': {
		-1: {'H': "HTS"},
	},
	'M': {
		-1: {'M': "RI"},
	},
	'N': {
		-1: {'N': "SS2"},
	},
	'O': {
		-1: {'O': "SS3"},
	},
	'V': {
		-1: {'V': "SPA"},
	},
	'W': {
		-1: {'W': "EPA"},
	},
	'X': {
		-1: {'X': "SOS"},
	},
	'Z': {
		-1: {'Z': "DECID"},
	},
	'c': {
		-1: {'c': "RIS"},
	},
	'=': {
		-1: {'=': "DECKPAM"},
	},
	'>': {
		-1: {'>': "DECPNM"},
	},
	'\\': {
		-1: {'\\': "ST"},
	},
	'6': {
		-1: {'6': "DECBI"},
	},
	'7': {
		-1: {'7': "DECSC"},
	},
	'8': {
		-1: {'8': "DECRC"},
	},
	'9': {
		-1: {'9': "DECFI"},
	},
	'#': {
		'3': {'3': "DECDHLT"},
		'4': {'4': "DECDHLB"},
		'5': {'5': "DECSWL"},
		'6': {'6': "DECDWL"},
		'8': {'8': "DECALN"},
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

func ControlCharacter(r rune) (name string, isEnd bool) {
	if isControlSequence(r) {
		return "", false
	}

	if val, ok := controlCharacter[r]; ok {
		return val, true
	}

	return "", true
}

func ControlSequence(rs []rune) (name string, isEnd bool) {
	rs = convertC1ControlSequence(rs)

	if val, ok := controlSequences[rs[1]]; ok {
		if len(rs) == 2 {
			if val1, ok := val[-1]; ok {
				if val2, ok := val1[rs[len(rs)-1:][0]]; ok {
					return val2, true
				}
			}
		} else {
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
	}

	return "", false
}

func isControlSequence(r rune) bool {
	return r == 0x1b || isC1ControlSequence(r)
}

func isC1ControlSequence(r rune) bool {
	return r == 0x9b || r == 0x9d || r == 0x90 || r == 0x9e || r == 0x9f
}

func convertC1ControlSequence(rs []rune) []rune {
	if isC1ControlSequence(rs[0]) {
		return append(controlCharacterToEscapeSequence[rs[0]], rs[1:]...)
	}

	return rs
}
