package database

import (
	"app/errors"
	"database/sql"
)

func scan_item(rows *sql.Rows, item *Item) error {
	err := rows.Scan(
		&item.ID,
		&item.Code,
		&item.Name,
		&item.PreparationCode,
		&item.Preparation,
		&item.Energy,
		&item.Protein,
		&item.Lipids,
		&item.Carbohydrate,
		&item.Fiber,
		&item.CategoryId,
		&item.CreatedAt,
	)

	return err
}

func scan_cantegory(rows *sql.Rows, item *Category) error {
	err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.CreatedAt,
	)

	return err
}

func FindCategoryByName(db *sql.DB, category string) (*[]Category, error) {
	var categories []Category

	rows, err := db.Query(`
		SELECT *
		FROM categories 
		WHERE name LIKE ?
	`, category+"%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category Category
		err = scan_cantegory(rows, &category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return &categories, nil
}

func FindItemByCategory(db *sql.DB, category int) (*[]Item, error) {
	var items []Item

	rows, err := db.Query(`
		SELECT *
		FROM items 
		WHERE category_id = ?
	`, category)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item Item
		err = scan_item(rows, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}

func FindItemByName(db *sql.DB, find string) (*[]Item, error) {
	var items []Item

	rows, err := db.Query(`
		SELECT *
		FROM items 
		WHERE name LIKE ?
	`, find+"%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item Item
		err = scan_item(rows, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}

func InsertCategory(db *sql.DB, name string) int64 {
	stmt, err := db.Prepare("INSERT INTO categories(name) values(?)")
	errors.CheckError(err)

	res, err := stmt.Exec(name)
	errors.CheckError(err)

	id, err := res.LastInsertId()
	errors.CheckError(err)

	return id
}

func InsertItem(db *sql.DB, category int64, item []string) int64 {
	stmt, err := db.Prepare(`INSERT INTO items(
		code, 
		name,
		preparation_code,
		preparation,
		energy,
		protein,
		lipids,
		carbohydrate,
		fiber,
		category_id
	) values(?,?,?,?,?,?,?,?,?,?)`)
	errors.CheckError(err)

	res, err := stmt.Exec(
		item[0],
		item[1],
		item[2],
		item[3],
		item[4],
		item[5],
		item[6],
		item[7],
		item[8],
		category,
	)
	errors.CheckError(err)

	id, err := res.LastInsertId()
	errors.CheckError(err)

	return id
}
