package main

import (
	"fmt"
	"github.com/gorilla/mux"
	router "github.com/kevinrizkhy/LCWB/-/controller/router"
	"net/http"
)

func main() {
	fmt.Println("Running on http://localhost:8080 ...")
	r := mux.NewRouter()
	r.HandleFunc("/", router.Home)
	r.HandleFunc("/navbar/{type}/{pID}", router.Navbar)
	r.HandleFunc("/create", router.Create)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./-/view/"))))
	http.Handle("/assets/", r)
	http.ListenAndServe(":8080", r)
}
