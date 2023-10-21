package server

import (
	"net/http"

	"github.com/ericoliveiras/gate-guard/internal/config"
	"github.com/ericoliveiras/gate-guard/internal/db"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DB     *sqlx.DB
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		DB:     db.Init(config),
		Config: config,
	}
}

func (server *Server) Start(addr string) error {
	return http.ListenAndServe(":"+addr, nil)
}
