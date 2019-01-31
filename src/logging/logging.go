package logging

import (
	"fmt"
)

func Log(s string) {
	fmt.Println("-> %s\n", s)
}
