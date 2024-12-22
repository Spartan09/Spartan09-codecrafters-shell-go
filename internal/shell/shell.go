package shell

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/builtin"
	"github.com/codecrafters-io/shell-starter-go/internal/command"
	"github.com/codecrafters-io/shell-starter-go/internal/external"
	"os"
	"strings"
)

type Shell struct {
	Commands map[string]command.Command
}

func NewShell() *Shell {
	s := &Shell{
		Commands: make(map[string]command.Command),
	}
	s.registerBuiltins()
	return s
}

func (s *Shell) IsBuiltin(name string) bool {
	_, exists := s.Commands[name]
	return exists
}

func (s *Shell) registerBuiltins() {
	s.Commands["exit"] = &builtin.ExitCommand{}
	s.Commands["echo"] = &builtin.EchoCommand{}
	s.Commands["type"] = &builtin.TypeCommand{Checker: s}
}

func (s *Shell) Run() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("reading input: %w", err)
		}

		if err := s.Execute(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func (s *Shell) Execute(input string) error {
	parts := strings.Fields(strings.TrimSpace(input))
	if len(parts) == 0 {
		return nil
	}

	cmd, exists := s.Commands[parts[0]]
	if exists {
		return cmd.Execute(parts[1:])
	}

	if err := external.Execute(parts); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
