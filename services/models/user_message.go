package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var UserModel *UserMessageModel

type UserMessage struct {
	ID         string    `bson:"_id,omitempty"`
	UserID     string    `bson:"user_id"`
	ReplyToken string    `bson:"reply_token"`
	Content    string    `bson:"content"`
	CreateTime time.Time `bson:"create_time"`
}

type UserMessageModel struct {
	dbconn *DBConnection
}

func InitUserModel() *UserMessageModel {
	return &UserMessageModel{
		dbconn: DBConn,
	}
}

func (model *UserMessageModel) Close() {
	model.dbconn.Close()
}

func (model *UserMessageModel) Save(usermsg []UserMessage) error {
	if len(usermsg) == 0 {
		return errors.New("UserMessageModel.Save : input usermsg is empty")
	}

	collection := model.dbconn.db.Collection("user_messages")
	newDocs := model.insertManyTransfer(usermsg)
	_, err := collection.InsertMany(context.Background(), newDocs)
	return err
}

func (model *UserMessageModel) QueryAll() ([]UserMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection("user_messages")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var userMessages []UserMessage
	for cursor.Next(ctx) {
		var msg UserMessage
		if err := cursor.Decode(&msg); err != nil {
			log.Println("User message decode error! in services/models/user_messages.go -> QueryAll : " + err.Error())
		}

		userMessages = append(userMessages, msg)
	}

	return userMessages, nil
}

func (model *UserMessageModel) QueryByUser(userID string) ([]UserMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection("user_messages")

	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var userMessages []UserMessage
	for cursor.Next(ctx) {
		var msg UserMessage
		if err := cursor.Decode(&msg); err != nil {
			log.Println("User message decode error! in services/models/user_messages.go -> QueryByUser : " + err.Error())
			continue
		}

		userMessages = append(userMessages, msg)
	}

	return userMessages, nil
}

func (model *UserMessageModel) insertManyTransfer(inputs []UserMessage) []interface{} {
	var ret []interface{}
	for _, t := range inputs {
		ret = append(ret, t)
	}

	return ret
}
