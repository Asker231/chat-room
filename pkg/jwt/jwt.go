package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct{
	Email string
}

type JWT struct{
	SECRET string	
}

func NewJWT(secret string)*JWT{
	return &JWT{
		SECRET: secret,
	}
}


func(j *JWT)Generate(data JWTData)(string,error){
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":data.Email,
	})
	token,err := t.SignedString(j.SECRET)
	if err != nil{
		return err.Error(),nil
	}
	return token,nil
}

func(j *JWT)ParseJWT(token string)(bool,*JWTData){
	t,err := jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
		return []byte(j.SECRET),nil
	})
	if err != nil{
		return false,nil
	}
	email := t.Claims.(jwt.MapClaims)["email"]

	return t.Valid,&JWTData{
		Email: email.(string),
	}
}