package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Printf("Error while reading command: %s", err)
			return
		}

		fmt.Println(command[:len(command)-1] + ": command not found")
	}
}
