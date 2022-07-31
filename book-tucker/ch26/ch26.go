package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readArgs() (string, string) {
	allArgs := os.Args[1:]
	logMsgUsage := "[Usage]\n\t ch26.exe [finding word] [file name]"

	if len(allArgs) < 2 {
		log.Fatal(logMsgUsage)
	} else {
		fmt.Printf("Args=%v\n", allArgs)
	}

	return allArgs[0], allArgs[1]
}

func main() {
	// Read Arguments (word & file name)
	findingWord, fileName := readArgs()
	fmt.Println("Args: ", findingWord, fileName)

	// File Open & Read
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Searching word
	scanner := bufio.NewScanner(file)
	lineNo := 1
	for scanner.Scan() {

		line := scanner.Text()
		if strings.Contains(line, findingWord) {
			fmt.Printf("%d\t%s\n", lineNo, line)
		}
		lineNo++
	}
}
