package models

import (
	"database/sql"

	"gopkg.in/boj/redistore.v1"
)

type Model struct {
	DB        *sql.DB
	RediStore *redistore.RediStore
}
