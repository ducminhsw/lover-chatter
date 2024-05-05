package server

import (
	"database/sql"
	"log/slog"
)

type ServerConfig struct {
	ListenAddr string
	Logger     *slog.Logger
	Database   *sql.DB
}

type Server struct {
	Conf ServerConfig
}

func DefaultServer() *Server {
	return &Server{
		Conf: ServerConfig{
			ListenAddr: "8000",
			Logger:     &slog.Logger{},
			Database:   nil,
		},
	}
}

func NewServer(cfg ServerConfig) *Server {
	server := DefaultServer()
	if len(cfg.ListenAddr) > 0 {
		server.Conf.ListenAddr = cfg.ListenAddr
	}
	if cfg.Logger != nil {
		server.Conf.Logger = cfg.Logger
	}
	if cfg.Database != nil {
		server.Conf.Database = cfg.Database
	}
	return server
}
