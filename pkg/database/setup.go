package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.GetEnv("DB_USER", ""), env.GetEnv("DB_PASSWORD", ""), env.GetEnv("DB_HOST", "127.0.0.1"), env.GetEnv("DB_PORT", "3306"), env.GetEnv("DB_NAME", ""))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	err = DB.AutoMigrate(&models.User{}, &models.UserSession{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("successfully migrate database!")

}

func SetupMongoDB() {
	uri := env.GetEnv("MONGODB_URI", "")
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	coll := client.Database("messaging_db").Collection("message_history")
	MongoDB = coll

	log.Println("successfully connected to mongoDB")
}
