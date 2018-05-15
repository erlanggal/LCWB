package utility

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var t *template.Template

const (
	BaseURL    = "http://localhost:8080"
	PreviewCSS = "-/view/css/preview.css"
)

func ExecuteTemplate(title, url, html_path, css_path, js_path string, folderList []string, w http.ResponseWriter) {
	html := ConcatFile(html_path)
	html = strings.Replace(html, "css/style.css", BaseURL+"/assets/css/preview.css", -1)
	css := ConcatFile(css_path)
	js := ConcatFile(js_path)
	writeFile([]byte(css), PreviewCSS)
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

func FolderList(pathS string) []string {
	var list []string
	parent_folder := strings.Split(pathS, "/")[1]
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if f.IsDir() {
			if f.Name() != "css" && f.Name() != "js" && f.Name() != parent_folder {
				list = append(list, f.Name())
			}
		}
		return nil
	})
	return list
}

func ConcatFile(pathS string) string {
	var files string
	file, _ := os.Open(pathS)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		files += fmt.Sprintln(scanner.Text())
	}
	return files
}

func writeFile(file []byte, path string) {
	ioutil.WriteFile(path, file, 0644)
}
