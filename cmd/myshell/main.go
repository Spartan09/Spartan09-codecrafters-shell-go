package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Command interface {
	Execute(args []string) error
	Name() string
}

type Shell struct {
	commands map[string]Command
}

func NewShell() *Shell {
	s := &Shell{
		commands: make(map[string]Command),
	}
	s.registerBuiltins()
	return s
}

func (s *Shell) registerBuiltins() {
	s.commands["exit"] = &ExitCommand{}
	s.commands["echo"] = &EchoCommand{}
	s.commands["type"] = &TypeCommand{shell: s}
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

	cmd, exists := s.commands[parts[0]]
	if !exists {
		return fmt.Errorf("%s: command not found", parts[0])
	}

	return cmd.Execute(parts[1:])
}

type ExitCommand struct{}

func (c *ExitCommand) Name() string { return "exit" }
func (c *ExitCommand) Execute(args []string) error {
	os.Exit(0)
	return nil
}

type EchoCommand struct{}

func (c *EchoCommand) Name() string { return "echo" }
func (c *EchoCommand) Execute(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

type TypeCommand struct {
	shell *Shell
}

func (c *TypeCommand) Name() string { return "type" }
func (c *TypeCommand) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("type: incorrect number of arguments")
	}
	if _, exists := c.shell.commands[args[0]]; exists {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return nil
	}
	return c.searchInPath(args[0])
}

func (c *TypeCommand) searchInPath(cmd string) error {
	envPath := os.Getenv("PATH")
	paths := strings.Split(envPath, ":")
	for _, path := range paths {
		fp := filepath.Join(path, cmd)
		if _, err := os.Stat(fp); err == nil {
			fmt.Println(fp)
			return nil
		}
	}
	return fmt.Errorf("%s: not found", cmd)
}

func main() {
	shell := NewShell()
	if err := shell.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
