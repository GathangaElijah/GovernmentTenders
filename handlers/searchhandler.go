package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"tenders/models"
	"tenders/utilities"
)
type SearchResult struct{
	Results []models.Contract
}
func Search(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("query")
	d, err := utilities.FileReader()
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	contractsData := utilities.ContractsMapper(d)
	var result []models.Contract
	if query != "" {
		for _, data := range contractsData {
			if strings.Contains(strings.ToLower(data.ContractNumber), strings.ToLower(query)) {

				result = append(result, models.Contract{
					ContractNumber: data.ContractNumber,
					PEName: data.PEName,
					CreatedAt : data.CreatedAt,
				})
			}
		}
	}
	var templ *template.Template
	templ, err = template.ParseGlob("./static/templates/*.html")

	if err != nil {
		http.Error(writer, "Error loading index template", http.StatusInternalServerError)
		return
	}
	indexTemplate := templ.Lookup("searchresults.html")
	data := SearchResult{Results: result}
	log.Println(result)
	indexTemplate.Execute(writer, data)
}
