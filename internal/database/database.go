package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Server()
}

type service struct {
	db *sql.DB
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() Service {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",username, password, host, port, database)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error:", err)
	}
	s := &service{db: db}
	return s
}

func (s *service)Server()  {
}
