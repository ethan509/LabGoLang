package main

import (
	"bufio"
	"common"
	"excelcontroller"
	"fmt"
	"log"
	"os"
	"strings"
)

func getInputFileName() string {
	fmt.Printf("input source file name(relative path):")

	in := bufio.NewReader(os.Stdin)
	inputStr, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputStr = strings.TrimSuffix(inputStr, "\n")
	inputStr = strings.TrimSuffix(inputStr, "\r")

	path, _ := os.Getwd()
	fmt.Println(path)

	iputFileName := path + "\\" + inputStr

	return iputFileName
}

func main() {
	iputFileName := getInputFileName()

	excelcontroller.Init()
	excelcontroller.ReadExcel(iputFileName)

	fmt.Printf("Final validated Phone number:%d\n", len(common.ExcelInfos))
	fmt.Println("-------------------------------------------------------------------------")
}
