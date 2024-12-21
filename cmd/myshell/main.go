package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	_, exists := commands[args[0]]
	if exists {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return
	}
	envPath := os.Getenv("PATH")
	paths := strings.Split(envPath, ":")
	//fmt.Printf("CMD: %s\n", cmd)
	//fmt.Printf("PATH: %s\n", envPath)
	for _, path := range paths {
		fp := filepath.Join(path, args[0])
		if _, err := os.Stat(fp); err == nil {
			fmt.Println(fp)
			return
		}
	}
	fmt.Printf("%s: not found\n", args[0])
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
