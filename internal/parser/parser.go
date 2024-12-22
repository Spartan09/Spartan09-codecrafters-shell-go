package parser

import "strings"

type ShellParser struct {
	state   state
	current strings.Builder
	args    []string
}

func NewParser() *ShellParser {
	return &ShellParser{
		state: stateNormal,
		args:  make([]string, 0),
	}
}

func (p *ShellParser) Parse(input string) []string {
	for i := 0; i < len(input); i++ {
		ch := input[i]
		switch p.state {
		case stateNormal:
			switch ch {
			case '\'':
				p.state = stateSingleQuote
			case '"':
				p.state = stateDoubleQuote
			case ' ', '\t':
				p.addArgument()
			default:
				p.current.WriteByte(ch)
			}

		case stateSingleQuote:
			switch ch {
			case '\'':
				p.state = stateNormal
			default:
				p.current.WriteByte(ch)
			}
		case stateDoubleQuote:
			switch ch {
			case '"':
				p.state = stateNormal
			default:
				p.current.WriteByte(ch)
			}
		}

	}
	p.addArgument()
	return p.args
}

func (p *ShellParser) addArgument() {
	if p.current.Len() > 0 {
		p.args = append(p.args, p.current.String())
		p.current.Reset()
	}
}
