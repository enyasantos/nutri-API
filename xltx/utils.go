package xltx

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func nextChar(ch byte) byte {
	if ch += 1; ch > 'z' {
		return 'a'
	}
	return ch
}

func is_category(f *excelize.File, sheet_name string, row int) (bool, string, error) {
	var values []string

	cells := []string{
		"A" + strconv.Itoa(row),
		"B" + strconv.Itoa(row),
	}

	for _, c := range cells {
		temp, err := f.GetCellValue(sheet_name, c)
		if err != nil {
			fmt.Println(err)
			return false, "", err
		}

		values = append(values, temp)
	}

	if values[0] == values[1] {
		return true, values[0], nil
	}

	return false, "", nil
}
