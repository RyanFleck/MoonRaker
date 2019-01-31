package main

import (
	"logger"
)

func main() {

	logger.Log("Test")
	logger.Err("Aaaaa!")
	logger.Log("This should not print.")

}
