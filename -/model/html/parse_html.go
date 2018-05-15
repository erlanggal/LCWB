package html

import (
	file "github.com/kevinrizkhy/LCWB/-/model/file"
	"html/template"
	"net/http"
	"strings"
)

var t *template.Template

const (
	BaseURL    = "http://localhost:8080"
	PreviewCSS = "-/view/css/preview.css"
)

func ExecuteTemplate(title, url, path string, folderList []string, w http.ResponseWriter) {
	html, css, js := file.ConcatFile(path)
	html = strings.Replace(html, "css/style.css", BaseURL+"/assets/css/preview.css", -1)
	file.WriteFile(css, PreviewCSS)
	url_split := strings.Split(url, "/")
	data := map[string]interface{}{
		"BaseURL":  BaseURL,
		"Title":    title,
		"list":     folderList,
		"url":      url_split[1] + "/" + url_split[2],
		"html_str": html,
		"css_str":  css,
		"js_str":   js,
	}
	t, _ = template.ParseFiles(
		"-/view/list.html",
	)
	t.ExecuteTemplate(w, "layout", data)
}
