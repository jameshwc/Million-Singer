package model

import (
	"context"
	"fmt"
	"time"

	"github.com/jameshwc/Million-Singer/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jameshwc/Million-Singer/conf"
)

var db *sql.DB
var mongoClient *mongo.Client
var mongoDB *mongo.Database
var lyricsCollection *mongo.Collection

func Setup(externalDB *sql.DB) {
	var err error
	db = externalDB
	if db == nil {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DBconfig.User,
			conf.DBconfig.Password,
			conf.DBconfig.Host,
			conf.DBconfig.Name))
		if err != nil {
			log.Fatalf("models.Setup SQL err: %v", err)
		}
	}

	mongoClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Fatal("models.Setup mongodb err: %v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal("models.Setup mongodb err: %v", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("models.Setup mongodb err: %v", err)
	}

	mongoDB = mongoClient.Database("million_singer")
	lyricsCollection = mongoDB.Collection("lyrics")
}
