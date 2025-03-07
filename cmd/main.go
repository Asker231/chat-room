package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	//init router
	router := http.NewServeMux()

	fmt.Println("")




	//init server 
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
	
}