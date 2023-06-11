package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func DatabaseConnection(host, port, user, dbname, password, sslmode string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return db, nil
}
