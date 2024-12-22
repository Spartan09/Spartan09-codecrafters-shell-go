// internal/builtin/echo.go
package builtin

import (
	"fmt"
	"os"
	"strings"
)

type EchoCommand struct{}

func (c *EchoCommand) Name() string { return "echo" }
func (c *EchoCommand) Execute(args []string, redirectFile string) error {
	output := strings.Join(args, " ")

	if redirectFile != "" {
		f, err := os.OpenFile(redirectFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Fprintln(f, output)
	} else {
		fmt.Println(output)
	}
	return nil
}
