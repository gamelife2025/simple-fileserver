package simplefileserver

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var tpl_brower = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>File Server</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
			margin: 0;
			padding: 0;
		}
		.container {
			max-width: 800px;
			margin: 50px auto;
			background: #fff;
			padding: 20px;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		}
		h1 {
			text-align: center;
			color: #333;
		}
		ul {
			list-style: none;
			padding: 0;
		}
		li {
			margin: 10px 0;
		}
		a {
			text-decoration: none;
			font-weight: bold;
		}
		.directory a {
			color: #007BFF;
		}
		.directory a:hover {
			text-decoration: underline;
		}
		.file a {
			color: #86bbd0;
		}
	</style>
</head>
<body>
	<div class="container">
		<h1>Directory listing for {{.Path}}</h1>
		<ul>
			{{if .Parent}}
			<li class="directory"><a href="{{.Parent}}">..</a></li>
			{{end}}
			{{range .Entries}}
			<li class="{{if .IsDir}}directory{{else}}file{{end}}">
				{{if .IsDir}}目录{{else}}文件{{end}} : <a href="{{$.Path}}/{{.Name}}">{{.Name}}</a>
			</li>
			{{end}}
		</ul>
	</div>
</body>
</html>
`

func Brower(c *gin.Context) {
	requestedPath := c.Param("filepath")
	if requestedPath == "/" {
		requestedPath = ""
	}
	fullPath := filepath.Join(DEFAULT_DIR_ROOT, requestedPath)

	// 获取文件信息
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		c.String(http.StatusNotFound, "File or directory not found")
		return
	}

	// 如果是目录，则列出目录内容
	if fileInfo.IsDir() {
		files, err := os.ReadDir(fullPath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read directory")
			return
		}
		tmpl := template.Must(template.New("index").Parse(tpl_brower))

		// 渲染模板并传递文件列表
		entries := make([]gin.H, len(files))
		for i, file := range files {
			entries[i] = gin.H{
				"Name":  file.Name(),
				"IsDir": file.IsDir(),
			}
		}

		parentPath := ""
		if requestedPath != "" {
			parentPath = filepath.Dir(requestedPath)
			if parentPath == "." {
				parentPath = "/"
			}
		}

		tmpl.Execute(c.Writer, gin.H{
			"Path":    requestedPath,
			"Entries": entries,
			"Parent":  parentPath,
		})
	} else {
		// 如果是文件，则直接提供文件下载
		c.File(fullPath)
	}
}
