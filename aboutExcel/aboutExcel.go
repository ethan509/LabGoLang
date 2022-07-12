package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	x1, err := excelize.OpenFile("./lotto_1013.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}

	str := x1.GetCellValue("lotto_1013", "N3")
	fmt.Println(str)

	fmt.Println("SheetCount=", x1.SheetCount)
	fmt.Println("Sheet=", x1.Sheet)
	fmt.Println("GetSheetMap=", x1.GetSheetMap())
	fmt.Println("GetActiveSheetIndex=", x1.GetActiveSheetIndex())

	// 당첨번호 출력
	rows := x1.GetRows("lotto_1013")
	for _, row := range rows {
		if len(row[1]) != 0 {
			fmt.Printf("row[%s]: %s, %s, %s, %s, %s, %s\n", row[1], row[13], row[14], row[15], row[16], row[17], row[18])
		}
	}

}
