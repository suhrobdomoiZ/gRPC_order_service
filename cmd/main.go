package main

import (
	"homework/internal/server"
	"os"
)

func main() {

	myServer := server.NewServer("./config/.env")

	if err := myServer.Run(); err != nil {
		myServer.Logger.Error("main: failed to serve", "error", err.Error())
		os.Exit(1)
	}
}
