package main

import (
	"incidentflow/internal/server"
	"incidentflow/cmd"
	"os"
)


func main() {
	// Check if the first argument is "generate" (CLI mode)
	if len(os.Args) > 1 && os.Args[1] == "generate" {
		cmd.Execute() // Run CLI command
	} else {
		server.StartServer() // Otherwise, start API server
	}
}