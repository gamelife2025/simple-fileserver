package simplefileserver

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	DEFAULT_DIR_ROOT = "upload"

	REG_DIR_PATH, _ = regexp.CompilePOSIX(`^\/(\w+\/?)+$`)
)

func UploadFiles(c *gin.Context) {
	dir := c.Query("dir")
	if dir != "" && !strings.HasPrefix("/", dir) {
		dir = "/" + dir
		if !REG_DIR_PATH.MatchString(dir) {
			c.String(http.StatusBadRequest, "invalid dir")
			return
		}
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var count int
	files := form.File["files"]
	for _, file := range files {
		err = c.SaveUploadedFile(file, fileDst(dir, file.Filename))
		if err == nil {
			count++
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", count))
}

func fileDst(dir, filename string) string {
	var targetdir string = DEFAULT_DIR_ROOT + dir
	return fmt.Sprintf("%s%s", targetdir, filename)
}
