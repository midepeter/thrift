package db

import (
	"log"

	m "github.com/golang-migrate/migrate/v4"
)

type Migrate struct {
	M *m.Migrate
}

func New(sourceUrl, databaseUrl string) *Migrate {
	m, err := m.New(sourceUrl, databaseUrl)
	if err != nil {
		return nil
	}

	m.Steps(2)
	return  &Migrate{
		M: m,
	}
}

func(m *Migrate) Up() error {
	var err error 
	defer m.Close()
	if m == nil {
		log.Fatalf("Migrate instance invalid")
	}

	if err = m.Up(); err != nil {
		return err
	}
	return nil
}

func (m *Migrate) Down() error {
	var err error 
	if m == nil {
		log.Fatalf("Migrate instance invalid")
	}

	if err = m.Down(); err != nil {
		return err
	}
	return nil
}

func (m *Migrate) Close() {
	m.Close()
}