package database

import "time"

type Category struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Item struct {
	ID              int       `json:"id,omitempty"`
	Code            string    `json:"code,omitempty"`
	Name            string    `json:"name,omitempty"`
	PreparationCode int32     `json:"preparation_code,omitempty"`
	Preparation     string    `json:"preparation,omitempty"`
	Energy          string    `json:"energy,omitempty"`
	Protein         string    `json:"protein,omitempty"`
	Lipids          string    `json:"lipids,omitempty"`
	Carbohydrate    string    `json:"carbohydrate,omitempty"`
	Fiber           string    `json:"fiber,omitempty"`
	CategoryId      int64     `json:"category_id,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}
