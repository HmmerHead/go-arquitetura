package main

import (
	"database/sql"

	dbApadter "github.com/HmmerHead/go-arquit/adapters/db"
	app "github.com/HmmerHead/go-arquit/application"
	_ "github.com/mattn/go-sqlite3"
)

// Posso testar via comando go run
func mainsz() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := dbApadter.NewProductDb(db)
	ProductService := app.NewProductService(productDbAdapter)
	product, _ := ProductService.Create("Prod", 30)

	ProductService.Enable(product)
}
