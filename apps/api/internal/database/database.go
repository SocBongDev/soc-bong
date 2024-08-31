package database

import (
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/pocketbase/dbx"
	"go.nhat.io/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func New(cfg *config.DatabaseSecret) (*dbx.DB, error) {
	driverName, err := otelsql.Register(
		"libsql",
		otelsql.AllowRoot(),
		otelsql.TraceQueryWithoutArgs(),
		otelsql.TraceRowsClose(),
		otelsql.TraceRowsAffected(),
		// otelsql.WithDatabaseName("my_database"),        // Optional.
		otelsql.WithSystem(semconv.DBSystemSqlite), // Optional.
	)
	if err != nil {
		return nil, err
	}

	db, err := dbx.MustOpen(driverName, cfg.GetUrl())
	if err != nil {
		return nil, err
	}
	db.LogFunc = logger.Info

	return db, nil
}
