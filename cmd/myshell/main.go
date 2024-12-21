package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CommandHandler func(args []string)

var commands map[string]CommandHandler

func exitHandler(args []string) {
	os.Exit(0)
}

func echoHandler(args []string) {
	echo := strings.Join(args, " ")
	fmt.Printf("%s\n", echo)
}

func typeHandler(args []string) {
	if len(args) != 1 {
		return
	}
	if _, exists := commands[args[0]]; exists {
		fmt.Printf("%s is a shell builtin\n", args[0])
	} else {
		fmt.Printf("%s: not found\n", args[0])
	}
}

func main() {
	commands = map[string]CommandHandler{
		"exit": exitHandler,
		"echo": echoHandler,
		"type": typeHandler,
	}
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		command = strings.TrimSuffix(command, "\n")
		parts := strings.Fields(command)

		if len(parts) == 0 {
			continue
		}

		if handler, exists := commands[parts[0]]; exists {
			handler(parts[1:]) // Pass remaining args to handler
		} else {
			fmt.Printf("%s: command not found\n", parts[0])
		}
	}
}
