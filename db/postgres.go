package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Postgres struct {
	Conn *pgx.Conn
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Open(ctx context.Context, url string) (*pgx.Conn, error) {
	var err error
	p.Conn, err = pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("Database connection failed %v", err)
	}

	err = p.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("Database connection dead %v", err)
	}

	defer p.Close(ctx)

	return p.Conn, nil
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.Conn.Ping(ctx)
}

func (p *Postgres) Close(ctx context.Context) error {
	return p.Conn.Close(ctx)
}
