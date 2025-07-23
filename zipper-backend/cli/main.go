package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Application running")

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run cli/main.go [compress|decompress] <file>")
		return
	}
	fmt.Println("Inside the application")

	command := os.Args[1]
	path := os.Args[2]

	fmt.Println(command)
	fmt.Println(path)
}
