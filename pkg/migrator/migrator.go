package migrator

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"

	"github.com/pressly/goose/v3"
)

type Migrator struct {
	db         *sql.DB
	migrations fs.FS
}

func New(db *sql.DB, migrations fs.FS) *Migrator { //конструктор
	return &Migrator{
		db:         db,
		migrations: migrations,
	}
}

func (m *Migrator) Up() error { //поднимаем миграции
	if err := m.setup(); err != nil {
		return fmt.Errorf("migrator.Up: %w", err)
	}
	if err := goose.Up(m.db, "."); err != nil {
		return fmt.Errorf("migrator: failed to up: %w", err)
	}
	return nil
}

func (m *Migrator) setup() error { //настройки мигратора
	goose.SetLogger(goose.NopLogger()) //отключает логирование для миграций
	goose.SetTableName("schema_migrations")

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return fmt.Errorf("migrator.setup %w", err)
	}

	goose.SetBaseFS(m.migrations)

	return nil

}

func (m *Migrator) Down() error { //роняем миграции
	if err := m.setup(); err != nil {
		return fmt.Errorf("migrator.Down: %w", err)
	}

	if err := goose.Down(m.db, "."); err != nil {
		return fmt.Errorf("migrator: failed to down: %w", err)
	}

	return nil

}

/*
//go:embed migrations/*.sql
var embedMigrations embed.FS
*/

func EmbedMigrations(db *sql.DB, efs embed.FS, dir string) (*Migrator, error) {
	sub, err := fs.Sub(efs, dir)
	if err != nil {
		return nil, fmt.Errorf("migrator.EmbedMigrations: %w", err)
	}

	return New(db, sub), nil
}
