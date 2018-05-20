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
	BaseURL = "http://localhost:7070"
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
	data := map[string]interface{}{
		"BaseURL": BaseURL,
	}
	if r.Method == "POST" {
		//save di db
		filename := r.FormValue("templatename")
		path := r.FormValue("path")
		track := ""
		if filename != "" {
			track = filename + "!@#$%^&*()_+"
		} else {
			track = r.FormValue("track") + path + "!@#$%^&*()_+"
		}

		folder_path := ""

		if step == 1 {
			data["Title"] = "NavBar"
			data["type"] = "1"
			folder_path = "Navigation Bar/Top/"
		} else if step == 2 {
			data["Title"] = "NavBar"
			data["type"] = "2"
			folder_path = "Navigation Bar/Top/"
		} else if step == 3 {

		} else if step == 4 {
		}

		if step < 3 {
			array_title, array_html, array_path := file.GetFileList(".html", folder_path)
			data["array_title"] = array_title
			data["array_html"] = array_html
			data["array_path"] = array_path
			data["Track"] = track
			data["step"] = step + 1
			t, _ = template.ParseFiles(
				"-/view/form/form_list.html",
			)
			t.ExecuteTemplate(w, "layout", data)
		} else {
			track_split := strings.Split(track, "!@#$%^&*()_+")
			name := track_split[0]
			arr_path := track_split[1 : len(track_split)-1]
			html, css, js := file.ConcatFile(arr_path)
		}

	} else {
		t, _ = template.ParseFiles(
			"-/view/form/form.html",
		)
		t.ExecuteTemplate(w, "layout", data)
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
