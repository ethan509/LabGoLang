package main

import (
	"common"
	"excelcontroller"
	"fmt"
)

func main() {
	excelcontroller.Init()
	excelcontroller.ReadExcel("C:/dev/workspace/goLang/readExcelppl/220907_추석용 통합본.xlsx")

	//excelcontroller.ValidatePhoneNumber()

	fmt.Printf("Final Count:%d\n", len(common.ExcelInfos))
	fmt.Println("-------------------------------------------------------------------------")
}
