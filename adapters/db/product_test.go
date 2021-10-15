package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/HmmerHead/go-arquit/adapters/db"
	app "github.com/HmmerHead/go-arquit/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products("id" string,"name" string,"price" float,"status" string);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abs", "Prod Test", 3, "disabled")`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abs")
	require.Nil(t, err)
	require.Equal(t, "Prod Test", product.GetName())
	require.Equal(t, 3.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := app.NewProduct()
	product.Name = "Product"
	product.Price = 25
	product.Status = "disabled"

	prodResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, prodResult.GetName())
	require.Equal(t, product.Price, prodResult.GetPrice())
	require.Equal(t, product.Status, prodResult.GetStatus())

	product.Status = "enabled"

	prodResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, prodResult.GetName())
	require.Equal(t, product.Price, prodResult.GetPrice())
	require.Equal(t, product.Status, prodResult.GetStatus())
}
