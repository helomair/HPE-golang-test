package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserModel *UserMessageModel

type UserMessage struct {
	ID         string    `bson:"_id,omitempty" json:"_id"`
	UserID     string    `bson:"user_id" json:"user_id" validate:"required,max=35,min=30"`
	ReplyToken string    `bson:"reply_token" json:"reply_token" validate:"required,max=35,min=30"`
	Content    string    `bson:"content" json:"content" validate:"max=300"`
	ExpireTime time.Time `bson:"expire_time" json:"expire_time"`
	CreateTime time.Time `bson:"create_time" json:"create_time"`
}

type UserMessageModel struct {
	dbconn *DBConnection
}

func init() {
	UserModel = InitUserModel()
}

func InitUserModel() *UserMessageModel {
	return &UserMessageModel{
		dbconn: DBConn,
	}
}

func (model *UserMessageModel) Close() {
	model.dbconn.Close()
}

// Save a slice of UserMessage to db
func (model *UserMessageModel) Save(usermsg UserMessage) (string, error) {
	if usermsg == (UserMessage{}) {
		return "", errors.New("UserMessageModel.Save : input usermsg is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection("user_messages")
	result, err := collection.InsertOne(ctx, usermsg)
	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}

// Query all documents from db
func (model *UserMessageModel) QueryAll() ([]UserMessage, error) {
	return model.query(bson.M{})
}

func (model *UserMessageModel) QueryById(id string) ([]UserMessage, error) {
	return model.query(bson.M{"_id": id})
}

func (model *UserMessageModel) QueryByUser(userID string) ([]UserMessage, error) {
	return model.query(bson.M{"user_id": userID})
}

// func (model *UserMessageModel) DeleteById(id string) error {
// }

func (model *UserMessageModel) query(filter interface{}) ([]UserMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection("user_messages")

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var userMessages []UserMessage
	for cursor.Next(ctx) {
		var msg UserMessage
		if err := cursor.Decode(&msg); err != nil {
			log.Println("User message decode error! " + err.Error())
			continue
		}

		userMessages = append(userMessages, msg)
	}

	return userMessages, nil
}
