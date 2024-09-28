package simplefileserver

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	DEFAULT_DIR_ROOT = "upload"

	REG_DIR_PATH, _ = regexp.CompilePOSIX(`^\/(\w+\/?)+$`)

	REG_VALID_PATH = regexp.MustCompile(`^[a-zA-Z0-9/_-]+$`)
)

func UploadFiles(c *gin.Context) {
	dir := c.Query("dir")
	var targetPath = DEFAULT_DIR_ROOT
	if dir != "" {
		if strings.HasPrefix("/", dir) {
			targetPath = targetPath + dir
		} else {
			targetPath = targetPath + "/" + dir
		}
	}
	if !strings.HasSuffix("/", targetPath) {
		targetPath = targetPath + "/"
	}

	if !REG_VALID_PATH.MatchString(targetPath) {
		c.String(http.StatusBadRequest, "invalid dir")
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var count int
	files := form.File["files"]
	for _, file := range files {
		err = c.SaveUploadedFile(file, fileDst(targetPath, file.Filename))
		if err == nil {
			count++
		} else {
			log.Default().Printf("uploda %s err:%s\n", file.Filename, err.Error())
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", count))
}

func fileDst(dir, filename string) string {
	return fmt.Sprintf("%s%s", dir, filename)
}
