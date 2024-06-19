package main

import "github.com/dominikpalatynski/ContentCreator/content"

// import "github.com/dominikpalatynski/ContentCreator/server"

func main() {
	// server := server.NewAPIServer()

	// server.Run()
	content.GenerateContent("Example Prompt")
}