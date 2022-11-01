package xltx

import (
	"app/errors"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ReadFile(file File) ([]string, [][][]string) {
	f, err := excelize.OpenFile(file.Path)
	errors.CheckError(err)

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheet_name := "Sheet1"
	initial_char := byte('A')
	next_char := initial_char
	var values [][][]string
	var temp [][]string
	var items_by_category []string
	var categories []string

	for i := file.FirstCell; i <= file.LastCell; i++ {
		is_category, category, err := is_category(f, sheet_name, i)
		errors.CheckError(err)

		if is_category {
			if len(temp) > 0 {
				values = append(values, temp)
				temp = nil
			}
			categories = append(categories, category)
		} else {
			// A até I
			items_by_category = nil
			next_char = initial_char
			for j := 1; j < 10; j++ {
				cell := string(next_char) + strconv.Itoa(i)
				value, err := f.GetCellValue(sheet_name, cell)
				items_by_category = append(items_by_category, value)
				errors.CheckError(err)

				next_char = nextChar(next_char)
			}
			temp = append(temp, items_by_category)
		}

	}
	values = append(values, temp)

	return categories, values
}
