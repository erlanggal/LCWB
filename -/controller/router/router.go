package router

import (
	"fmt"
	"github.com/gorilla/mux"
	file "github.com/wellcode/LCWB/-/model/file"
	function "github.com/wellcode/LCWB/-/model/html"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var t *template.Template

const (
	BaseURL = "http://localhost:80"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ = template.ParseFiles(
		"-/view/all.html",
	)
	data := map[string]interface{}{
		"BaseURL": BaseURL,
	}
	t.ExecuteTemplate(w, "layout", data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	param_step := r.URL.Query()["step"]
	step_str := "0"
	if param_step != nil {
		step_str = param_step[0]
	}
	step, _ := strconv.Atoi(step_str)
	if r.Method == "POST" {
		fmt.Sprintln(step)
		//save di db
		filename := r.FormValue("filename")
		data := r.FormValue("data") + "!@#$%^&*()_+" + filename
		fmt.Sprintln(w, data)
		title, html := file.GetFile(".html", "Navigation Bar/Top/")
		fmt.Println(title)
		fmt.Println(html)
	} else {
		if step == 0 {
			t, _ = template.ParseFiles(
				"-/view/form/form.html",
			)
			data := map[string]interface{}{
				"BaseURL": BaseURL,
			}
			t.ExecuteTemplate(w, "layout", data)
		} else if step == 1 {
			//step := param_step[0]
		}
	}
}

func Navbar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	id := vars["pID"]
	var folder []string
	var path string
	if types == "top" {
		folder = file.FolderList("Navigation Bar/Top")
		path = "Navigation Bar/Top/" + id
	} else if types == "left" {
		folder = file.FolderList("Navigation Bar/Left")
		path = "Navigation Bar/Left/" + id
	}
	title := "NAVBAR - " + strings.ToUpper(types)
	function.ExecuteTemplate(title, r.URL.String(), path, folder, w)
}
