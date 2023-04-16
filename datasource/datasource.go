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
	getListProducts      string = "SELECT * FROM products"
	getProductByID       string = "SELECT * FROM products WHERE id=?"
	getProductByCategory string = "SELECT * FROM products WHERE id_category=?"
	createProduct        string = "INSERT INTO products(name, description, price, id_category) VALUE (?, ?, ?, ?)"
	deleteProduct        string = "DELETE FROM products WHERE id=?"
	updateProduct        string = "UPDATE products SET name=?, description=?, price=?, id_category=? WHERE id=?"
)

// DATASOURCE Categories

func DBConnect() (db *sqlx.DB) {
	Driver := "mysql"
	Host := "tcp(127.0.0.1:3306)"
	User := "root"
	Password := "Ubuntu1234$"
	DBName := "db_cachengo"

	db, err := sqlx.Open(Driver, fmt.Sprintf("%s:%s@%s/%s", User, Password, Host, DBName))
	if err != nil {
		return nil
	}

	if db.Ping() != nil {
		return nil
	}

	db.SetMaxOpenConns(100)  // The default is 0 (unlimited)
	db.SetMaxIdleConns(50)   // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(0) // 0, connections are reused forever.

	return db
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
	category := models.Category{}

	list, err := db.Query(getListProducts)
	if err != nil {
		panic(err.Error())
	}
	for list.Next() {
		var id, id_category int
		var name, description string
		var price float32
		err = list.Scan(&id, &name, &description, &price, &id_category)
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

		listAllProducts = append(listAllProducts, product)
	}
	return listAllProducts
}

func GetListProductsByCategory(id_category string) []models.Product {

	db := DBConnect()
	product := models.Product{}
	listAllProducts := []models.Product{}
	category := models.Category{}

	list, err := db.Query(getProductByCategory, id_category)
	if err != nil {
		panic(err.Error())
	}
	for list.Next() {
		var id, id_category int
		var name, description string
		var price float32
		err = list.Scan(&id, &name, &description, &price, &id_category)
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

		listAllProducts = append(listAllProducts, product)
	}
	return listAllProducts
}

func GetAll() *map[string]models.AllProducts {

	listAllProductsByCategories := make(map[string]models.AllProducts)

	listCate := GetListCategories()
	for _, category := range listCate {
		listAllProducts := []models.Product{}
		product := models.Product{}
		var idcategory = category.ID
		var namecategory = category.Name
		listPro := GetListProductsByCategory(strconv.Itoa(idcategory))

		for _, prod := range listPro {
			var id = prod.ID
			var id_category = prod.Category.ID
			var name = prod.Name
			var description = prod.Description
			var price = prod.Price

			category = GetCategoryByID(strconv.Itoa(id_category))
			product.ID = id
			product.Name = name
			product.Description = description
			product.Price = price
			product.Category.ID = id_category
			product.Category.Name = category.Name
			listAllProducts = append(listAllProducts, product)

		}
		listAllProductsByCategories[namecategory] = listAllProducts

	}
	return &listAllProductsByCategories
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
