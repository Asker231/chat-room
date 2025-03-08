package res

import (
	"encoding/json"
	"fmt"
	
	"net/http"
)

func Response(w http.ResponseWriter,data any,statusCode int){
	if err := json.NewEncoder(w).Encode(data); err != nil{
		fmt.Println(err)
	}
}