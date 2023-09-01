package models

import (
	"HPE-golang-test/configs"
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mutex sync.Mutex
var dbConnection *DBConnection

type DBConnection struct {
	client *mongo.Client
	db     *mongo.Database
}

func GetDBConnInstance() *DBConnection {
	if dbConnection == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if dbConnection == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			dbUrl := "mongodb://" + configs.Configs.MongoDBInfo.Host + ":" + configs.Configs.MongoDBInfo.Port
			client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

			if err != nil {
				panic("MongoDB connection failed! in services/models/mongodb_connection.go : " + err.Error())
			}

			db := client.Database(configs.Configs.MongoDBInfo.DbName)

			dbConnection = &DBConnection{
				client: client,
				db:     db,
			}
		}
	}

	return dbConnection
}

func (dbconn *DBConnection) Close() {
	if dbconn.client != nil {
		dbconn.client.Disconnect(context.Background())
	}
}
