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
	ID         string    `bson:"_id,omitempty" json:"_id"`
	UserID     string    `bson:"user_id" json:"user_id"`
	ReplyToken string    `bson:"reply_token" json:"reply_token"`
	Content    string    `bson:"content" json:"content"`
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
func (model *UserMessageModel) Save(usermsg []UserMessage, targetCollection string) error {
	if len(usermsg) == 0 {
		return errors.New("UserMessageModel.Save : input usermsg is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection(targetCollection)
	newDocs := model.insertManyTransfer(usermsg)
	_, err := collection.InsertMany(ctx, newDocs)
	return err
}

// Query all documents from db
func (model *UserMessageModel) QueryAll(targetCollection string) ([]UserMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection(targetCollection)
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

func (model *UserMessageModel) QueryByUser(userID string, targetCollection string) ([]UserMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := model.dbconn.db.Collection(targetCollection)

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
