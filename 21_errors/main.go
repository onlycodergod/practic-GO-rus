package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type errRead struct {
	message string
}

func (e *errRead) Error() string {
	return e.message
}

type errLimit struct {
	message    string
	limit      int
	lastString string
}

func (e *errLimit) Error() string {
	return e.message
}

var (
	errReadMessage = "read error"
)

func main() {
	pathIn := "data/in.txt"

	lineCount, err := readAndCountLines(pathIn)

	if err != nil {
		var eRead *errRead
		var errLimit *errLimit

		isErrRead := errors.As(err, &eRead)
		isErrLimit := errors.As(err, &errLimit)
		if isErrLimit {
			fmt.Printf("%s, limit: %d, last string: %s\n", errLimit.message, errLimit.limit, errLimit.lastString)
			return
		} else if isErrRead {
			log.Printf("read error: %v\n", err)
			return
		}
	}
	fmt.Printf("Total string: %d\n", lineCount)
}

func readAndCountLines(pathIn string) (int, error) {
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

	// count number of lines and check limit
	var lineCount int
	s := bufio.NewScanner(fIn)
	limit := 100
	for s.Scan() {
		lineCount++
		if lineCount > limit {
			eLimit := errLimit{
				"string count exceed limit, please read another file =)",
				limit,
				s.Text(),
			}
			return lineCount, fmt.Errorf("%w", &eLimit)
		}
	}

	// check read error
	r := bufio.NewReader(fIn)
	for i := 1; i <= lineCount; i++ {
		_, _, err := r.ReadLine()
		if err != nil && err != io.EOF {
			return 0, fmt.Errorf("readAndCountLines error: %w", &errRead{message: errReadMessage})
		}
	}
	return lineCount, nil
}
