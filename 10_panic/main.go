package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pathIn := "data/in.txt"
	pathOut := "data/data_out.txt"

	// check panic and show output file
	defer func() {
		if panicErr := recover(); panicErr != "" {
			fmt.Println(panicErr)
			showFile(pathOut)
		}
	}()

	// open and close input file
	fIn, err := os.Open(pathIn)
	if err != nil {
		fmt.Println("Reading input file error", err)
	}
	defer fIn.Close()

	// create and close output file
	fOut, err := os.Create(pathOut)
	if err != nil {
		fmt.Print("Creating output file error", err)
	}
	defer fOut.Close()

	// read input file line by line and write in output file
	var str string
	var strData []string
	var Name, Address, City string
	var panicStr string

	s := bufio.NewScanner(fIn)
	for i := 1; s.Scan(); i++ {
		strData = strings.Split(s.Text(), "|")
		Name = strData[0]
		Address = strData[1]
		City = strData[2]

		if Name != "" && Address != "" && City != "" {
			str = fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", i, Name, Address, City)
			fOut.WriteString(str)
		} else {
			panicStr = fmt.Sprintf("parse error: empty field on string %d", i)
			panic(panicStr)
		}
	}
}

func showFile(path string) {
	f, _ := os.Open(path)
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
