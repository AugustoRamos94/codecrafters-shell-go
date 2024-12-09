package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var builtInCommands = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
	"pwd":  true,
	"cd":   true,
	// "cat":  true,
}

func handleCommand(cmd string, args []string) {
	switch cmd {
	case "exit":
		exit(args)
	case "echo":
		echo(args)
	case "type":
		_type(args)
	case "pwd":
		pwd()
	case "cd":
		cd(cmd, args)
	// case "cat":
	// 	cat(args)
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

func exit(args []string) {
	exitCode, _ := strconv.Atoi(args[0])
	os.Exit(exitCode)
}

func echo(args []string) {
	fmt.Println(strings.Join(args[:], " "))
}

func _type(args []string) {
	_, ok := builtInCommands[args[0]]
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
}

func pwd() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
}

func cd(cmd string, args []string) {
	dirname := args[0]

	if dirname == "~" {
		dirname, _ = os.UserHomeDir()
	}

	if err := os.Chdir(dirname); err != nil {
		fmt.Printf("%s: %s: No such file or directory\n", cmd, dirname)
	}
}

// func cat(args []string) {
// 	for _, fileName := range args {
// 		file, _ := os.Open(fileName)
// 		fmt.Fprint(os.Stdout, "filename: ", fileName)
// 		defer file.Close()

// 		r := bufio.NewReader(file)

// 		for {
// 			line, _, _ := r.ReadLine()
// 			if len(line) > 0 {
// 				fmt.Fprint(os.Stdout, "%q", line)
// 			}
// 		}
// 	}
// }
