package database

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type migrator struct {
	m *migrate.Migrate
}

var _ Migrator = migrator{}

func (this migrator) Up() error {
	return this.m.Up()
}

func (this migrator) Down(steps int) error {
	return this.m.Steps(-steps)
}

func (this migrator) Force(version int) error {
	return this.m.Force(version)
}

func NewMigrator(cfg *config.DatabaseSecret) (Migrator, error) {
	dbx, err := New(cfg)
	if err != nil {
		log.Fatalln("database.New err: ", err)
		return nil, err
	}

	db := dbx.DB()
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalln("sqlite3.WithInstance err: ", err)
		return nil, err
	}

	fSrc, err := (&file.File{}).Open("./migrations")
	if err != nil {
		log.Fatalln("file.Open err: ", err)
		return nil, err
	}
	m, err := migrate.NewWithInstance(
		"file",
		fSrc,
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatalln("migrate.NewWithInstance err: ", err)
		return nil, err
	}

	return &migrator{m}, nil
}
