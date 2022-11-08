package repositories

import (
	"app/database"
	"database/sql"
)

type category struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *category {
	return &category{db}
}

func (r category) FindCategoryByName(name string) (*[]database.Category, error) {
	categories, error := database.FindCategoryByName(r.db, name)
	if error != nil {
		return nil, error
	}

	return categories, nil
}
