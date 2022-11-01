package main

import (
	"app/database"
	"app/errors"
	"app/xltx"
)

func create_database() {
	file := xltx.File{
		Path:      "tab01.xlsx",
		FirstCell: 5,
		LastCell:  1994,
	}

	categories, items := xltx.ReadFile(file)

	db, err := database.Connection()
	errors.CheckError(err)

	var categories_id []int64
	for _, c := range categories {
		id := database.InsertCategory(db, c)
		categories_id = append(categories_id, id)
	}

	for i, v := range items {
		for _, item := range v {
			category_id := categories_id[i]
			database.InsertItem(db, category_id, item)
		}
	}

	database.CloseConnection(db)

}
