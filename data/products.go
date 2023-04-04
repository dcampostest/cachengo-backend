package data

import "github.com/pedidosya/golan-rest-simple/models"

var Products = models.AllProducts{
	{
		ID:          1,
		Name:        "Coca Cola",
		Description: "Coca cola 50 ml",
		Price:       400.00,
	},
	{
		ID:          2,
		Name:        "Combo de fernet",
		Description: "Combo de fernet brancha + coca cola 2.5lts",
		Price:       4500,
	},
	{
		ID:          3,
		Name:        "Limonada",
		Description: "Sumo de limon exprimido",
		Price:       500,
	},
}
