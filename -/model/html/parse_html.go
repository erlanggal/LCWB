package html

import (
	file "github.com/wellcode/LCWB/-/model/file"
	"html/template"
	"net/http"
	"strings"
)

var t *template.Template

const (
	BaseURL    = "http://localhost:7070"
	PreviewCSS = "-/view/css/preview.css"
	PreviewJS  = "-/view/js/preview.js"
)

func KatalogTemplate(title, url, path string, folderList []string, w http.ResponseWriter) {
	html, css, js := file.ConcatFile([]string{path})
	file.WriteFile(css, PreviewCSS)
	file.WriteFile(js, PreviewJS)
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
