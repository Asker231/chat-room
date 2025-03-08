package user

import "github.com/Asker231/chat-room.git/pkg/db"

type UserRepo struct{
	Db *db.DataBase
}

func NewUserRepo(db *db.DataBase)*UserRepo{
	return &UserRepo{
		Db: db,
	}
}

func(r *UserRepo)AddUser(user *User)(bool,error){
	result := r.Db.Db.Create(user)
	if result.Error != nil{
		return false,result.Error
	}
	return true,nil
}