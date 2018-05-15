package router

import (
	"fmt"
	"github.com/gorilla/mux"
	function "github.com/kevinrizkhy/LCWB/-/model/utility"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var t *template.Template

const (
	BaseURL = "http://localhost:8080"
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
		data := r.FormValue("data")
		//data_split := strings.Split(data, "!@#$%^&*()_+")
		fmt.Fprintln(w, filename+data)
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
	var html_path, css_path, js_path string
	if types == "top" {
		folder = function.FolderList("Navigation Bar/Top")
		css_path = "Navigation Bar/Top/" + id + "/css/style.css"
		html_path = "Navigation Bar/Top/" + id + "/index.html"
		js_path = "Navigation Bar/Top/" + id + "/js/js.js"
	} else if types == "left" {
		folder = function.FolderList("Navigation Bar/Left")
		css_path = "Navigation Bar/Left/" + id + "/css/style.css"
		html_path = "Navigation Bar/Left/" + id + "/index.html"
		js_path = "Navigation Bar/Left/" + id + "/js/js.js"
	}
	title := "NAVBAR - " + strings.ToUpper(types)
	function.ExecuteTemplate(title, r.URL.String(), html_path, css_path, js_path, folder, w)
}
