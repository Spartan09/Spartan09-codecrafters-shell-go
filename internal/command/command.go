package command

type Command interface {
	Name() string
	Execute(args []string, redirectFile string) error
}

type BuiltinChecker interface {
	IsBuiltin(name string) bool
}
