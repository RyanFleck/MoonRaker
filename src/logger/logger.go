package logger

import (
	"fmt"
	"os"
)

// Log prints a formatted log string to the console.
func Log(msg string) {
	symlog("---", "general", msg)
}

// Err prints a formatted error message to the console.
func Err(msg string) {
	symerr("!!!", "general", msg)
}

// Logger should be instantiated for every file. Exposes logging methods.
type Logger struct {
	Name string
}

// Log prints a formatted log string to the console.
func (l Logger) Log(msg string) {
	symlog("---", l.Name, msg)
}

// EnterMethod prints a Log with formatting so a method entry is implied.
func (l Logger) EnterMethod(msg string) {
	symlog(">-<", l.Name, msg)
}

// MethodLog prints a log implying it is coming from within a method.
func (l Logger) MethodLog(msg string) {
	symlog(">->", l.Name, msg)
}

// ExitMethod
func (l Logger) ExitMethod(msg string) {
	symlog("<->", l.Name, msg)
}

// Err prints a formatted error message to the console.
func (l Logger) Err(msg string) {
	symerr("!!!", l.Name, msg)
}

/*
 * SymLog functions:
 *   prints a message with a symbol.
 */

func symlog(sym string, pkg string, msg string) {
	fmt.Fprintf(os.Stdout, "%s %s ->\t%s\n", sym, pkg, msg)
}

func symerr(sym string, pkg string, err string) {
	fmt.Fprintf(os.Stderr, "%s %s ->\t%s\n", sym, pkg, err)
	os.Exit(1)
}
