package main

import (
	"github.com/York-Shawn/activity/activitylog/internal/server"
)

func main() {
	server.NewHTTPServer(":8080").ListenAndServe()
}
