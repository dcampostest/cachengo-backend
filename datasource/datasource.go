package datasource

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pedidosya/golan-rest-simple/models"
)

/* type Datasource struct {
	config *models.DbConfiguration
	db     *sqlx.DB
} */

const (
	getListCategories    string = "SELECT * FROM categories"
	getCategoryByID      string = "SELECT * FROM categories WHERE id=?"
	getCategoryByName    string = "SELECT * FROM categories WHERE name=?"
	createCategory       string = "INSERT INTO categories(name, description) VALUE (?, ?)"
	deleteCategory       string = "DELETE FROM categories WHERE id=?"
	updatecategory       string = "UPDATE categories SET name=?, description=? WHERE id=?"
	getListProducts      string = "SELECT products.id, products.name, products.description, products.price, categories.name FROM products INNER JOIN categories ON products.id_category=categories.id;"
	getProductByID       string = "SELECT * FROM products WHERE id=?"
	getProductByCategory string = "SELECT products.id, products.name, products.description, products.price, products.id_category, categories.name FROM categories INNER JOIN products ON categories.id=products.id_category"
	createProduct        string = "INSERT INTO products(name, description, price, id_category) VALUE (?, ?, ?, ?)"
	deleteProduct        string = "DELETE FROM products WHERE id=?"
	updateProduct        string = "UPDATE products SET name=?, description=?, price=?, id_category=? WHERE id=?"
)

// DATASOURCE Categories

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
	return category
}

func GetCategoryByName(name string) models.Category {
	category := models.Category{}
	dbconnect := DBConnect()
	getCategoryByName, err := dbconnect.Query(getCategoryByName, name)
	if err != nil {
		panic(err.Error())
	}
	for getCategoryByName.Next() {
		var id int
		var name, description string
		err = getCategoryByName.Scan(&id, &name, &description)
		if err != nil {
			panic(err.Error())
		}
		category.ID = id
		category.Name = name
		category.Description = description
	}
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

// DATASOURCE Products

func GetListProducts() []models.Product {

	db := DBConnect()
	product := models.Product{}
	listAllProducts := []models.Product{}

	list, err := db.Query(getListProducts)
	if err != nil {
		panic(err.Error())
	}
	for list.Next() {
		var id int
		var name, description, nameCategory string
		var price float32

		err = list.Scan(&id, &name, &description, &price, &nameCategory)
		if err != nil {
			panic(err.Error())
		}
		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Category.Name = nameCategory

		listAllProducts = append(listAllProducts, product)
	}
	return listAllProducts
}

func GetListProductsByCategory() []models.Product {

	db := DBConnect()
	product := models.Product{}
	listAllProducts := []models.Product{}

	list, err := db.Query(getProductByCategory)
	if err != nil {
		panic(err.Error())
	}
	for list.Next() {
		var id, id_category int
		var name, description, nameCategory string
		var price float32
		err = list.Scan(&id, &name, &description, &price, &id_category, &nameCategory)
		if err != nil {
			panic(err.Error())
		}
		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Category.ID = id_category
		product.Category.Name = nameCategory

		listAllProducts = append(listAllProducts, product)
	}
	return listAllProducts
}

func GetAll() any {

	listAllProductsByCategories := make(map[string]models.AllProducts)

	listCate := GetListCategories()
	listPro := GetListProductsByCategory()

	for _, category := range listCate {
		listAllProducts := []models.Product{}

		for _, product := range listPro {

			if product.Category.ID == category.ID {
				listAllProducts = append(listAllProducts, product)
			}
		}
		var namecategory = category.Name
		listAllProductsByCategories[namecategory] = listAllProducts

	}
	return listAllProductsByCategories
}

func CreateProduct(name string, description string, price string, idcategory string) {
	dbconnect := DBConnect()
	insertarRegistro, err := dbconnect.Prepare(createProduct)
	if err != nil {
		panic(err.Error())
	}
	insertarRegistro.Exec(name, description, price, idcategory)
}

func DeleteProduct(id string) {
	dbconnect := DBConnect()
	delRegistro, err := dbconnect.Prepare(deleteProduct)
	if err != nil {
		panic(err.Error())
	}
	delRegistro.Exec(id)
}

func GetProductByID(id string) models.Product {
	product := models.Product{}
	category := models.Category{}
	dbconnect := DBConnect()
	getProductByID, err := dbconnect.Query(getProductByID, id)
	if err != nil {
		panic(err.Error())
	}
	for getProductByID.Next() {
		var id, id_category int
		var name, description string
		var price float32
		err = getProductByID.Scan(&id, &name, &description, &price, &id_category)
		if err != nil {
			panic(err.Error())
		}
		category = GetCategoryByID(strconv.Itoa(id_category))
		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Category.ID = id_category
		product.Category.Name = category.Name
	}
	return product
}

func UpdateProduct(id string, name string, description string, price string, id_category string) {
	dbconnect := DBConnect()
	updateRegistro, err := dbconnect.Prepare(updateProduct)
	if err != nil {
		panic(err.Error())
	}
	updateRegistro.Exec(name, description, price, id_category, id)
}
