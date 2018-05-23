package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/wellcode/LCWB/-/config"
	router "github.com/wellcode/LCWB/-/controller/router"
	database "github.com/wellcode/LCWB/-/model/db"
)

func main() {
	fmt.Println("Running on " + config.Base_URL + " ...")
	database.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/", router.Home)
	r.HandleFunc("/sign", router.SignIn)
	r.HandleFunc("/navbar/{type}/{pID}", router.Navbar)
	r.HandleFunc("/about/{pID}", router.About)
	r.HandleFunc("/form/{pID}", router.Form)
	r.HandleFunc("/footer/{pID}", router.Footer)
	r.HandleFunc("/create", router.Create)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./-/view/"))))
	http.Handle("/assets/", r)
	http.ListenAndServe(config.Base_Port, r)
}
