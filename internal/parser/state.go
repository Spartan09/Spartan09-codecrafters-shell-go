package parser

type state int

const (
	stateNormal state = iota
	stateSingleQuote
)
