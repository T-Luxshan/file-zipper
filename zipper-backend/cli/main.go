package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"zipper-backend/pkg/compressor"
)

func main() {
	fmt.Println("Zipper running....")

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run cli/main.go [compress|decompress] <file>")
		return
	}
	fmt.Println("Inside the application")

	command := os.Args[1]
	path := os.Args[2]

	switch strings.ToLower(command) {
	case "compress":
		info, err := os.Stat(path)
		if err != nil {
			log.Fatalf("specified path does not exist")
		}
		if info.IsDir() {
			msg, err := compressor.CmpressFolder(path)
			if err != nil {
				log.Fatal("Compressing failed: ", err)
			}
			fmt.Println(msg)
		} else {
			msg, err := compressor.CompressFile(path)
			if err != nil {
				log.Fatal("Compressing failed: ", err)
			}
			fmt.Println(msg)
		}
		return
	case "decompress":
		//TODO: DO the logic for decompress
		info, err := os.Stat(path)
		if err != nil {
			log.Fatalf("specified path does not exist")
		}
		if info.IsDir() {
			msg, err := compressor.CmpressFolder(path)
			if err != nil {
				log.Fatal("Compressing failed: ", err)
			}
			fmt.Println(msg)
		} else {
			msg, err := compressor.CompressFile(path)
			if err != nil {
				log.Fatal("Compressing failed: ", err)
			}
			fmt.Println(msg)
		}
		return
	default:
		log.Fatal("Unknown command: " + command)
	}

}
