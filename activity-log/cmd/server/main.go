package main

import (
	"github.com/York-Shawn/activity/activity-log/internal/server"
)

func main() {
	server.NewHTTPServer(":8080").ListenAndServe()
}
