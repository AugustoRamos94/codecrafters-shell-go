package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

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
		default:
			fmt.Println(commands[0] + ": command not found")
		}
	}
}
