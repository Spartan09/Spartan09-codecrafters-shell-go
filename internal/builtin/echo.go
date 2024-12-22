package builtin

import (
	"fmt"
	"strings"
)

type EchoCommand struct{}

func (c *EchoCommand) Name() string { return "echo" }
func (c *EchoCommand) Execute(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
