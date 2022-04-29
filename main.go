package main

import (
	"activity/internal/server"
)

func main() {
	server.NewHTTPServer(":8080").ListenAndServe()
}
