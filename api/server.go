package api

import (
	db "bookkeeping-backend/database/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Queries
	router *gin.Engine
}

func NewServer(store *db.Queries) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)

	server.router = router

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
