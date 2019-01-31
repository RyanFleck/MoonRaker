package main

import (
	"logger"
)

func main() {

	l := logger.Logger{
		Name: "Main",
	}
	l.Log("Test")
	l.Err("Aaaaa!")
	l.Log("This should not print.")

}
