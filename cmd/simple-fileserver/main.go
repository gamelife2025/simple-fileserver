package main

import (
	"flag"
	"net/http"

	simplefileserver "github.com/gamelife2025/simple-fileserver"
	"github.com/gin-gonic/gin"
)

var (
	auth = flag.String("auth", "", "example: auth123") //控制上传认证
	addr = flag.String("addr", "127.0.0.1:8000", "")
)

func middlewareAUTH(c *gin.Context) {
	if c.GetHeader("auth") != *auth {
		c.Status(http.StatusForbidden)
		c.Abort()
	}
}

func main() {
	flag.StringVar(&simplefileserver.DEFAULT_DIR_ROOT, "dir", "", "/upload")
	flag.Parse()

	router := gin.Default()
	router.GET("/*filepath", simplefileserver.Brower)
	if *auth != "" {
		router.Use(middlewareAUTH)
	}
	router.POST("/upload", simplefileserver.UploadFiles)
	router.Run(*addr)
}
