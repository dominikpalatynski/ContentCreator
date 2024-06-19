package server

import (
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	router *gin.Engine
}

func NewAPIServer() *APIServer{
	r := gin.Default()

	server := &APIServer{
		router: r,
	}

	return server
}

func (s *APIServer) Run() {

	s.registerRoutes()


	s.router.Run("6000")
}

func (s *APIServer) registerRoutes() {

}