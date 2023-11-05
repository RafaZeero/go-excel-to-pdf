package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("teste_go.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get value from cell by given worksheet name and cell reference.
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		fmt.Println("linha", i+1)

		for _, colCell := range row {
			// fmt.Printf("type: %T, value: %v", colCell, colCell)
			num, err := strconv.Atoi(colCell)
			if err != nil {
				fmt.Println(err)
				break
			}

			// if num is type integer, multiply by 2
			if reflect.TypeOf(num).Kind() == reflect.Int {
				fmt.Printf("type: %T, value: %v\n", num, num)
				num *= 2 // multiply by 2
				fmt.Printf("value: %v\n", num)
			}
		}
		fmt.Println()
	}
}
