package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//config model
type Config struct {
	DBhost     string
	DBname     string
	DBport     string
	DBpassWord string
	DBsslMode  string
	DBuser     string

	JwtKey  string
	Issuer  string
	AppEnv  string
	AppPort string
}

var AppConfig *Config

//helper func 
func getEnv(key, Fallback string) string {
	if value,exist:=os.LookupEnv(key);exist{
		return value
	}else{return Fallback}
}

//load config data from .env file
func LoadConfig(){
	if err:= godotenv.Load();err!=nil{
		log.Println("no .env file found")
	}

	AppConfig=&Config{
		DBhost: getEnv("DB_HOST","localhost"),
		DBname: getEnv("DB_NAME",""),
		DBport: getEnv("DB_PORT",""),
		DBpassWord: getEnv("DB_PASSWORD",""),
		DBuser: getEnv("DB_USER",""),
		DBsslMode: getEnv("DB_SSLMODE",""),
		JwtKey: getEnv("JWT_KEY",""),
		AppEnv: getEnv("APP_ENV","developement"),
		AppPort: getEnv("APP_PORT",":8080"),
		Issuer: getEnv("ISSUER",""),
	}
	log.Println("config loaded")
}


