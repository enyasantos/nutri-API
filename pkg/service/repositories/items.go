package repositories

import (
	"app/database"
	"database/sql"
)

type items struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *items {
	return &items{db}
}

func (r items) FindItemByName(name string) (*[]database.Item, error) {
	items, error := database.FindItemByName(r.db, name)
	if error != nil {
		return nil, error
	}

	return items, nil
}

func (r items) FindItemByCategory(category int) (*[]database.Item, error) {
	items, error := database.FindItemByCategory(r.db, category)
	if error != nil {
		return nil, error
	}

	return items, nil
}
