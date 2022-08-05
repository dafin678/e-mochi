package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed on connecting to database server")
	}

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var DBConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	DBConfig.DBHost = getEnv("DB_HOST", "localhost")
	DBConfig.DBUser = getEnv("DB_USER", "postgres")
	DBConfig.DBPassword = getEnv("DB_PASSWORD", "dafin")
	DBConfig.DBName = getEnv("DB_NAME", "mochistore")
	DBConfig.DBPort = getEnv("DB_PORT", "5432")

	appConfig.AppName = getEnv("APP_NAME", "EmochiStore")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")
	server.Initialize(appConfig, DBConfig)
	server.Run(":" + appConfig.AppPort)
}
