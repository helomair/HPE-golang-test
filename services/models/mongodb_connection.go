package models

import (
	"HPE-golang-test/configs"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBConn *DBConnection

type DBConnection struct {
	client *mongo.Client
	db     *mongo.Database
}

func InitDBConnection() *DBConnection {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	dbUrl := "mongodb://" + configs.Configs.MongoDBInfo.Host + ":" + configs.Configs.MongoDBInfo.Port
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

	if err != nil {
		panic("MongoDB connection failed! in services/models/mongodb_connection.go : " + err.Error())
	}

	db := client.Database(configs.Configs.MongoDBInfo.DbName)

	return &DBConnection{
		client: client,
		db:     db,
	}
}

func (dbconn *DBConnection) Close() {
	if dbconn.client != nil {
		dbconn.client.Disconnect(context.Background())
	}
}
