package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"github.com/QBasy/bailanys-server-go/pkg/logger"
)

func Connect(dsn string, log logger.Logger) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error("db connect failed", logger.Err(err), logger.String("dsn", redactDSN(dsn)))
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Error("db ping failed", logger.Err(err), logger.String("dsn", redactDSN(dsn)))
		panic(err)
	}

	log.Info("db connected")
	return db
}

func redactDSN(dsn string) string {
	if len(dsn) > 64 {
		return dsn[:64] + "..."
	}
	return dsn
}
