package data

import (
	"github.com/pedidosya/golan-rest-simple/models"
)

var Categories = models.AllCategories{
	{
		ID:          1,
		Name:        "Gaseosas",
		Description: "gaseosas",
	},
	{
		ID:          2,
		Name:        "Combos",
		Description: "caombos",
	},
	{
		ID:          3,
		Name:        "Mocktails",
		Description: "limonadas y pomeladas",
	},
}
