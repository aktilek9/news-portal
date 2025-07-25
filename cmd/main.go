package main

import (
	"log"
	news_portal "news-portal"
	"news-portal/db"
	"news-portal/pkg/handler"
	"news-portal/pkg/jwt"
	"news-portal/pkg/repository"
	"news-portal/pkg/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}
	db, err := db.NewDBConnection(db.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	})

	if err != nil {
		log.Fatalf("failed to connect to database: %s", err.Error())
	}

	router := gin.New()

	jwt := jwt.NewJWTService(os.Getenv("JWT_SECRET"))
	repos := repository.NewRepository(db)
	service := service.NewService(repos, jwt)
	handler.RegisterEndpoint(router, service, jwt)

	server := new(news_portal.Server)
	if err := server.Run("8080", router); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}
