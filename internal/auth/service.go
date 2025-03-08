package auth

import (
	"github.com/Asker231/chat-room.git/internal/user"
	"golang.org/x/crypto/bcrypt"
)


type AuthService struct{
	UserRepo *user.UserRepo
}


func NewAuthService(repo *user.UserRepo)*AuthService{
	return &AuthService{
		UserRepo: repo,
	}
}

func(a *AuthService)RegisterUser(name,email,password string)(bool,error){
	pass,_ := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	user := user.User{
		Name: name,
		Email: email,
		Password: string(pass),
	}
   success,err :=  a.UserRepo.AddUser(&user)
   if err != nil{
	 	return false,err
   }
   return success,nil
}