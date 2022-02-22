package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	pathIn := "data/in.txt"

	// open input file and check errors
	fIn, err := os.Open(pathIn)
	if err != nil {
		fmt.Printf("open input file error: %v\n", err)
	}

	// close input file and check errors
	defer func() {
		if err := fIn.Close(); err != nil {
			fmt.Printf("close input file error: %v\n", err)
		}
	}()

	// count number of lines
	var lineCount int
	s := bufio.NewScanner(fIn)
	for s.Scan() {
		lineCount++
	}

	// check EOF error
	r := bufio.NewReader(fIn)
	for i := 1; i <= lineCount; i++ {
		_, _, err := r.ReadLine()
		if err != nil && err != io.EOF {
			fmt.Printf("read error: %v\n", err)
		}
	}

	fmt.Printf("Total string: %d\n", lineCount)
}
