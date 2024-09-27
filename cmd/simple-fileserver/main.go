package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	simplefileserver "github.com/zqb7/simple-fileserver"
)

var (
	auth = flag.String("auth", "", "example: auth123")
	addr = flag.String("addr", "127.0.0.1:8000", "")
)

func middlewareAUTH(c *gin.Context) {
	if c.GetHeader("auth") != *auth {
		c.Status(http.StatusForbidden)
		c.Abort()
	}
}

func main() {
	flag.Parse()

	router := gin.Default()
	if *auth != "" {
		router.Use(middlewareAUTH)
	}
	router.POST("/upload", simplefileserver.UploadFiles)
	router.Run(*addr)
}
