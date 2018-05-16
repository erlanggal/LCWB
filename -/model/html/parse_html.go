package html

import (
	file "github.com/wellcode/LCWB/-/model/file"
	"html/template"
	"net/http"
	"strings"
)

var t *template.Template

const (
	BaseURL    = "http://localhost:80"
	PreviewCSS = "-/view/css/preview.css"
	PreviewJS = "-/view/js/preview.js"
)

func ExecuteTemplate(title, url, path string, folderList []string, w http.ResponseWriter) {
	html, css, js := file.ConcatFile(path)
	html = strings.Replace(html, "css/style.css", BaseURL+"/assets/css/preview.css", -1)
	html = strings.Replace(html, "<script type='text/javascript' src='js/js.js'></script>", "", -1)
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
