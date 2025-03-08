package auth

import (
	"fmt"
	"net/http"

	"github.com/Asker231/chat-room.git/pkg/req"
)

type AuthHandler struct{
	Service *AuthService	
}

func NewAuthHandler(router *http.ServeMux,setvice *AuthService){
	a := &AuthHandler{
		Service: setvice,
	}

	router.HandleFunc("POST /auth/register",a.Register())
}

func(a *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
			body := req.DecodeBody[RegisterRequest](w,r)

			isRegister,err := a.Service.RegisterUser(body.Name,body.Email,body.Password)
			if err != nil{
				return
			}
			fmt.Println(isRegister)

	}
}

func(a *AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_  = req.DecodeBody[LoginRequest](w,r)

	}
}

