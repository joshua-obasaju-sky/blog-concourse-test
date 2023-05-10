package db

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
  *DBConfig
}

type DBConfig struct {
	Username string
	Password string
	Host string
	Port string
}


func GetConfig () *Config{
  err := godotenv.Load()

  if err != nil {
	panic(err)
  }


  return &Config{
	DBConfig: &DBConfig{
		 Username: os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	},
  }



}