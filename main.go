package main

import (
	"fmt"
	"os"
)

func main() {
	// No need for named parameters
	// such as '--port or -p'
	port := os.Args[1]

	fmt.Println(port)
}
