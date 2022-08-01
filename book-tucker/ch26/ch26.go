package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func readArgs() (string, []string) {
	allArgs := os.Args[1:]
	logMsgUsage := "[Usage]\n\t ch26.exe [finding word] [file name1] ..."

	if len(allArgs) < 2 {
		log.Fatal(logMsgUsage)
	} else {
		fmt.Printf("Args=%v\n", allArgs)
	}

	return allArgs[0], allArgs[1:]
}

func SearchingWord(wg *sync.WaitGroup, findingWord string, fileName string) {

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
		//fmt.Println("Args: ", line)
		if strings.Contains(line, findingWord) {
			fmt.Printf("[%s]\t%d\t%s\n", fileName, lineNo, line)
		}
		lineNo++
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// Read Arguments (word & file name)
	findingWord, fileNames := readArgs()
	fmt.Println("Args: ", findingWord, fileNames)

	wg.Add(len(fileNames))

	for _, fileName := range fileNames {
		go SearchingWord(&wg, findingWord, fileName)
	}
	wg.Wait()

}
