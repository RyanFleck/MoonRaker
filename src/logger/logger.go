package logger

import (
	"fmt"
	"os"
)

// Log prints a formatted log string to the console.
func Log(s string) {
	fmt.Fprintf(os.Stdout, "-> %v\n", s)
}

// Err prints a formatted error message to the console.
func Err(s string) {
	fmt.Fprintf(os.Stderr, "!! %s\n", s)
	os.Exit(1)
}
