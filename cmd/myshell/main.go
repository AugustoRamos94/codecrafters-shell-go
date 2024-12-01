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

		commands := strings.Split(input, " ")

		switch commands[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(commands[1:], " "))
		case "type":
			_, ok := commandsMapping[commands[1]]
			if !ok {
				fmt.Println(commands[1] + ": not found")
				continue
			}

			fmt.Println(commands[1] + " is a shell builtin")
		default:
			fmt.Println(commands[0] + ": command not found")
		}
	}
}
