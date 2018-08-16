package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgres(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
