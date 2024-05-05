package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func OpenDatabaseConn(urlstr string, logger slog.Logger) *sql.DB {
	conn, err := sql.Open("postgres", urlstr)
	if err != nil {
		logger.Error("database connection error", err)
		os.Exit(1)
	}
	return conn
}
