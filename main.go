package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var tmplh *template.Template

func init() {
	tmplh = template.Must(template.ParseGlob("template/*"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tmplh.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}