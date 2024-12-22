package command

type Command interface {
	Execute(args []string) error
	Name() string
}

type BuiltinChecker interface {
	IsBuiltin(name string) bool
}
