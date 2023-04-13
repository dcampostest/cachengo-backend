package models

type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	Category    Category `json:"references:ID"`
}

type AllProducts []Product
