package main

import (
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template
var contactTemplate *template.Template
var faqTemplate *template.Template
var pageNoteFoundTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := pageNoteFoundTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}
	faqTemplate, err = template.ParseFiles(
		"views/faq.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}
	pageNoteFoundTemplate, err = template.ParseFiles(
		"views/404.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	var notFound http.Handler = http.HandlerFunc(pageNotFound)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = notFound
	http.ListenAndServe(":3000", r)
}