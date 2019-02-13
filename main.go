package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: parser <file>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	FatalIfError(err)

	logScanner := NewLogScanner(bufio.NewReader(file))
	logScanner.ScanAll()
	fmt.Println(logScanner.LogStatistics.String())
}
