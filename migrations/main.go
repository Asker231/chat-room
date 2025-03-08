package main

import (
	"fmt"
	"os"

	"github.com/Asker231/chat-room.git/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main(){
	//init env
	if err := godotenv.Load("../.env");err != nil{
		fmt.Println(err.Error())
		return 
	} 
	//init db 
	db,err := gorm.Open(postgres.Open(os.Getenv("DNS")),&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		return 
	}
	err = db.AutoMigrate(user.User{})
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("this")
}