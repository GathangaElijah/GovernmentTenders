package main

import (
	"log"
	"net/http"
	"tenders/handlers"
)

func main() {
   
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/search", handlers.Search)

    fileHandler := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("templates",fileHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("problem creating the server\n", err)
		return
	}
}
