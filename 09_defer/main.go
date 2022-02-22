package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// show runtime
	defer logTime()()

	pathIn := "data/in.txt"
	pathOut := "data/out.txt"

	// open and close input file
	fIn, err := os.Open(pathIn)
	if err != nil {
		fmt.Println("Reading input file error", err)
	}
	defer fIn.Close()

	// create and close output file
	fOut, err := os.Create(pathOut)
	if err != nil {
		fmt.Println("Creating output file error", err)
	}
	defer fOut.Close()
	var sCount, bCount int
	defer func() {
		fOut.Close()
		fmt.Printf("Created: %v string, %v bytes\n", sCount, bCount)
	}()

	// read input file line by line and write in output file
	var str string
	var bytes int
	s := bufio.NewScanner(fIn)
	for i := 1; s.Scan(); i++ {
		str = strconv.Itoa(i) + " " + s.Text() + "\n"
		bytes, _ = fOut.WriteString(str)
		sCount++
		bCount = bCount + bytes
	}
}

// runtime
func logTime() func() {
	t := time.Now()
	return func() {
		fmt.Printf("Duration: %s\n", time.Since(t))
	}
}
