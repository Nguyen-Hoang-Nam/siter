package controlfunction

var controlCharacterToEscapeSequence = map[rune][]rune{
	rune(CSI): {rune(ESC), '['},
	rune(OSC): {rune(ESC), ']'},
	rune(DCS): {rune(ESC), 'P'},
	rune(PM):  {rune(ESC), '^'},
	rune(APC): {rune(ESC), '_'},
}

var controlSequences = map[rune]map[rune]map[rune]FunctionName{
	// nF Escape sequences
	' ': {
		'A':  {'A': ISO2022},
		'B':  {'B': ISO2022},
		'C':  {'C': ISO2022},
		'D':  {'D': ISO2022},
		'E':  {'E': ISO2022},
		'F':  {'F': S7C1T},
		'G':  {'G': S8C1T},
		'H':  {'H': ISO2022},
		'I':  {'I': ISO2022},
		'J':  {'J': ISO2022},
		'K':  {'K': ISO2022},
		'L':  {'L': ISO4873L1},
		'M':  {'M': ISO4873L2},
		'N':  {'N': ISO4873L3},
		'P':  {'P': ISO2022},
		'R':  {'R': ISO2022},
		'S':  {'S': ISO2022},
		'T':  {'T': ISO2022},
		'U':  {'U': ISO2022},
		'V':  {'V': ISO2022},
		'W':  {'W': ISO2022},
		'Z':  {'Z': ISO2022},
		'[':  {'[': ISO2022},
		'\\': {'\\': ISO2022},
	},
	'#': {
		'3': {'3': DECDHLT},
		'4': {'4': DECDHLB},
		'5': {'5': DECSWL},
		'6': {'6': DECDWL},
		'8': {'8': DECALN},
	},
	'%': {
		-1: {'G': UTF8},
		'/': {
			'I': UTF8,
			'L': UTF16,
			'F': UTF32,
		},
	},

	// Fp Escape sequences
	'6': {-1: {'6': DECBI}},
	'7': {-1: {'7': DECSC}},
	'8': {-1: {'8': DECRC}},
	'9': {-1: {'9': DECFI}},
	'=': {-1: {'=': DECKPAM}},
	'>': {-1: {'>': DECPNM}},

	// Esc + 7 bit
	'@': {-1: {'@': PAD}},
	'A': {-1: {'A': HOP}},
	'B': {-1: {'B': BPH}},
	'C': {-1: {'C': NBH}},
	'D': {-1: {'D': IND}},
	'E': {-1: {'E': NEL}},
	'F': {-1: {'F': SSA}},
	'G': {-1: {'G': ESA}},
	'H': {-1: {'H': HTS}},
	'I': {-1: {'I': HTJ}},
	'J': {-1: {'H': VTS}},
	'K': {-1: {'H': PLD}},
	'L': {-1: {'H': PLU}},
	'M': {-1: {'M': RI}},
	'N': {-1: {'N': SS2}},
	'O': {
		-1:  {'O': SS3},
		'P': {'P': F1},
		'Q': {'Q': F2},
		'R': {'R': F3},
		'S': {'S': F4},
	},
	'Q':  {-1: {'O': PU1}},
	'R':  {-1: {'R': PU2}},
	'S':  {-1: {'O': STS}},
	'T':  {-1: {'O': CCH}},
	'U':  {-1: {'O': MW}},
	'V':  {-1: {'V': SPA}},
	'W':  {-1: {'W': EPA}},
	'X':  {-1: {'X': SOS}},
	'Z':  {-1: {'Z': DECID}},
	'\\': {-1: {'\\': ST}},

	// FS Escape sequences
	0x60: {-1: {0x60: DMI}},
	0x61: {-1: {0x61: INT}},
	0x62: {-1: {0x62: EMI}},
	0x63: {-1: {0x63: RIS}},
	0x64: {-1: {0x64: CMD}},
	0x6e: {-1: {0x6e: LS2}},
	0x6f: {-1: {0x6f: LS3}},
	0x7c: {-1: {0x7c: LS3R}},
	0x7d: {-1: {0x7d: LS2R}},
	0x7e: {-1: {0x7e: LS1R}},

	// Esc + 8 bit
	0x82: {-1: {0x82: BPH}},
	0x83: {-1: {0x83: NBH}},
	0x84: {-1: {0x84: IND}},
	0x85: {-1: {0x85: NEL}},
	0x86: {-1: {0x86: SSA}},
	0x87: {-1: {0x87: ESA}},
	0x88: {-1: {0x88: HTS}},
	0x89: {-1: {0x89: HTJ}},
	0x8a: {-1: {0x8a: VTS}},
	0x8b: {-1: {0x8b: PLD}},
	0x8c: {-1: {0x8c: PLU}},
	0x8d: {-1: {0x8d: RI}},
	0x8e: {-1: {0x8e: SS2}},
	0x8f: {-1: {0x8f: SS3}},
	0x91: {-1: {0x91: PU1}},
	0x92: {-1: {0x92: PU2}},
	0x93: {-1: {0x93: STS}},
	0x94: {-1: {0x94: CCH}},
	0x95: {-1: {0x95: MW}},
	0x96: {-1: {0x96: SPA}},
	0x97: {-1: {0x97: EPA}},
	0x98: {-1: {0x98: SOS}},
	0x99: {-1: {0x99: DECID}},
	0x9c: {-1: {0x9c: ST}},

	// CSI
	'[': {
		-1: {
			'@':  ICH,
			'A':  CUU,
			'B':  CUD,
			'C':  CUF,
			'D':  CUB,
			'E':  CNL,
			'F':  CPL,
			'G':  CHA,
			'H':  CUP,
			'J':  CHT,
			'K':  EL,
			'L':  IL,
			'M':  DL,
			'N':  EF,
			'P':  DCH,
			'Q':  SSE,
			'R':  CPR,
			'S':  SU,
			'T':  SD,
			'U':  NP,
			'V':  PP,
			'W':  CTC,
			'X':  ECH,
			'Y':  CVT,
			'Z':  CBT,
			'[':  SRS,
			'\\': PTX,
			']':  SDS,
			'^':  SIMD,
			'`':  HPA,
			'a':  HPR,
			'b':  REP,
			'c':  DA,
			'd':  VPA,
			'e':  VPR,
			'f':  HVP,
			'g':  TBC,
			'h':  SM,
			'i':  MC,
			'j':  HPB,
			'k':  VPB,
			'l':  RM,
			'm':  SGR,
			'n':  DSR,
			'o':  DAQ,
			'p':  DECSCL,
			'q':  DECSCA,
			's':  SCOSC,
			'u':  SCORC,
			'w':  DECEFR,
			'~':  FUNCTIONKEY,
		},
		'?': {
			'J': DECSED,
			'K': DECSEL,
			'h': DECSET,
			'i': MC,
			'l': DECRST,
			'n': DSR,
		},
		'>': {'c': DA2},
		'=': {'c': DA3},
		'!': {'p': DECSTR},
	},

	// OSC
	']': {
		-1: {
			rune(BEL): OSC,
			rune(ST):  OSC,
		},
	},

	// DCS
	'P': {
		-1:  {rune(ST): DCS},
		'!': {rune(ST): DECRPTUI},
		'"': {rune(ST): DECCKD},
	},

	// APC
	'_': {-1: {rune(ST): APC}},

	// PM
	'^': {-1: {rune(ST): PM}},
}

func ControlSequence(rs []rune) (name FunctionName, isEnd bool) {
	rs = convertC1ControlSequence(rs)
	if val, ok := controlSequences[rs[1]]; ok {
		if len(rs) == 2 {
			if rs[1] != '[' && rs[1] != ']' && rs[1] != 'P' && rs[1] != '_' && rs[1] != '^' {
				if val1, ok := val[-1]; ok {
					if val2, ok := val1[rs[len(rs)-1:][0]]; ok {
						return val2, true
					}
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

	return -1, false
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
