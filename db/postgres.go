package db

import (
	"database/sql"
)

type Postgres struct {
	conn *sql.Conn
	db   *sql.DB
}

func (p *Postgres) Open(url string) (*sql.Conn, error) {
	return nil, nil
}
