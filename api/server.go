package api

import (
	db "github.com/Al3xDo/simple_bank/db/sqlc"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// server serves HTTP requests for out banking services
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routine
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// register a custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validaCurrency)
	}

	// add routes to router
	// if you pass multiple funcs -> last one is handler and others are middleware
	router.POST("/accounts", server.CreateAccount)
	router.GET("/account/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccount)

	router.POST("/transfers", server.createTransfer)

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
