package main

import (
	"log"
	"net/http"

	"github.com/Asker231/chat-room.git/internal/auth"
)
func main() {
	//init router
	router := http.NewServeMux()


	//init  handlers
	auth.NewAuthHandler(router)



	//init server 
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
	
}