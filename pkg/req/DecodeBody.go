package req

import (
	"encoding/json"
	"net/http"

	"github.com/Asker231/chat-room.git/pkg/res"
)


func DecodeBody[T any](w http.ResponseWriter, r *http.Request)(*T){
	 var payload T

	 //decode data
	 if err := json.NewDecoder(r.Body).Decode(&payload);err != nil{
		res.Response(w,err.Error(),404)
	 }
	 //validate data
	 if err := ValidateBody(payload); err != nil{
		res.Response(w,err.Error(),204)
	 }
	 
	 //return data

	 return &payload

}