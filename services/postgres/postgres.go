package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "ludvig"
	password = "ludvig"
	dbname   = "ludvig"
)

func NewPostgres() *sql.DB {
	fmt.Println(os.Getenv("DATABASE_URL"))
	psql := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
