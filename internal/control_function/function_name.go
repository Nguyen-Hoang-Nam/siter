package controlfunction

type FunctionName int

const (
	NUL FunctionName = iota
	SOH
	STX
	ETX
	EOT
	ENQ
	ACK
	BEL
	BS
	HT
	LF
	VT
	FF
	CR
	SO
	SI
	DLE
	DC1
	DC2
	DC3
	DC4
	NAK
	SYN
	ETB
	CAN
	EM
	SUB
	ESC
	FS
	GS
	RS
	US
	DEL FunctionName = 0x7f

	PAD FunctionName = iota + 0x5f
	HOP
	BPH
	NBH
	IND
	NEL
	SSA
	ESA
	HTS
	HTJ
	VTS
	PLD
	PLU
	RI
	SS2
	SS3
	DCS
	PU1
	PU2
	STS
	CCH
	MW
	SPA
	EPA
	SOS
	DECID
	SCI
	CSI
	ST
	OSC
	PM
	APC

	ICH
	CUU
	CUD
	CUF
	CUB
	CNL
	CPL
	CHA
	CUP
	CHT
	ED
	EL
	IL
	DL
	EF
	EA
	DCH
	SSE
	CPR
	SU
	SD
	NP
	PP
	CTC
	ECH
	CVT
	CBT
	SRS
	PTX
	SDS
	SIMD
	_
	HPA
	HPR
	REP
	DA
	VPA
	VPR
	HVP
	TBC
	SM
	MC
	HPB
	VPB
	RM
	SGR
	DSR
	DAQ
	DECSCL
	DECSCA
	SCOSC
	SCORC
	DECEFR

	DECSED
	DECSEL
	DECSET
	DECRST

	DA2
	DA3
	DECSTR

	DMI
	INT
	EMI
	RIS
	CMD
	LS0
	LS1
	LS2
	LS3
	LS3R
	LS2R
	LS1R

	// ISO 2022
	S7C1T
	S8C1T
	ISO4873L1
	ISO4873L2
	ISO4873L3
	UTF8
	UTF16
	UTF32
	ISO2022

	DECDHLT
	DECDHLB
	DECSWL
	DECDWL
	DECALN

	DECBI
	DECSC
	DECRC
	DECFI
	DECKPAM
	DECPNM

	DECRPTUI
	DECCKD
)
