package auth

import (
	"net/http"
	"github.com/Asker231/chat-room.git/pkg/req"
)

type AuthHandler struct{
	
}

func NewAuthHandler(router *http.ServeMux){
	a := &AuthHandler{}

	router.HandleFunc("POST /auth/register",a.Register())
	router.HandleFunc("GET /auth",a.Register())
}

func(a *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
			_  = req.DecodeBody[RegisterRequest](w,r)
	}
}

func(a *AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_  = req.DecodeBody[LoginRequest](w,r)

	}
}

