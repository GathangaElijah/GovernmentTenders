package handlers

import (
	"net/http"
	// "tenders/utilities"
	"text/template"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	var templ *template.Template
	templ, err := template.ParseGlob("./static/templates/*.html")

	if err != nil {
		http.Error(writer, "Error loading index template", http.StatusInternalServerError)
		return
	}
	
	indexTemplate := templ.Lookup("index.html")
	err = indexTemplate.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error Executing index template data", http.StatusInternalServerError)
		return
	}
}
