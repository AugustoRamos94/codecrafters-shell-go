package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Printf("Error while reading command: %s", err)
		return
	}

	fmt.Printf("%s: command not found", s)
}
