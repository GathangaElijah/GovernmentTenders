package main

import (
	"fmt"
	"log"
	"net/http"
	"tenders/handlers"
	"tenders/utilities"
)

func main() {
    data, err := utilities.FileReader()
    if err != nil{
        fmt.Println("error\n",err)
        return
    }
    fmt.Println("This is data", data[1])
    fmt.Println("This is the length of the data", len(data[1]))
    
    fmt.Println(utilities.ContractsMapper(data))
	http.HandleFunc("/", handlers.Index)

    fileHandler := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("templates",fileHandler))
	err = http.ListenAndServe(":8080", nil)
    // log.Println("Server started at : http:local")
	if err != nil {
		log.Println("problem creating the server\n", err)
		return
	}
}
