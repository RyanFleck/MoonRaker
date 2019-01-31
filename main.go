package main

import (
	"fetch"
	"logger"
)

func main() {

	l := logger.Logger{
		Name: "Main",
	}

	l.Log("Copyright (c) 2019 Ryan Fleck")
	fetch.Test()

}
