package server

import (
	"os"

	"github.com/dominikpalatynski/ContentCreator/util"
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

	util.LoadEnv()

	s.router.Run(":"+ os.Getenv("PORT"))
}

func (s *APIServer) registerRoutes() {

}