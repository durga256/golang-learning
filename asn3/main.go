package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fileName := "./" + strings.Join(os.Args[1:], "")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
