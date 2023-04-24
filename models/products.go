package models

import "github.com/cxmarcelo/alura-go-web/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAllProducts() []Product {
	db := db.ConnectDatabase()

	selecAllProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selecAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selecAllProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func InsertProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	insert, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	deleteQuery, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteQuery.Exec(id)

	defer db.Close()
}

func FindById(id string) Product {
	db := db.ConnectDatabase()

	productDb, err := db.Query("select * from products where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDb.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()
	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	updateScript, err := db.Prepare("update products set name = $1, description = $2, price = $3, quantity = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	updateScript.Exec(name, description, price, quantity, id)

	defer db.Close()
}
