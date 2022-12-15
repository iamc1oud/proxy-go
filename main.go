package main

import (
	"github.com/proxy-go/server"
)

func main() {
	// Start server with TLS
	server.InitializeServer("https://dev.famstar.in")
}
