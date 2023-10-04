package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ReservationModel *ReserveModel

type Reserve struct {
	ID         string    `bson:"_id,omitempty" json:"_id"`
	UserID     string    `bson:"user_id" json:"user_id" validate:"required,max=35,min=30"`
	ReplyToken string    `bson:"reply_token" json:"reply_token" validate:"required,max=35,min=30"`
	Content    string    `bson:"content" json:"content" validate:"max=300"`
	ExpireTime time.Time `bson:"expire_time" json:"expire_time"`
	CreateTime time.Time `bson:"create_time" json:"create_time"`
}

type ReserveModel struct {
	dbconn *DBConnection
}

func init() {
	ReservationModel = InitReservationModel()
}

func InitReservationModel() *ReserveModel {
	return &ReserveModel{
		dbconn: DBConn,
	}
}

func (model *ReserveModel) Close() {
	model.dbconn.Close()
}

// Save a slice of Reserve to db
func (model *ReserveModel) Save(usermsg Reserve) (string, error) {
	if usermsg == (Reserve{}) {
		return "", errors.New("ReserveModel.Save : input usermsg is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection("reservations")
	result, err := collection.InsertOne(ctx, usermsg)
	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}

// Query all documents from db
func (model *ReserveModel) QueryAll() ([]Reserve, error) {
	return model.query(bson.M{})
}

func (model *ReserveModel) QueryById(id string) ([]Reserve, error) {
	return model.query(bson.M{"_id": id})
}

func (model *ReserveModel) QueryByUser(userID string) ([]Reserve, error) {
	return model.query(bson.M{"user_id": userID})
}

// func (model *ReserveModel) DeleteById(id string) error {
// }

func (model *ReserveModel) query(filter interface{}) ([]Reserve, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection("reservations")

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var Reserves []Reserve
	for cursor.Next(ctx) {
		var msg Reserve
		if err := cursor.Decode(&msg); err != nil {
			log.Println("Reservation decode error! " + err.Error())
			continue
		}

		Reserves = append(Reserves, msg)
	}

	return Reserves, nil
}
