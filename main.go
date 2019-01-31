package main

import (
	"fmt"
	"logging"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")
	logging.Log("Test")
	fmt.Println("The time is", time.Now())
}
