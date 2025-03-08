package db

import (


	"github.com/Asker231/chat-room.git/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type DataBase struct{
	Db *gorm.DB

}
func NewDb(conf *configs.DbConfig)*DataBase{
	db,err := gorm.Open(postgres.Open(conf.DNS),&gorm.Config{})
	if err != nil{
		return nil
	}
	return &DataBase{
		Db: db,
	}
}
