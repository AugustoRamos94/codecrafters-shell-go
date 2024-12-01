package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	commandsMapping := map[string]bool{
		"exit": true,
		"echo": true,
		"type": true,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		args := strings.Split(input, " ")
		cmd := args[0]
		args = args[1:]

		switch cmd {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(args[:], " "))
		case "type":
			_, ok := commandsMapping[args[0]]
			if !ok {
				fmt.Println(args[0] + ": not found")
				continue
			}

			fmt.Println(args[0] + " is a shell builtin")
		default:
			fmt.Println(cmd + ": command not found")
		}
	}
}
