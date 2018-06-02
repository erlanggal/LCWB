package home

import (
	"html/template"
	"net/http"

	"github.com/wellcode/LCWB/-/config"
	session "github.com/wellcode/LCWB/-/model/session"
)

var t *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if session.CheckSession(r) != nil {
		t, _ = template.ParseFiles(
			"-/view/home.html",
		)
		data := map[string]interface{}{
			"BaseURL": config.Base_URL,
		}
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/sign", 302)
	}

}
