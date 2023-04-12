package datasource

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pedidosya/golan-rest-simple/models"
)

/* type Datasource struct {
	config *models.DbConfiguration
	db     *sqlx.DB
} */

const (
	getListProducs    string = "SELECT * FROM products"
	getProductByID    string = ""
	getListCategories string = "SELECT * FROM categories"
	getCategoryByID   string = "SELECT * FROM categories WHERE id=?"
	createCategory    string = "INSERT INTO categories(name, description) VALUE (?, ?)"
	deleteCategory    string = "DELETE FROM categories WHERE id=?"
	updatecategory    string = "UPDATE categories SET name=?, description=? WHERE id=?"
)

/*var db = []models.DbConfiguration{
	{User: "root", Pwd: "Ubuntu1234", Server: "localhost", Database: "db_cachengo",Port: 3306},
}*/

func DBConnect() (connectdb *sqlx.DB) {
	Driver := "mysql"
	Host := "tcp(127.0.0.1:3306)"
	User := "root"
	Password := "Ubuntu1234$"
	DBName := "db_cachengo"

	connectdb, err := sqlx.Open(Driver, fmt.Sprintf("%s:%s@%s/%s", User, Password, Host, DBName))
	if err != nil {
		panic(err.Error())
	}
	if err = connectdb.Ping(); err != nil {
		return nil
	}

	return connectdb
}

// GetUserByToken returns user by token
func GetListProducts() *[]models.Product {

	db := DBConnect()
	listProductsData := []models.Product{}
	err := db.Select(&listProductsData, getListProducs)

	if err != nil {
		panic(err.Error())
	}
	return &listProductsData
}

func GetListCategories() []models.Category {

	db := DBConnect()
	category := models.Category{}
	listAllCategories := []models.Category{}

	list, err := db.Query(getListCategories)
	if err != nil {
		panic(err.Error())
	}

	for list.Next() {
		var id int
		var name, description string
		err = list.Scan(&id, &name, &description)
		if err != nil {
			panic(err.Error())
		}
		category.ID = id
		category.Name = name
		category.Description = description

		listAllCategories = append(listAllCategories, category)
	}
	// fmt.Println(listAllCategories)
	return listAllCategories
}

func CreateCategory(name string, description string) {
	dbconnect := DBConnect()
	insertarRegistro, err := dbconnect.Prepare(createCategory)
	if err != nil {
		panic(err.Error())
	}
	insertarRegistro.Exec(name, description)
}

func DeleteCategory(id string) {
	dbconnect := DBConnect()
	delRegistro, err := dbconnect.Prepare(deleteCategory)
	if err != nil {
		panic(err.Error())
	}
	delRegistro.Exec(id)
}

func GetCategoryByID(id string) models.Category {
	category := models.Category{}
	dbconnect := DBConnect()
	getCategoryByID, err := dbconnect.Query(getCategoryByID, id)
	if err != nil {
		panic(err.Error())
	}
	for getCategoryByID.Next() {
		var id int
		var name, description string
		err = getCategoryByID.Scan(&id, &name, &description)
		if err != nil {
			panic(err.Error())
		}
		category.ID = id
		category.Name = name
		category.Description = description
	}
	// fmt.Println(category)
	return category
}

func UpdateCategory(id string, name string, description string) {
	dbconnect := DBConnect()
	updateRegistro, err := dbconnect.Prepare(updatecategory)
	if err != nil {
		panic(err.Error())
	}
	updateRegistro.Exec(name, description, id)
}
