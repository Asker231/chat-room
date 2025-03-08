package auth

import "net/http"

type AuthHandler struct{
	
}

func(a *AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}

func(a *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}