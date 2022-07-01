package db

import (
	"context"
	"database/sql"
	"errors"
)

type Db struct {
	Db *sql.DB
}

func (d *Db) Setup(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("sql", dsn)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, errors.New("Database connection not available")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	d.Db = db
	return db, nil
}

//Database methods
func (d *Db) Insert(ctx context.Context, stmt string) error {
	_, err := d.Db.PrepareContext(ctx, stmt)
	if err != nil {
		return err
	}

	result, err := d.Db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	r, _ := result.RowsAffected()
	if r == 0 {
		return errors.New("Database query failed")
	}
	return nil
}

func (d *Db) Update(ctx context.Context, stmt string) error {
	_, err := d.Db.PrepareContext(ctx, stmt)
	if err != nil {
		return err
	}

	result, err := d.Db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	r, _ := result.RowsAffected()
	if r == 0 {
		return errors.New("Database query failed")
	}
	return nil
}

func (d *Db) Delete(ctx context.Context, stmt string) error {
	_, err := d.Db.PrepareContext(ctx, stmt)
	if err != nil {
		return err
	}

	result, err := d.Db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	r, _ := result.RowsAffected()
	if r == 0 {
		return errors.New("Database query failed")
	}
	return nil
}

func (d *Db) Select(ctx context.Context, stmt string) error {
	_, err := d.Db.PrepareContext(ctx, stmt)
	if err != nil {
		return err
	}

	result, err := d.Db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	r, _ := result.RowsAffected()
	if r == 0 {
		return errors.New("Database query failed")
	}
	return nil
}
