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
		input = strings.Trim(input, "\n")

		args := parseCommand(input)
		cmd := args[0]
		args = args[1:]

		handleCommand(cmd, args)
	}
}
