package database

import (
	"blog-api/models"
	"blog-api/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *gorm.DB
var MongoDB *mongo.Database

func Connect() {
	switch config.AppConfig.DBType {
	case "sqlite":
		connectSQLite()
	case "postgres":
		connectPostgres()
	case "mongodb":
		connectMongoDB()
	default:
		log.Fatal("Invalid db_type in config")
	}
}

func connectSQLite() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.AppConfig.SQLitePath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQLite:", err)
	}
	DB.AutoMigrate(&models.BlogPost{})
}

func connectPostgres() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.AppConfig.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Postgres:", err)
	}
	DB.AutoMigrate(&models.BlogPost{})
}

func connectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(config.AppConfig.MongoURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	MongoDB = client.Database(config.AppConfig.MongoDBName)
}
