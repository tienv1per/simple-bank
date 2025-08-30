package api

import (
	"github.com/gin-gonic/gin"
	db "simple-bank/db/sqlc"
)

// Server serves HTTP for banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfer", server.createTransfer)

	server.router = router

	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
