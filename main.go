package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: You must pass one filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	file, err :=os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
