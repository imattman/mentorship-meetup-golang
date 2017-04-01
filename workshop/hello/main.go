package main

import (
	"fmt"
	"os"
)

func main() {
	who := "World"
	if len(os.Args) > 1 {
		who = os.Args[1]
	}
	fmt.Printf("Hello, %s!\n", who)
}
