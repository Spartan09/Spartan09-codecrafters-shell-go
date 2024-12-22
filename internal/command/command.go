package command

import "github.com/codecrafters-io/shell-starter-go/internal/parser"

type Command interface {
	Name() string
	Execute(args []string, redirect *parser.RedirectInfo) error
}

type BuiltinChecker interface {
	IsBuiltin(name string) bool
}
