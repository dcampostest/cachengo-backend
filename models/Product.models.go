package models

type Product struct {
	ID          int     `json:"ID"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
}

type AllProducts []Product
