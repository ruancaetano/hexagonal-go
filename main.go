package main

import (
	"database/sql"

	"github.com/ruancaetano/hexagonal-go/adapters/db"
	"github.com/ruancaetano/hexagonal-go/application"
)

func main() {
	conn, _ := sql.Open("sqlite3", "db.sqlite")
	defer conn.Close()

	productDbAdapter := db.NewProductDb(conn)
	productService := application.NewProductService(productDbAdapter)
	productService.Create("Example", 100)
}
