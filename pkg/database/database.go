package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var DB *gorm.DB

var MongoDB *mongo.Collection
