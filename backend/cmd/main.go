package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.ducminhsw.prepare-project/internal/database"
	"github.ducminhsw.prepare-project/internal/server"
)

func main() {
	// init
	addr := flag.String("addr", "localhost:8000", "HTTP port network address")
	dbcn := flag.String("dbcn", "postgres://postgres:Taochiunhe1712@@localhost/beloved-database?sslmode=disable", "Database connection address")
	flag.Parse()

	// create config params
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	sqlConn := database.OpenDatabaseConn(*dbcn, *logger)
	defer sqlConn.Close()

	// make the configuration
	config := server.ServerConfig{
		ListenAddr: *addr,
		Logger:     logger,
		Database:   sqlConn,
	}

	svr := server.NewServer(config)

	server := http.Server{
		Addr:         (*svr).Conf.ListenAddr,
		Handler:      svr.RegisterRoutes(),
		ErrorLog:     slog.NewLogLogger(svr.Conf.Logger.Handler(), slog.LevelError),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("server shutdown", err)
		os.Exit(1)
	}
}
