package main

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// daha iyi yonlendirmeler icin
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/deneme", deneme)


	http.ListenAndServe(":80", r)
}

func home(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	view,_ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}

func deneme(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	username := r.FormValue("username")
	view,_ := template.ParseFiles("deneme.html")
	
	data := username
	view.Execute(w, data)
}
