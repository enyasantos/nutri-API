package main

import (
	"app/database"
	"app/errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//create_database()

	db, err := database.Connection()
	errors.CheckError(err)

	category := "frutas"
	categories := database.FindCategoryByName(db, category)
	for _, category := range categories {
		fmt.Println(
			category.ID,
			category.Name,
			category.CreatedAt,
		)
	}

	find := "arroz i"
	items := database.FindItemByName(db, find)
	for _, item := range items {
		fmt.Println(
			item.ID,
			item.Code,
			item.Name,
			item.PreparationCode,
			item.Preparation,
			item.Energy,
			item.Protein,
			item.Lipids,
			item.Carbohydrate,
			item.Fiber,
			item.CategoryId,
			item.CreatedAt,
		)
	}

	category_id := categories[0].ID
	items = database.FindItemByCategory(db, category_id)
	for _, item := range items {
		fmt.Println(
			item.ID,
			item.Code,
			item.Name,
			item.PreparationCode,
			item.Preparation,
			item.Energy,
			item.Protein,
			item.Lipids,
			item.Carbohydrate,
			item.Fiber,
			item.CategoryId,
			item.CreatedAt,
		)
	}

	database.CloseConnection(db)

}
