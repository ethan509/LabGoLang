package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type LineInfo struct {
	LineNo int
	Line   string
}

type FindInfo struct {
	FileName string
	lines    *list.List
}

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

func SearchingWord(wg *sync.WaitGroup, ch chan FindInfo, findInfos *list.List, findingWord string, fileName string) {

	// File Open & Read
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var findInfo FindInfo
	findInfo.FileName = fileName

	lineInfos := list.New()
	var lineInfo LineInfo

	// Searching word
	scanner := bufio.NewScanner(file)
	lineNo := 1
	for scanner.Scan() {

		line := scanner.Text()
		//fmt.Println("Args: ", line)
		if strings.Contains(line, findingWord) {

			lineInfo.LineNo = lineNo
			lineInfo.Line = line

			lineInfos.PushBack(lineInfo)

			fmt.Printf("[%s]\t%d\t%s\n", fileName, lineInfo.LineNo, lineInfo.Line)
		}
		lineNo++
	}
	findInfo.lines = lineInfos

	ch <- findInfo
	findInfos.PushBack(findInfo)

	wg.Done()
}

func Result(ch chan FindInfo) {
	for findInfo := range ch {
		for e := findInfo.lines.Front(); e != nil; e = e.Next() {
			fmt.Printf("[%s]\t%v\n", findInfo.FileName, e.Value)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Read Arguments (word & file name)
	findingWord, fileNames := readArgs()
	fmt.Println("Args: ", findingWord, fileNames)

	wg.Add(len(fileNames))
	ch := make(chan FindInfo)
	//findInfos := list.New()

	Result(ch)
	// for _, fileName := range fileNames {
	// 	go SearchingWord(&wg, &ch, findInfos, findingWord, fileName)
	// }
	wg.Wait()

}
