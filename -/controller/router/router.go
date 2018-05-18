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

		fmt.Sprintln(data)
		if step == 1 {
			data["Title"] = "NavBar"
			data["type"] = "1"
		} else if step == 2 {
			fmt.Fprint(w, track)
		} else if step == 3 {
		} else if step == 4 {
		}
		array_html_template := []template.HTML{}
		array_title, array_html, array_path := file.GetFile(".html", "Navigation Bar/Top/")
		for i := 0; i < len(array_html); i++ {
			array_html[i] = strings.Replace(array_html[i], "<link rel=\"stylesheet\" href=\"css/style.css\">", "", -1)
			array_html[i] = strings.Replace(array_html[i], "<script type=\"text/javascript\" src=\"/js/js.js\"></script>", "", -1)
			var temp string
			temp = array_html[i]
			array_html_template = append(array_html_template, template.HTML(temp))
		}
		fmt.Println(len(array_html_template))
		data["array_title"] = array_title
		data["array_html"] = array_html_template
		data["array_path"] = array_path
		data["Track"] = track
		t, _ = template.ParseFiles(
			"-/view/form/form_list.html",
		)
		t.ExecuteTemplate(w, "layout", data)

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
