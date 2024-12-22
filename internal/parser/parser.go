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

func (p *ShellParser) Parse(input string) ([]string, string) {
	var redirectFile string
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
			case '\\':
				if i+1 < len(input) {
					p.current.WriteByte(input[i+1])
					i++
				}

			case '>':
				// Add current argument (without the '1' if it exists)
				str := p.current.String()
				if i > 0 && input[i-1] == '1' && len(str) > 0 {
					str = str[:len(str)-1] // Remove the '1'
				}
				p.current.Reset()
				if str != "" {
					p.args = append(p.args, str)
				}

				// Skip spaces after >
				i++
				for i < len(input) && (input[i] == ' ' || input[i] == '\t') {
					i++
				}

				// Collect filename
				start := i
				for i < len(input) && input[i] != ' ' && input[i] != '\t' {
					i++
				}
				redirectFile = input[start:i]
				i-- // Back up one to handle next character properly
				continue

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
			case '\\':
				if i+1 < len(input) {
					next := input[i+1]
					if next == '"' || next == '\\' {
						p.current.WriteByte(next)
						i++
					} else {
						p.current.WriteByte('\\')
					}
				}
			default:
				p.current.WriteByte(ch)
			}
		}

	}
	p.addArgument()
	return p.args, redirectFile
}

func (p *ShellParser) addArgument() {
	if p.current.Len() > 0 {
		p.args = append(p.args, p.current.String())
		p.current.Reset()
	}
}
