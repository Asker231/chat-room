package user

import (
	"errors"

	"github.com/Asker231/chat-room.git/pkg/db"
)

type UserRepo struct{
	Db *db.DataBase
}

func NewUserRepo(db *db.DataBase)*UserRepo{
	return &UserRepo{
		Db: db,
	}
}
//add user func
func(r *UserRepo)AddUser(user *User)(bool,error){
	findUser := r.FindByEmail(user.Email)
	if findUser != nil{
		return false,errors.New("Пользователь с таким email уже существует")
	}
	result := r.Db.Db.Create(user)
	if result.Error != nil{
		return false,result.Error
	}
	return true,nil
}

//find by email func


func(r *UserRepo)FindByEmail(email string)(*User){
	var user User
	result := r.Db.Db.First(&user,"email = ?",email)
	if result.Error != nil{
		return nil
	}
	return &user
}