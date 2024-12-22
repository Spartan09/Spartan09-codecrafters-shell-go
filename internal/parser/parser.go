package parser

import "strings"

type RedirectInfo struct {
	StdoutFile string
	StderrFile string
	IsAppend   bool
}

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

func (p *ShellParser) Parse(input string) ([]string, *RedirectInfo) {
	redirect := &RedirectInfo{}

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
				isStderr := false
				// Check if it's '2>' or '1>' or just '>'
				if i > 0 && input[i-1] == '2' {
					isStderr = true
					// Remove the '2' from current argument
					str := p.current.String()
					if len(str) > 0 {
						str = str[:len(str)-1] // Remove the '2'
						p.current.Reset()
						if str != "" {
							p.args = append(p.args, str)
						}
					}
				} else if i > 0 && input[i-1] == '1' {
					// Remove the '1' from current argument
					str := p.current.String()
					if len(str) > 0 {
						str = str[:len(str)-1] // Remove the '1'
						p.current.Reset()
						if str != "" {
							p.args = append(p.args, str)
						}
					}
				} else {
					p.addArgument()
				}

				// Handle >> case
				if i+1 < len(input) && input[i+1] == '>' {
					redirect.IsAppend = true
					i++
				}

				// Skip spaces after > or >>
				i++
				for i < len(input) && (input[i] == ' ' || input[i] == '\t') {
					i++
				}

				// Collect filename
				start := i
				for i < len(input) && input[i] != ' ' && input[i] != '\t' {
					i++
				}
				if isStderr {
					redirect.StderrFile = input[start:i]
				} else {
					redirect.StdoutFile = input[start:i]
				}
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
	return p.args, redirect
}

func (p *ShellParser) addArgument() {
	if p.current.Len() > 0 {
		p.args = append(p.args, p.current.String())
		p.current.Reset()
	}
}
