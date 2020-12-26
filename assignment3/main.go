package main

import (
	"io"
	"log"
	"os"
)

func main() {
	strSlice := os.Args
	fileName := strSlice[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Something wrong happened")
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
