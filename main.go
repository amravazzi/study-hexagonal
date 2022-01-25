package main

import (
	"database/sql"

	db2 "github.com/amravazzi/study-hexagonal/adapters/db"

	"github.com/amravazzi/study-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	productService.Create("Product Example", 30)
}