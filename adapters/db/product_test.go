package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ruancaetano/hexagonal-go/adapters/db"
	"github.com/ruancaetano/hexagonal-go/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	craeteTable(Db)
	insertProducts(Db)
}

func craeteTable(db *sql.DB) {
	queryString := `CREATE TABLE products (
			"id" string,
			"name" string,
			"status" string,
			"price" float
		);`

	stmt, err := db.Prepare(queryString)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func insertProducts(db *sql.DB) {
	queryString := `INSERT INTO products (id, name, status, price) VALUES ("1","P1","disabled",0);`

	stmt, err := db.Prepare(queryString)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("1")

	require.Nil(t, err)

	require.Equal(t, "P1", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "1", product.GetID())
}

func TestProductDb_Sav(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Status = "disabled"
	product.Price = 0.0

	createdProduct, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetName(), createdProduct.GetName())
	require.Equal(t, product.GetPrice(), createdProduct.GetPrice())
	require.Equal(t, product.GetStatus(), createdProduct.GetStatus())

	product.Price = 10
	product.Enable()

	updatedProduct, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetName(), updatedProduct.GetName())
	require.Equal(t, product.GetPrice(), updatedProduct.GetPrice())
	require.Equal(t, product.GetStatus(), updatedProduct.GetStatus())
}
