package models

import "time"

type Product struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
