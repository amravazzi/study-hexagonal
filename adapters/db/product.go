package db

import (
	"database/sql"

	"github.com/amravazzi/study-hexagonal/application"
)

type ProductDb struct {
	db *sql.DB
	_  "github.com/nattn/go-sqlite3"
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {

}
