package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var commandsMapping = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
	"pwd":  true,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		args := strings.Split(input, " ")
		cmd := args[0]
		args = args[1:]

		handleCommand(cmd, args)
	}
}

func handleCommand(cmd string, args []string) {
	switch cmd {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(args[:], " "))
	case "type":
		_, ok := commandsMapping[args[0]]
		if ok {
			fmt.Println(args[0] + " is a shell builtin")
			return
		}

		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			fp := filepath.Join(path, args[0])
			if _, err := os.Stat(fp); err == nil {
				fmt.Println(fp)
				return
			}
		}

		fmt.Println(args[0] + ": not found")
	case "pwd":
		pwd, _ := os.Getwd()
		fmt.Println(pwd)
	default:
		externalCommand := exec.Command(cmd, args...)
		externalCommand.Stderr = os.Stderr
		externalCommand.Stdout = os.Stdout
		err := externalCommand.Run()
		if err != nil {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
