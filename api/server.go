package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// server serves HTTP requests for out banking services
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routine
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	// if you pass multiple funcs -> last one is handler and others are middleware
	router.POST("/accounts", server.CreateAccount)
	router.GET("/account/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccount)
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
