package main

import (
	"log"
	"os"

	"github.com/federicodosantos/class-task-tracker/internal/model"
	"github.com/federicodosantos/class-task-tracker/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when load env file:%v", err)
	}

	router := gin.Default()

	db := db.DBInit()

	err = db.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Class{},
		&model.Task{},
		&model.CompletedTask{},
	)

	if err != nil {
		log.Fatalf("Error when migrate database:%v", err)
	}
	log.Println("Success migrate database!")

	router.Run(":" + os.Getenv("APP_PORT"))
}
