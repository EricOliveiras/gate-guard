package server

import (
	"log"
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
	conn, err := db.Init(config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return &Server{
		DB:     conn,
		Config: config,
	}
}

func (server *Server) Start(addr string) error {
	return http.ListenAndServe(":"+addr, nil)
}
