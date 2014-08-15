package home

import (
	"html/template"
	"net/http"
)

//LayoutPath of home page
const LayoutPath string = "templates/layout.html"

//GetHomePage FUNCTION return home page
func GetHomePage(rw http.ResponseWriter, req *http.Request) {
	var Templates *template.Template

	type Page struct {
		Title string
	}

	p := Page{
		Title: "home",
	}

	Templates = template.Must(template.ParseFiles("templates/home/home.html", LayoutPath))
	err := Templates.ExecuteTemplate(rw, "base", p)
	if err != nil {

	}
}
