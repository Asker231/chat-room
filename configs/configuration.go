package configs

import (
	"os"

	"github.com/joho/godotenv"
)




type AppConfig struct{
	DB DbConfig
	Auth AuthConfig
}

type DbConfig struct{
	DNS string
}

type AuthConfig struct{
	TOKEN string
}

func NewConfig()*AppConfig{
	if err := godotenv.Load(".env");err != nil{
		return nil
	}
	return &AppConfig{
		DB: DbConfig{
			DNS: os.Getenv("DNS"),
		},
		Auth:AuthConfig{
			TOKEN: os.Getenv("TOKEN"),
		},
		
	}

}