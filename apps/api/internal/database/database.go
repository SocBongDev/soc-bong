package database

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/pocketbase/dbx"
)

func New(cfg *config.DatabaseSecret) (*dbx.DB, error) {
	db, err := dbx.MustOpen("libsql", cfg.Url)
	if err != nil {
		return nil, err
	}
	db.LogFunc = log.Printf

	return db, nil
}
