package main

import (
	"log"
	"net/http"

	"github.com/Asker231/chat-room.git/configs"
	"github.com/Asker231/chat-room.git/internal/auth"
	"github.com/Asker231/chat-room.git/internal/user"
	"github.com/Asker231/chat-room.git/pkg/db"
)
func main() {
	conf := configs.NewConfig()
	//init router
	router := http.NewServeMux()
	//init db
	dtabase := db.NewDb(&conf.DB)
	//repo
	userRepo:= user.NewUserRepo(dtabase)
	//init service
	serviceAuth := auth.NewAuthService(userRepo)
	//init  handlers
	auth.NewAuthHandler(router,serviceAuth)

	//init server 
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
	
}