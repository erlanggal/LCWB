package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
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
	PreviewCSS = "-/css/preview.css"
)

func data() map[string]interface{} {
	data := map[string]interface{}{
		"BaseURL": "http://localhost:8080",
		"css":     "http://localhost:8080/assets/css/",
	}
	return data
}

func main() {
	fmt.Println("Running on http://localhost:8080 ...")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/navbar/{type}/{pID}", navbar)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./-/"))))
	http.Handle("/assets/", r)
	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ = template.ParseFiles(
		"-/all.html",
	)
	t.ExecuteTemplate(w, "layout", data())
}

func navbar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	id := vars["pID"]
	var css_path, html_path string
	var folder []string
	if types == "top" {
		folder = folderList("Navigation Bar/Top")
		css_path = "Navigation Bar/Top/" + id + "/css/style.css"
		html_path = "Navigation Bar/Top/" + id + "/index.html"
	} else if types == "left" {
		folder = folderList("Navigation Bar/Left")
		css_path = "Navigation Bar/Left/" + id + "/css/style.css"
		html_path = "Navigation Bar/Left/" + id + "/index.html"
	}
	t, _ = template.ParseFiles(
		"-/list.html",
	)

	title := "NAVBAR - " + strings.ToUpper(types)
	html := concatFile(html_path)
	html = strings.Replace(html, "css/style.css", BaseURL+"/assets/css/preview.css", -1)
	css := concatFile(css_path)
	js := ""
	writeFile([]byte(css), PreviewCSS)
	executeTemplate(title, r.URL.String(), html, css, js, folder, w)
}

func executeTemplate(title, url, html, css, js string, folderList []string, w http.ResponseWriter) {
	url_split := strings.Split(url, "/")
	data := data()
	data["Title"] = title
	data["list"] = folderList
	data["url"] = url_split[1] + "/" + url_split[2]
	data["html"] = html
	data["css"] = css
	data["js"] = js
	t.ExecuteTemplate(w, "layout", data)
}

func folderList(pathS string) []string {
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

func concatFile(pathS string) string {
	var files string
	file, err := os.Open(pathS)
	if err != nil {
		fmt.Println(err)
	}
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
