package req

import "github.com/go-playground/validator/v10"


func ValidateBody[T any](el T)(error){
	v := validator.New()
	if err := v.Struct(el); err != nil{
		return err
	}
	return nil
}