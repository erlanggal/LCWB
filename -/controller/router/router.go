package router

import (
	"github.com/gorilla/mux"
	config "github.com/wellcode/LCWB/-/config"
	create "github.com/wellcode/LCWB/-/controller/create"
	sign "github.com/wellcode/LCWB/-/controller/sign"
	file "github.com/wellcode/LCWB/-/model/file"
	html_exe "github.com/wellcode/LCWB/-/model/html"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var t *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ = template.ParseFiles(
		"-/view/all.html",
	)
	data := map[string]interface{}{
		"BaseURL": config.Base_URL,
	}
	t.ExecuteTemplate(w, "layout", data)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	sign.SignIn(w, r)
}

func Create(w http.ResponseWriter, r *http.Request) {
	param_step := r.URL.Query()["step"]
	step_str := "0"
	if param_step != nil {
		step_str = param_step[0]
	}
	step, _ := strconv.Atoi(step_str)
	create.Create(w, r, step)
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
	html_exe.KatalogTemplate(title, r.URL.String(), path, folder, w)
}
