package main

import (
	"strings"
	"unicode"
)

func parseCommand(command string) []string {
	var args []string
	var current strings.Builder
	inSingleQuotes := false
	inDoubleQuotes := false
	escaped := false
	for _, char := range command {
		switch {
		case char == '\'':
			if escaped && inDoubleQuotes {
				current.WriteRune('\\')
			}
			if inDoubleQuotes || escaped {
				current.WriteRune(char)
			} else {
				inSingleQuotes = !inSingleQuotes
			}
			escaped = false
		case char == '"':
			if inSingleQuotes || escaped {
				current.WriteRune(char)
			} else {
				inDoubleQuotes = !inDoubleQuotes
			}
			escaped = false
		case char == '\\':
			if inSingleQuotes || escaped {
				current.WriteRune(char)
				escaped = false
			} else {
				escaped = true
			}
		case unicode.IsSpace(char):
			if escaped && (inDoubleQuotes || inSingleQuotes) {
				current.WriteRune('\\')
			}
			if inSingleQuotes || inDoubleQuotes || escaped {
				current.WriteRune(char)
			} else if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
			escaped = false
		default:
			if escaped && inDoubleQuotes {
				current.WriteRune('\\')
			}
			current.WriteRune(char)
			escaped = false
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}
