package main

import "github.com/dominikpalatynski/ContentCreator/server"

func main() {
	server := server.NewAPIServer()

	server.Run()
}